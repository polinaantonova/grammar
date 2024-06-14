package dict

import (
	"strings"
	"testing"

	"github.com/polina/grammar/pkg/rules"
	"github.com/polina/grammar/pkg/tree"
	"github.com/stretchr/testify/require"
)

func TestDict_Match(t *testing.T) {
	newDict := NewDict([]string{"mum", "washed", "window"}, "test")

	t.Run("match", func(t *testing.T) {
		res := newDict.Match(strings.Fields("mum washed window"))
		require.Equal(t, rules.NewParseResult([]string{"washed", "window"}, tree.Node{
			Name:     "test",
			Children: []tree.Node{{Name: "mum"}}}), res)
	})

	t.Run("empty", func(t *testing.T) {
		res := newDict.Match([]string{})
		require.Nil(t, res)
	})

	t.Run("single word", func(t *testing.T) {
		res := newDict.Match([]string{"mum"})
		require.Equal(t, rules.NewParseResult([]string{}, tree.Node{
			Name:     "test",
			Children: []tree.Node{{Name: "mum"}},
		}), res)
	})
}
