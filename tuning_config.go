package godruid

type TuningConfig interface{}

type tuningConfigKafka struct {
	Type             string `json:"type"`
	MaxRowPerSegment int    `json:"maxRowsPerSegment"`
}

func TuningConfigKafka(maxRowPerSegment int) TuningConfig {
	return &tuningConfigKafka{
		Type:             "kafka",
		MaxRowPerSegment: maxRowPerSegment,
	}
}
