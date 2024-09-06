package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gertd/go-pluralize"

	"github.com/polina/grammar/pkg/grammar/rules"
	"github.com/polina/grammar/pkg/grammar/rules/check_full"
	"github.com/polina/grammar/pkg/grammar/rules/dict"
	"github.com/polina/grammar/pkg/grammar/rules/or"
	"github.com/polina/grammar/pkg/grammar/rules/placeholder"
	"github.com/polina/grammar/pkg/grammar/rules/regex"
	"github.com/polina/grammar/pkg/grammar/rules/sequence"
	"github.com/polina/grammar/pkg/grammar/tree"
)

func main() {
	vFin := dict.NewDict([]string{"washed", "said"}, "v")
	vInf := dict.NewDict([]string{"wash"}, "v")

	//nouns читаем из файла
	file, err := os.Open("pkg/enGram/nouns.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	nounScanner := bufio.NewScanner(file)
	nouns := make([]string, 0, 200000)

	//----------------добавляем мн ч сущ в словарь
	pluralize := pluralize.NewClient()

	for nounScanner.Scan() {
		noun := nounScanner.Text()
		nouns = append(nouns, noun)
		if pluralize.IsSingular(noun) {
			nouns = append(nouns, pluralize.Plural(noun))
		}
	}
	n := dict.NewDict(nouns, "n")

	adj := dict.NewDict([]string{"big"}, "adj")
	det := dict.NewDict([]string{"a", "the", "this", "that"}, "det")
	mod := dict.NewDict([]string{"can", "must", "has to"}, "v")
	num := regex.NewRegEx(`^(0|[1-9][0-9]*)$`, "num")

	np := or.NewOr([]rules.Rule{
		sequence.NewSequence([]rules.Rule{det, adj, n}, "np"),
		sequence.NewSequence([]rules.Rule{det, n}, "np"),
		sequence.NewSequence([]rules.Rule{num, adj, n}, "np"),
		sequence.NewSequence([]rules.Rule{num, n}, "np"),
	})

	s := placeholder.NewPlaceholder()

	vp := or.NewOr([]rules.Rule{
		sequence.NewSequence([]rules.Rule{vFin, s}, "vp"),
		sequence.NewSequence([]rules.Rule{mod, vInf, np}, "vp"),
		sequence.NewSequence([]rules.Rule{vFin, np}, "vp"),
		vFin,
	})

	s.SetRule(check_full.NewCheckFull(sequence.NewSequence([]rules.Rule{np, vp}, "s")))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		res := s.Match(tokens)

		if res == nil {
			fmt.Println("ungrammatical sentence")
			continue
		}

		fmt.Println(scanner.Text())
		tree.PrintNode(res.Node(), os.Stdout)
	}
}
