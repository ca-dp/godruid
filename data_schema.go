package godruid

type DataSchema struct {
	Datasource      string         `json:"dataSource"`
	GranularitySpec Granlarity     `json:"granularitySpec"`
	Parser          *Parser        `json:"parser,omitempty"`
	MetricsSpec     []*MetricsSpec `json:"metricsSpec,omitempty"`
}
