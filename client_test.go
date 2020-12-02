package godruid

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupby(t *testing.T) {
	t.Parallel()
	patterns := map[string]struct {
		input    *QueryGroupBy
		expected string
	}{
		"success": {
			input: &QueryGroupBy{
				DataSource:  "campaign",
				Intervals:   []string{"2014-09-01T00:00/2020-01-01T00"},
				Granularity: GranAll,
				Filter:      FilterAnd(FilterJavaScript("hour", "function(x) { return(x >= 1) }"), nil),
				LimitSpec:   LimitDefault(5),
				Dimensions:  []DimSpec{"campaign_id"},
				Aggregations: []Aggregation{
					*AggRawJson(`{ "type" : "count", "name" : "count" }`),
					*AggLongSum("impressions", "impressions"),
				},
				PostAggregations: []PostAggregation{
					PostAggArithmetic(
						"imp/count", "/",
						[]PostAggregation{
							PostAggFieldAccessor("impressions"),
							PostAggRawJson(`{ "type" : "fieldAccess", "fieldName" : "count" }`)})},
			},
		},
	}
	for msg, p := range patterns {
		t.Run(msg, func(t *testing.T) {
			client := Client{
				Url:        "http://localhost:8082",
				HttpClient: &http.Client{},
				Debug:      true,
			}
			err := client.Query(p.input, "")
			fmt.Printf("requst: %s", client.LastRequest)
			assert.NoError(t, err)
			fmt.Printf("response: %s", client.LastResponse)
			fmt.Printf("query.SubmitResult: %v", p.input.QueryResult)
		})
	}
}

func TestSearch(t *testing.T) {
	t.Parallel()
	patterns := map[string]struct {
		input    *QuerySearch
		expected string
	}{
		"success": {
			input: &QuerySearch{
				DataSource:       "campaign",
				Intervals:        []string{"2014-09-01T00:00/2020-01-01T00"},
				Granularity:      GranAll,
				SearchDimensions: []string{"campaign_id", "hour"},
				Query:            SearchQueryInsensitiveContains(1313),
				Sort:             SearchSortLexicographic,
			},
		},
	}
	for msg, p := range patterns {
		t.Run(msg, func(t *testing.T) {
			client := Client{
				Url:        "http://localhost:8082",
				HttpClient: &http.Client{},
				Debug:      true,
			}
			err := client.Query(p.input, "")
			fmt.Printf("requst: %s", client.LastRequest)
			assert.NoError(t, err)
			fmt.Printf("response: %s", client.LastResponse)
			fmt.Printf("query.SubmitResult: %v", p.input.QueryResult)
		})
	}
}
