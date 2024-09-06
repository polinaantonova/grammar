package suggest

import "github.com/polina/grammar/pkg/suggest/prefixTree"

type Tree interface {
	Add(query *prefixTree.QueryStat)
	Search(prefix string) []*prefixTree.QueryStat
}
