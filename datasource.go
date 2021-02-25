package godruid

type DataSource struct {
	Type  string      `json:"type"`
	Name  string      `json:"name,omitempty"`
	Query interface{} `json:"query,omitempty"`
}

func DataSourceQuery(query interface{}) *DataSource {
	return &DataSource{
		Type:  "query",
		Query: query,
	}
}

func DataSourceTable(name string) *DataSource {
	return &DataSource{
		Type: "table",
		Name: name,
	}
}
