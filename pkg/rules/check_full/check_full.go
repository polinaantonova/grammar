package check_full

import (
	"github.com/polina/grammar/pkg/rules"
)

type CheckFull struct {
	rule rules.Rule
}

func (ch CheckFull) Match(tokens []string) rules.ParseResult {
	if len(tokens) == 0 {
		return nil
	}

	res := ch.rule.Match(tokens)
	if res == nil || len(res.Remain()) > 0 {
		return nil
	}

	return res
}

func NewCheckFull(rule rules.Rule) rules.Rule {
	return CheckFull{rule: rule}
}
