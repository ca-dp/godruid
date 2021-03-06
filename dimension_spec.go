package godruid

type DimSpec interface{}

type EmptyDimension struct{}

type Dimension struct {
	Type         string           `json:"type"`
	Dimension    string           `json:"dimension"`
	OutputName   string           `json:"outputName"`
	ExtractionFn *DimExtractionFn `json:"extractionFn,omitempty"`
}

type DimExtractionFn struct {
	Type        string       `json:"type"`
	Expr        string       `json:"expr,omitempty"`
	Query       *SearchQuery `json:"query,omitempty"`
	Format      string       `json:"format,omitempty"`
	Function    string       `json:"function,omitempty"`
	TimeZone    string       `json:"timeZone,omitempty"`
	Locale      string       `json:"locale,omitempty"`
	Granularity string       `json:"granularity,omitempty"`
	AsMillis    bool         `json:"asMillis,omitempty"`
}

type DimFiltered struct {
	Type     string  `json:"type"`
	Delegate DimSpec `json:"delegate"`
	Pattern  string  `json:"pattern,omitempty"`
}

type TimeExtractionDimensionSpec struct {
	Type               string       `json:"type"`
	Dimension          string       `json:"dimension"`
	OutputName         string       `json:"outputName"`
	ExtractionFunction ExtractionFn `json:"extractionFn"`
}

func DimDefault(dimension, outputName string) DimSpec {
	return &Dimension{
		Type:       "default",
		Dimension:  dimension,
		OutputName: outputName,
	}
}

func DimExtraction(dimension, outputName string, fn *DimExtractionFn) DimSpec {
	return &Dimension{
		Type:         "extraction",
		Dimension:    dimension,
		OutputName:   outputName,
		ExtractionFn: fn,
	}
}

func DimExFnRegex(expr string) *DimExtractionFn {
	return &DimExtractionFn{
		Type: "regex",
		Expr: expr,
	}
}

func DimExFnPartial(expr string) *DimExtractionFn {
	return &DimExtractionFn{
		Type: "partial",
		Expr: expr,
	}
}

func DimExFnSearchQuerySpec(query *SearchQuery) *DimExtractionFn {
	return &DimExtractionFn{
		Type:  "searchQuery",
		Query: query,
	}
}

func DimExFnTime(timeFormat, timeZone string, locale string, granularity string, asMillis bool) *DimExtractionFn {
	return &DimExtractionFn{
		Type:        "timeFormat",
		Format:      timeFormat,
		TimeZone:    timeZone,
		Locale:      locale,
		Granularity: granularity,
		AsMillis:    asMillis,
	}
}

func DimExFnJavascript(function string) *DimExtractionFn {
	return &DimExtractionFn{
		Type:     "javascript",
		Function: function,
	}
}

func DimFilteredRegex(delegate DimSpec, pattern string) *DimFiltered {
	return &DimFiltered{
		Type:     "regexFiltered",
		Delegate: delegate,
		Pattern:  pattern,
	}
}
