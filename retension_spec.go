package godruid

import (
	"encoding/json"
)

type RetentionRulesSpec interface {
	setup()
	GetRawJSON() []byte
	onResponse(content []byte) error
}

type RetentionRule struct {
	Type             string            `json:"type"`
	Period           string            `json:"period,omitempty"`
	IncludeFuture    bool              `json:"includeFuture,omitempty"`
	Interval         string            `json:"interval,omitempty"`
	TieredReplicants *TieredReplicants `json:"tieredReplicants,omitempty"`
}

type RetentionRules struct {
	Rules               []*RetentionRule
	RetentionRuleResult *RetentionRulesResult `json:"-"`
	RawJSON             []byte                `json:"-"`
}

type RetentionRulesResult struct {
	rules []*RetentionRule
}

type TieredReplicants struct {
	Hot         int `json:"hot,omitempty"`
	DefaultTier int `json:"_default_tier,omitempty"`
}

func (s *RetentionRules) setup()             {}
func (s *RetentionRules) GetRawJSON() []byte { return nil }
func (s *RetentionRules) onResponse(content []byte) error {
	res := new(RetentionRulesResult)
	err := json.Unmarshal(content, res)
	if err != nil {
		return err
	}
	s.RetentionRuleResult = res
	s.RawJSON = content
	return nil
}

func LoadForeverRetensionRule() *RetentionRule {
	return &RetentionRule{
		Type: "loadForever",
	}
}

func DropBeforeByPeriodRetensionRule(period string) *RetentionRule {
	return &RetentionRule{
		Type:   "dropBeforeByPeriod",
		Period: period,
	}
}
