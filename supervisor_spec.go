package godruid

import (
	"encoding/json"
)

type SupervisorSpec interface {
	setup()
	GetRawJSON() []byte
	onResponse(content []byte) error
}

type SupervisorKafka struct {
	Type         string       `json:"type"`
	IOConfig     IOConfig     `json:"ioConfig"`
	TuningConfig TuningConfig `json:"tuningConfig"`
	DataSchema   *DataSchema  `json:"dataSchema"`
	SubmitResult Supervisor   `json:"-"`
	RawJSON      []byte       `json:"-"`
}

type Supervisor struct {
	Id string `json:"id"`
}

func (s *SupervisorKafka) setup()             {}
func (s *SupervisorKafka) GetRawJSON() []byte { return nil }
func (s *SupervisorKafka) onResponse(content []byte) error {
	res := new(Supervisor)
	err := json.Unmarshal(content, res)
	if err != nil {
		return err
	}
	s.SubmitResult = *res
	s.RawJSON = content
	return nil
}
