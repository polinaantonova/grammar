package dict

import (
	"github.com/polina/grammar/pkg/rules"
	"github.com/polina/grammar/pkg/tree"
)

type Dict struct {
	dict map[string]bool
	name string
}

func (d Dict) Match(tokens []string) rules.ParseResult {
	if len(tokens) == 0 {
		return nil
	}
	if d.dict[tokens[0]] {
		return rules.NewParseResult(tokens[1:], tree.Node{
			Name: d.name,
			Children: []tree.Node{
				{Name: tokens[0]}}})
	}
	return nil
}

func NewDict(words []string, name string) rules.Rule {
	newDict := make(map[string]bool, len(words))
	for _, word := range words {
		newDict[word] = true
	}

	return Dict{name: name, dict: newDict}
}
