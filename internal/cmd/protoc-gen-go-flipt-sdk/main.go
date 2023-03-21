package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	importPath  = "go.flipt.io/flipt/sdk"
	emptyImport = "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	var flags flag.FlagSet
	grpcAPIConfig := flags.String("grpc_api_configuration", "", "path to GRPC API configuration")

	protogen.Options{ParamFunc: flags.Set}.Run(func(gen *protogen.Plugin) error {
		if *grpcAPIConfig == "" {
			fmt.Fprintln(os.Stderr, "path to grpc API configuration required")
			os.Exit(1)
		}

		// We have some use of the optional feature in our proto3 definitions.
		// This broadcasts that our plugin supports it and hides the generated
		// warning.
		gen.SupportedFeatures |= uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}

			generateSubSDK(gen, f)
		}

		generateSDK(gen)

		generateGRPC(gen)

		generateHTTP(gen, *grpcAPIConfig)

		return nil
	})
}

func generateSDK(gen *protogen.Plugin) {
	g := gen.NewGeneratedFile("sdk.gen.go", importPath)
	g.P("// Code generated by protoc-gen-go-flipt-sdk. DO NOT EDIT.")
	g.P()
	g.P("package sdk")
	g.P()
	g.P("type Transport interface {")
	var types [][2]string
	for _, file := range gen.Files {
		if !file.Generate {
			continue
		}

		var (
			typ     = strings.Title(string(file.GoPackageName))
			method  = typ + "Client"
			returns = method
		)

		if len(file.Services) < 2 {
			returns = relativeImport(g, file, file.Services[0].GoName+"Client")
		}

		types = append(types, [...]string{typ, method})
		g.P(method, "() ", returns)
	}
	g.P("}")
	g.P()

	g.P(sdkBase)
	g.P()

	for _, t := range types {
		g.P("func (s SDK) ", t[0], "() *", t[0], "{")
		g.P("return &", t[0], "{")
		g.P("transport: s.transport.", t[1], "(),")
		g.P("tokenProvider: s.tokenProvider,")
		g.P("}")
		g.P("}\n")
	}
}

// generateSubSDK generates a .pb.sdk.go file containing a single SDK structure
// which represents an entire package from within the entire Flipt SDK API.
func generateSubSDK(gen *protogen.Plugin, file *protogen.File) (typ, client string) {
	filename := string(file.GoPackageName) + ".sdk.gen.go"
	g := gen.NewGeneratedFile(filename, importPath)
	g.P("// Code generated by protoc-gen-go-flipt-sdk. DO NOT EDIT.")
	g.P()
	g.P("package sdk")
	g.P()

	context := importPackage(g, "context")

	oneServicePackage := len(file.Services) == 1

	// define client structure
	typ = strings.Title(string(file.GoPackageName))
	client = relativeImport(g, file, typ+"Client")

	// We generate an interface which conjoins all the client interfaces
	// generated by the gRPC protoc generator.
	// Our gRPC and HTTP wrapper generators will take care of
	// bundling these clients appropriately for the SDK to consume.
	if !oneServicePackage {
		client = typ + "Client"

		g.P("type ", typ, "Client interface {")
		for _, srv := range file.Services {
			g.P(srv.GoName+"Client", "()", relativeImport(g, file, srv.GoName+"Client"))
		}
		g.P("}\n")

		g.P("type ", typ, " struct {")
		g.P("transport ", typ, "Client")
		g.P("tokenProvider ", "ClientTokenProvider")
		g.P("}\n")
	}

	for _, srv := range file.Services {
		serviceName := srv.GoName
		if oneServicePackage {
			serviceName = typ
		}

		g.P("type ", serviceName, " struct {")
		g.P("transport ", relativeImport(g, file, srv.GoName+"Client"))
		g.P("tokenProvider ", "ClientTokenProvider")
		g.P("}\n")

		if !oneServicePackage {
			g.P("func (s ", typ, ") ", srv.GoName, "() *", srv.GoName, "{")
			g.P("return &", srv.GoName, "{")
			g.P("transport: s.transport.", srv.GoName+"Client", "(),")
			g.P("tokenProvider: ", "s.tokenProvider,")
			g.P("}")
			g.P("}")
		}

		for _, method := range srv.Methods {
			var (
				signature       = []any{"func (x *", serviceName, ") ", method.GoName, "(ctx ", context("Context")}
				returnStatement = []any{"x.transport.", method.GoName, "(ctx, "}
			)

			if method.Input.GoIdent.GoImportPath != emptyImport {
				signature = append(signature, ", v *", method.Input.GoIdent)
				returnStatement = append(returnStatement, "v)")
			} else {
				returnStatement = append(returnStatement, "&", method.Input.GoIdent, "{})")
			}

			if method.Output.GoIdent.GoImportPath != emptyImport {
				g.P(append(signature, ") (*", method.Output.GoIdent, ", error) {")...)
				threadAuth(g, "x", "return nil, err")
				g.P(append([]any{"return "}, returnStatement...)...)
			} else {
				g.P(append(signature, ") error {")...)
				threadAuth(g, "x", "return err")
				g.P(append([]any{"_, err := "}, returnStatement...)...)
				g.P("return err")
			}

			g.P("}\n")
		}
	}
	return
}

func threadAuth(g *protogen.GeneratedFile, receiver, returnErr string) {
	metadata := importPackage(g, "google.golang.org/grpc/metadata")
	g.P("if ", receiver, ".tokenProvider != nil {")
	g.P("token, err := ", receiver, ".tokenProvider.ClientToken()")
	g.P("if err != nil { ", returnErr, " }")
	g.P()
	g.P("ctx = ", metadata("AppendToOutgoingContext"), `(ctx, "authorization", "Bearer "+token)`)
	g.P("}")
	g.P()
}

func unexport(v string) string {
	return strings.ToLower(v[:1]) + v[1:]
}

func importPackage(g *protogen.GeneratedFile, pkg string) func(string) string {
	return func(name string) string {
		return g.QualifiedGoIdent(protogen.GoIdent{
			GoImportPath: protogen.GoImportPath(pkg),
			GoName:       name,
		})
	}
}

func relativeImport(g *protogen.GeneratedFile, file *protogen.File, name string) string {
	return g.QualifiedGoIdent(protogen.GoIdent{
		GoImportPath: file.GoImportPath,
		GoName:       name,
	})
}

const sdkBase = `// ClientTokenProvider is a type which when requested provides a
// client token which can be used to authenticate RPC/API calls
// invoked through the SDK.
type ClientTokenProvider interface {
	ClientToken() (string, error)
}

// SDK is the definition of Flipt's Go SDK.
// It depends on a pluggable transport implementation and exposes
// a consistent API surface area across both transport implementations.
// It also provides consistent client-side instrumentation and authentication
// lifecycle support.
type SDK struct {
	transport Transport
    tokenProvider ClientTokenProvider
}

// Option is a functional option which configures the Flipt SDK.
type Option func(*SDK)

// WithClientTokenProviders returns an Option which configures
// any supplied SDK with the provided ClientTokenProvider.
func WithClientTokenProvider(p ClientTokenProvider) Option {
	return func(s *SDK) {
        s.tokenProvider = p
    }
}

// StaticClientTokenProvider is a string which is supplied as a static client token
// on each RPC which requires authentication.
type StaticClientTokenProvider string

// ClientToken returns the underlying string that is the StaticClientTokenProvider.
func (p StaticClientTokenProvider) ClientToken() (string, error) {
    return string(p), nil
}

// New constructs and configures a Flipt SDK instance from
// the provided Transport implementation and options.
func New(t Transport, opts ...Option) SDK {
    sdk := SDK{transport: t}

    for _, opt := range opts { opt(&sdk) }

    return sdk
}`