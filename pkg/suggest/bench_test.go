package suggest

import (
	"github.com/polina/grammar/pkg/suggest/fastRadixTree"
	"github.com/polina/grammar/pkg/suggest/prefixTree"
	"github.com/polina/grammar/pkg/suggest/radixTree"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func Benchmark_Trees(b *testing.B) {
	queries := ReadQueriesFile("/home/polina/data/query_freq.tsv.gz")
	frequentQueries := make([]*prefixTree.QueryStat, 0, 2000000)
	const freqLimit = 6
	for _, query := range queries {
		if query.Frequency() <= freqLimit {
			break
		}
		frequentQueries = append(frequentQueries, query)
	}

	build := func(t Tree, name string) {
		start := time.Now()
		for _, query := range frequentQueries {
			t.Add(query)
		}
		b.Logf("building %s tree: %v", name, time.Now().Sub(start))
	}

	var wg sync.WaitGroup

	simple := prefixTree.NewNode()
	radix := radixTree.NewNode()
	fastRadix := fastRadixTree.NewNode()

	wg.Add(3)

	go func() {
		defer wg.Done()
		build(simple, "simple")
	}()

	go func() {
		defer wg.Done()
		build(radix, "radix")
	}()

	go func() {
		defer wg.Done()
		build(fastRadix, "fastRadix")
	}()

	wg.Wait()

	getRandomPrefix := func() string {
		randomQuery := []rune(frequentQueries[rand.Intn(len(frequentQueries))].Name())
		queryLen := len(randomQuery)
		randomSubstring := randomQuery[:rand.Intn(queryLen)]
		return string(randomSubstring)
	}

	b.ResetTimer()
	b.Run("search", func(b *testing.B) {
		b.Run("simple", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				simple.Search(getRandomPrefix())
			}
		})

		b.Run("radix", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				radix.Search(getRandomPrefix())
			}
		})

		b.Run("fastRadix", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fastRadix.Search(getRandomPrefix())
			}
		})
	})
}
