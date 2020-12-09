package godruid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateSupervisor(t *testing.T) {
	t.Parallel()
	patterns := map[string]struct {
		input    *SupervisorKafka
		expected string
	}{
		"success": {
			input: &SupervisorKafka{
				Type:         "kafka",
				IOConfig:     IOConfigKafka("test-topic", "PT24H", "localhost:9092", "SCRAM-SHA-512", "SASL_PLAINTEXT", "org.apache.kafka.common.security.scram.ScramLoginModule required username=\"fee\" password=\"foo\";"),
				TuningConfig: TuningConfigKafka(1000000),
				DataSchema: &DataSchema{
					Datasource:      "test-datasource",
					GranularitySpec: GranUniform("HOUR", "none", false),
					Parser: &Parser{
						Type: "string",
						ParseSpec: &ParseSpec{
							Format:        "json",
							TimestampSpec: &TimestampSpec{Column: "timestamp", Format: "iso"},
						},
					},
					MetricsSpec: []*MetricsSpec{{Type: "count", Name: "count-name"}},
				},
			},
		},
	}
	for msg, p := range patterns {
		t.Run(msg, func(t *testing.T) {
			client := OverlordClient{
				Url:        "http://localhost:8081",
				HttpClient: &http.Client{},
				Debug:      true,
			}
			err := client.CreateOrUpdateSupervisor(context.Background(), p.input, "")
			fmt.Printf("requst: %s", client.LastRequest)
			assert.NoError(t, err)
			fmt.Printf("response: %s", client.LastResponse)
			fmt.Printf("query.SubmitResult: %v", p.input.SubmitResult)
		})
	}
}
