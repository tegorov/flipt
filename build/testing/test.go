package testing

import (
	"context"
	"encoding/json"

	"dagger.io/dagger"
)

func Unit(ctx context.Context, client *dagger.Client, flipt *dagger.Container) error {
	// create Redis service container
	redisSrv := client.Container().
		From("redis").
		WithExposedPort(6379).
		WithExec(nil)

	gitea := client.Container().
		From("gitea/gitea:1.21.1").
		WithExposedPort(3000).
		WithExec(nil)

	flipt = flipt.
		WithServiceBinding("gitea", gitea).
		WithExec([]string{"go", "run", "./build/internal/cmd/gitea/...", "-gitea-url", "http://gitea:3000", "-testdata-dir", "./internal/storage/fs/git/testdata"})

	out, err := flipt.Stdout(ctx)
	if err != nil {
		return err
	}

	var push map[string]string
	if err := json.Unmarshal([]byte(out), &push); err != nil {
		return err
	}

	minio := client.Container().
		From("quay.io/minio/minio:latest").
		WithExposedPort(9009).
		WithEnvVariable("MINIO_ROOT_USER", "user").
		WithEnvVariable("MINIO_ROOT_PASSWORD", "password").
		WithExec([]string{"server", "/data", "--address", ":9009"})

	flipt = flipt.
		WithServiceBinding("minio", minio).
		WithEnvVariable("AWS_ACCESS_KEY_ID", "user").
		WithEnvVariable("AWS_SECRET_ACCESS_KEY", "password").
		WithExec([]string{"go", "run", "./build/internal/cmd/minio/...", "-minio-url", "http://minio:9009", "-testdata-dir", "./internal/storage/fs/s3/testdata"})

	flipt, err = flipt.
		WithServiceBinding("redis", redisSrv).
		WithEnvVariable("REDIS_HOST", "redis:6379").
		WithEnvVariable("TEST_GIT_REPO_URL", "http://gitea:3000/root/features.git").
		WithEnvVariable("TEST_GIT_REPO_HEAD", push["HEAD"]).
		WithEnvVariable("TEST_S3_ENDPOINT", "http://minio:9009").
		WithExec([]string{"go", "test", "-race", "-p", "1", "-coverprofile=coverage.txt", "-covermode=atomic", "./..."}).
		Sync(ctx)
	if err != nil {
		return err
	}

	// attempt to export coverage if its exists
	_, _ = flipt.File("coverage.txt").Export(ctx, "coverage.txt")

	return nil
}
