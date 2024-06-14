package check_full

import (
	"strings"
	"testing"

	"github.com/polina/grammar/internal/rules"
	"github.com/polina/grammar/internal/rules/dict"
	"github.com/polina/grammar/internal/rules/sequence"
	"github.com/polina/grammar/internal/tree"
	"github.com/stretchr/testify/require"
)

func TestCheckFull_Match(t *testing.T) {
	v := dict.NewDict([]string{"washed"})
	n := dict.NewDict([]string{"window"})
	vp := sequence.NewSequence([]rules.Rule{v, n}, "vp")
	vp = NewCheckFull(vp)

	t.Run("match", func(t *testing.T) {
		res := vp.Match(strings.Fields("washed window"))
		require.Equal(t, rules.NewParseResult([]string{}, tree.Node{
			Name: "vp",
			Children: []tree.Node{
				{Name: "washed"},
				{Name: "window"},
			},
		}), res)
	})

	t.Run("empty", func(t *testing.T) {
		res := vp.Match([]string{})
		require.Nil(t, res)
	})

	t.Run("doesn't match", func(t *testing.T) {
		res := vp.Match(strings.Fields("mum washed window door"))
		require.Nil(t, res)
	})

}
