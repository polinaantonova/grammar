package suggest

import (
	"fmt"
	"github.com/polina/grammar/pkg/suggest/radixTree"
	"runtime"
)

func memoryCheck() {
	queries := ReadQueriesFile("/home/polina/data/query_freq.tsv.gz")

	//вычисляем по очереди при какой частоте переполнится оперативная память (5 ГБ)

	//node1 := prefixTree.NewNode()
	node2 := radixTree.NewNode()

	//для prefix tree частота 6
	//для radix tree частота 1
	count := 0
	i := 0
	for count <= len(queries) {
		for ; i < count+1000; i++ {
			//node1.Add(queries[i])
			node2.Add(queries[i])
			fmt.Println(queries[i].Frequency())
		}
		stats := &runtime.MemStats{}
		runtime.ReadMemStats(stats)
		fmt.Println("memory: ", stats.Alloc)
		if stats.Alloc > 5e+9 {
			break
		} else {
			count = count + 1000
		}
	}
}
