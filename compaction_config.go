package godruid

import (
	"encoding/json"
)

type CompactionConfig struct {
	DataSource            string                  `json:"dataSource"`
	TaskPriority          int                     `json:"taskPriority,omitempty"`
	InputSegmentSizeBytes int                     `json:"inputSegmentSizeBytes,omitempty"`
	MaxRowsPerSegment     int                     `json:"maxRowsPerSegment,omitempty"`
	SkipOffsetFromLatest  string                  `json:"skipOffsetFromLatest,omitempty"`
	TuningConfig          *CompactionTuningConfig `json:"tuningConfig,omitempty"`

	CompactionConfigResult *CompactionConfigResult `json:"-"`
	RawJSON                []byte                  `json:"-"`
}

type CompactionTuningConfig struct {
	maxRowsInMemory  int `json:"maxRowsInMemory,omitempty"`
	maxBytesInMemory int `json:"maxBytesInMemory,omitempty"`
}

type CompactionConfigResult struct {
	CompactionConfigs       []*CompactionConfig `json:"compactionConfig"`
	CompactionTaskSlotRatio float32             `json:"compactionTaskSlotRatio"`
	MaxCompactionTaskSlots  int                 `json:"maxCompactionTaskSlots"`
}

func (c *CompactionConfig) setup()             {}
func (c *CompactionConfig) GetRawJSON() []byte { return nil }
func (c *CompactionConfig) onResponse(content []byte) error {
	res := new(CompactionConfigResult)
	err := json.Unmarshal(content, res)
	if err != nil {
		return err
	}
	c.CompactionConfigResult = res
	c.RawJSON = content
	return nil
}
