package analytics

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

var errorNotFound = errors.New("event not found")

type AnalyticsStoreMutator interface {
	IncrementFlagEvaluationCounts(ctx context.Context, responses []*EvaluationResponse) error
	Close() error
}

type EvaluationResponse struct {
	FlagKey      string    `json:"flagKey,omitempty"`
	NamespaceKey string    `json:"namespaceKey,omitempty"`
	Reason       string    `json:"reason,omitempty"`
	Match        bool      `json:"match,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

// AnalyticsSinkSpanExporter implements SpanExporter.
type AnalyticsSinkSpanExporter struct {
	logger                *zap.Logger
	analyticsStoreMutator AnalyticsStoreMutator
}

// NewAnalyticsSinkSpanExporter is the constructor function for an AnalyticsSpanExporter.
func NewAnalyticsSinkSpanExporter(logger *zap.Logger, analyticsStoreMutator AnalyticsStoreMutator) *AnalyticsSinkSpanExporter {
	return &AnalyticsSinkSpanExporter{
		logger:                logger,
		analyticsStoreMutator: analyticsStoreMutator,
	}
}

// transformSpanEventToEvaluationResponse is a convenience function to transform a span event into an EvaluationResponse.
func transformSpanEventToEvaluationResponse(event sdktrace.Event) (*EvaluationResponse, error) {
	for _, attr := range event.Attributes {
		if string(attr.Key) == "flipt.evaluation.response" {
			evaluationResponseBytes := []byte(attr.Value.AsString())
			var evaluationResponse *EvaluationResponse

			if err := json.Unmarshal(evaluationResponseBytes, &evaluationResponse); err != nil {
				return nil, err
			}

			return evaluationResponse, nil
		}
	}

	return nil, errorNotFound
}

// ExportSpans transforms the spans into []*EvaluationResponse which the mutator takes to store into an analytics store.
func (a *AnalyticsSinkSpanExporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	evaluationResponses := make([]*EvaluationResponse, 0)

	for _, span := range spans {
		for _, event := range span.Events() {
			evaluationResponse, err := transformSpanEventToEvaluationResponse(event)
			if err != nil && !errors.Is(err, errorNotFound) {
				a.logger.Error("event not decodable into evaluation response", zap.Error(err))
				continue
			}

			evaluationResponses = append(evaluationResponses, evaluationResponse)
		}
	}

	return a.analyticsStoreMutator.IncrementFlagEvaluationCounts(ctx, evaluationResponses)
}

// Shutdown closes resources for an AnalyticsStoreMutator.
func (a *AnalyticsSinkSpanExporter) Shutdown(_ context.Context) error {
	return a.analyticsStoreMutator.Close()
}
