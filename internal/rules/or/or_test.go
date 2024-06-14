package or

import (
	"strings"
	"testing"

	"github.com/polina/grammar/internal/rules"
	"github.com/polina/grammar/internal/rules/dict"
	"github.com/polina/grammar/internal/rules/sequence"
	"github.com/polina/grammar/internal/tree"
	"github.com/stretchr/testify/require"
)

func TestOr_Match(t *testing.T) {
	v := dict.NewDict([]string{"washed"})
	n := dict.NewDict([]string{"mum", "window"})
	vp := NewOr([]rules.Rule{sequence.NewSequence([]rules.Rule{v, n}, "vp"), v})

	t.Run("match v", func(t *testing.T) {
		res := vp.Match(strings.Fields("washed"))
		require.Equal(t, rules.NewParseResult([]string{}, tree.Node{Name: "washed"}), res)
	})

	t.Run("match v+n", func(t *testing.T) {
		res := vp.Match(strings.Fields("washed window"))
		require.Equal(t, rules.NewParseResult([]string{}, tree.Node{Name: "vp", Children: []tree.Node{
			{
				Name: "washed",
			},
			{
				Name: "window",
			},
		}}), res)
	})

	t.Run("empty", func(t *testing.T) {
		res := vp.Match([]string{})
		require.Nil(t, res)
	})

	t.Run("doesn't match", func(t *testing.T) {
		res := vp.Match(strings.Fields("mum washed"))
		require.Nil(t, res)
	})
}
