package godruid

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupervisorKafka(t *testing.T) {
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
							DimensionSpec: &EmptyDimension{},
						},
					},
					MetricsSpec: []*MetricsSpec{
						{
							Type:      "variance",
							Name:      "valueVariance",
							FieldName: "value",
							InputType: "double",
							Estimator: "population",
						},
					},
				},
			},
			expected: `{
				"type": "kafka",
				"ioConfig": {
				  "type": "kafka",
				  "consumerProperties": {
					"bootstrap.servers": "localhost:9092",
					"sasl.mechanism": "SCRAM-SHA-512",
					"security.protocol": "SASL_PLAINTEXT",
					"sasl.jaas.config": "org.apache.kafka.common.security.scram.ScramLoginModule required username=\"fee\" password=\"foo\";"
				  },
				  "topic": "test-topic",
				  "lateMessageRejectionPeriod": "PT24H"
				},
				"tuningConfig": {
				  "type": "kafka",
				  "maxRowsPerSegment": 1000000
				},
				"dataSchema": {
				  "dataSource": "test-datasource",
				  "granularitySpec": {
					"type": "uniform",
					"segmentGranularity": "HOUR",
					"queryGranularity": "none",
					"rollup": false
				  },
				  "parser": {
					"type": "string",
					"parseSpec": {
					  "format": "json",
					  "timestampSpec": {
						"column": "timestamp",
						"format": "iso"
					  },
					  "dimensionsSpec": {}
					}
				  },
				  "metricsSpec": [
					{
					  "type": "variance",
					  "name": "valueVariance",
					  "fieldName": "value",
					  "inputType": "double",
					  "estimator": "population"
					}
				  ]
				}
			  }`,
		},
	}
	for msg, p := range patterns {
		t.Run(msg, func(t *testing.T) {
			actual, err := json.Marshal(p.input)
			assert.NoError(t, err)
			buf := new(bytes.Buffer)
			if err := json.Compact(buf, []byte(p.expected)); err != nil {
				assert.NoError(t, err)
			}
			assert.Equal(t, buf.String(), string(actual))
		})
	}
}
