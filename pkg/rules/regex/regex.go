package regex

import (
	"regexp"

	"github.com/polina/grammar/pkg/rules"
	"github.com/polina/grammar/pkg/tree"
)

type RegEx struct {
	regex string
	name  string
}

func (r RegEx) Match(tokens []string) rules.ParseResult {
	if len(tokens) == 0 {
		return nil
	}
	re := regexp.MustCompile(r.regex)

	if re.MatchString(tokens[0]) {
		return rules.NewParseResult(tokens[1:], tree.Node{
			Name: r.name,
			Children: []tree.Node{
				{Name: tokens[0]}}})
	}
	return nil
}

func NewRegEx(regex string, name string) rules.Rule {
	return RegEx{name: name, regex: regex}
}
