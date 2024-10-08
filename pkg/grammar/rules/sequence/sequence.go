package sequence

import (
	"github.com/polina/grammar/pkg/grammar/rules"
	"github.com/polina/grammar/pkg/grammar/tree"
)

type Sequence struct {
	rules []rules.Rule
	name  string
}

func (s Sequence) Match(tokens []string) rules.ParseResult {
	remain := tokens
	children := make([]tree.Node, 0, len(tokens))
	for _, rule := range s.rules {
		res := rule.Match(remain)
		if res == nil {
			return nil
		}
		remain = res.Remain()
		children = append(children, res.Node())
	}

	return rules.NewParseResult(remain, tree.Node{
		Name:     s.name,
		Children: children,
	})
}

func NewSequence(rules []rules.Rule, name string) rules.Rule {
	return Sequence{rules: rules, name: name}
}
