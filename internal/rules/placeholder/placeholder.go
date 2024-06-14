package placeholder

import (
	"github.com/polina/grammar/internal/rules"
)

type Placeholder struct {
	base rules.Rule
}

func (p *Placeholder) Match(tokens []string) rules.ParseResult {
	return p.base.Match(tokens)
}

func NewPlaceholder() *Placeholder {
	return &Placeholder{}
}

func (p *Placeholder) SetRule(r rules.Rule) {
	p.base = r
}
