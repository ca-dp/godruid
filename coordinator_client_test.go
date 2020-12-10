package godruid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateRetentionRule(t *testing.T) {
	t.Parallel()
	patterns := map[string]struct {
		rules    *RetentionRules
		expected string
	}{
		"success": {
			rules: &RetentionRules{
				Rules: []*RetentionRule{
					{Type: "dropBeforeByPeriod", Period: "P2M"},
					{Type: "loadForever"},
				},
			},
		},
	}
	for msg, p := range patterns {
		t.Run(msg, func(t *testing.T) {
			client := CoordinatorClient{
				Url:        "http://localhost:8081",
				HttpClient: &http.Client{},
				Debug:      true,
			}
			err := client.CreateOrUpdateRetentionRule(context.Background(), "wikiticker-2015-09-12-sampled", p.rules, "")
			fmt.Printf("requst: %s", client.LastRequest)
			assert.NoError(t, err)
			fmt.Println("err", err)
			fmt.Printf("response: %s", client.LastResponse)
			fmt.Printf("spec.RetentionRuleResult: %v", p.rules.RetentionRuleResult)
		})
	}
}
