package godruid

type TuningConfig interface{}

type tuningConfigKafka struct {
	Type              string `json:"type"`
	MaxRowsPerSegment int    `json:"maxRowsPerSegment"`
	MaxRowsInMemory   int    `json:"maxRowsInMemory,omitempty"`
	MaxBytesInMemory  int    `json:"maxBytesInMemory,omitempty"`
}

func TuningConfigKafka(maxRowsPerSegment, maxRowsInMemory, maxBytesInMemory int) TuningConfig {
	return &tuningConfigKafka{
		Type:              "kafka",
		MaxRowsPerSegment: maxRowsPerSegment,
		MaxRowsInMemory:   maxRowsInMemory,
		MaxBytesInMemory:  maxBytesInMemory,
	}
}
