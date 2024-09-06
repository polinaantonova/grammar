package suggest

import (
	"github.com/polina/grammar/pkg/suggest/prefixTree"
	"github.com/polina/grammar/pkg/suggest/radixTree"
)

func BuildRadixTree() *radixTree.Node {
	queries := ReadQueriesFile("/home/polina/data/query_freq.tsv.gz")
	frequentQueries := make([]*prefixTree.QueryStat, 0, 2000000)
	const freqLimit = 6
	for _, query := range queries {
		if query.Frequency() <= freqLimit {
			break
		}
		frequentQueries = append(frequentQueries, query)
	}

	radix := radixTree.NewNode()
	for _, query := range frequentQueries {
		radix.Add(query)
	}

	return radix
}
