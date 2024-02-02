package analytics_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	panalytics "go.flipt.io/flipt/internal/server/analytics"
	"go.flipt.io/flipt/internal/server/analytics/clickhouse"
	analyticstesting "go.flipt.io/flipt/internal/server/analytics/testing"
	"go.flipt.io/flipt/rpc/flipt/analytics"
)

type AnalyticsDBTestSuite struct {
	suite.Suite
	client *clickhouse.Client
}

func TestAnalyticsDBTestSuite(t *testing.T) {
	suite.Run(t, new(AnalyticsDBTestSuite))
}

func (a *AnalyticsDBTestSuite) SetupSuite() {
	setup := func() error {
		db, err := analyticstesting.Open()
		if err != nil {
			return err
		}

		c := &clickhouse.Client{
			Conn: db.DB,
		}

		a.client = c

		return nil
	}

	a.Require().NoError(setup())
}

func (a *AnalyticsDBTestSuite) TestAnalyticsMutationAndQuery() {
	t := a.T()

	now := time.Now().UTC()
	for i := 0; i < 5; i++ {
		err := a.client.IncrementFlagEvaluationCounts(context.TODO(), []*panalytics.EvaluationResponse{
			{
				NamespaceKey: "default",
				FlagKey:      "flag1",
				Match:        true,
				Reason:       "MATCH_EVALUATION_REASON",
				Timestamp:    now,
			},
			{
				NamespaceKey: "default",
				FlagKey:      "flag1",
				Match:        true,
				Reason:       "MATCH_EVALUATION_REASON",
				Timestamp:    now,
			},
			{
				NamespaceKey: "default",
				FlagKey:      "flag1",
				Match:        true,
				Reason:       "MATCH_EVALUATION_REASON",
				Timestamp:    now,
			},
			{
				NamespaceKey: "default",
				FlagKey:      "flag1",
				Match:        true,
				Reason:       "MATCH_EVALUATION_REASON",
				Timestamp:    now,
			},
			{
				NamespaceKey: "default",
				FlagKey:      "flag1",
				Match:        true,
				Reason:       "MATCH_EVALUATION_REASON",
				Timestamp:    now,
			},
		})
		require.Nil(t, err)
	}

	_, values, err := a.client.GetFlagEvaluationsCount(context.TODO(), &analytics.GetFlagEvaluationsCountRequest{
		NamespaceKey: "default",
		FlagKey:      "flag1",
		From:         now.Add(-time.Hour).Format(time.DateTime),
		To:           now.Format(time.DateTime),
	})
	require.Nil(t, err)

	assert.Len(t, values, 1)
	assert.Equal(t, values[0], float32(25))
}
