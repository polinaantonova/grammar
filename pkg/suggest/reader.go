package suggest

import (
	"bufio"
	"compress/gzip"
	"github.com/polina/grammar/pkg/suggest/prefixTree"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadQueriesFile(filename string) []*prefixTree.QueryStat {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileGz, err := gzip.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fileGz.Close()

	scanner := bufio.NewScanner(fileGz)

	queries := make([]*prefixTree.QueryStat, 0, 40000000)

	for scanner.Scan() {
		before, after, _ := strings.Cut(scanner.Text(), "\t")
		name := before
		freq, _ := strconv.Atoi(after)
		query := prefixTree.NewQuery(name, freq)
		queries = append(queries, query)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.SliceStable(queries, func(i, j int) bool {
		return queries[i].Frequency() > queries[j].Frequency()
	})
	return queries
}
