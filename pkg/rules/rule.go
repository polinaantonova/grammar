package rules

import (
	"github.com/polina/grammar/pkg/tree"
)

type ParseResult interface {
	Remain() []string
	Node() tree.Node
}

type parseResult struct {
	remain []string
	node   tree.Node
}

func (p parseResult) Remain() []string {
	return p.remain
}

func (p parseResult) Node() tree.Node {
	return p.node
}

func NewParseResult(remain []string, node tree.Node) ParseResult {
	return &parseResult{
		remain: remain,
		node:   node,
	}
}

type Rule interface {
	Match(tokens []string) ParseResult
}
