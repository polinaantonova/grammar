package sequence

import (
	"strings"
	"testing"

	"github.com/polina/grammar/pkg/grammar/rules"
	"github.com/polina/grammar/pkg/grammar/rules/dict"
	"github.com/polina/grammar/pkg/grammar/tree"
	"github.com/stretchr/testify/require"
)

func TestSequence_Match(t *testing.T) {
	v := dict.NewDict([]string{"washed"}, "v")
	n := dict.NewDict([]string{"mum", "window"}, "n")
	vp := NewSequence([]rules.Rule{v, n}, "vp")
	s := NewSequence([]rules.Rule{n, vp}, "s")

	t.Run("match", func(t *testing.T) {
		res := s.Match(strings.Fields("mum washed window"))
		require.Equal(t, rules.NewParseResult([]string{}, tree.Node{
			Name: "s",
			Children: []tree.Node{
				{
					Name:     "mum",
					Children: nil,
				},
				{
					Name: "vp",
					Children: []tree.Node{
						{Name: "washed"},
						{Name: "window"},
					},
				},
			},
		}), res)
	})

	t.Run("empty", func(t *testing.T) {
		res := s.Match([]string{})
		require.Nil(t, res)
	})

	t.Run("doesn't match", func(t *testing.T) {
		res := s.Match(strings.Fields("mum window washed"))
		require.Nil(t, res)
	})
}
