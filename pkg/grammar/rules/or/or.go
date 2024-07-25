package or

import (
	"github.com/polina/grammar/pkg/grammar/rules"
)

type Or struct {
	rules []rules.Rule
}

func (or Or) Match(tokens []string) rules.ParseResult {
	for _, rule := range or.rules {
		res := rule.Match(tokens)

		if res != nil {
			return res
		}
	}
	return nil
}

func NewOr(rules []rules.Rule) rules.Rule {
	return Or{rules: rules}
}
