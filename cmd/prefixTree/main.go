package main

import (
	"bufio"
	"cmp"
	"compress/gzip"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/polina/grammar/pkg/prefixTree"
)

func main() {
	file, err := os.Open("/home/esantonov/data/query_freq.tsv.gz")
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

	queries := make([]*prefixTree.Query, 0, 40000000)

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

	slices.SortStableFunc(queries, func(a, b *prefixTree.Query) int {
		return cmp.Compare(prefixTree.Frequency(b), prefixTree.Frequency(a))
	})

	frequentQueries := queries[:9]
}
