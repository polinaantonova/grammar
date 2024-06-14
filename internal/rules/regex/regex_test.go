package regex

import (
	"strings"
	"testing"

	"github.com/polina/grammar/internal/rules"
	"github.com/polina/grammar/internal/tree"
	"github.com/stretchr/testify/require"
)

func TestNumeral_Match(t *testing.T) {
	numeral := NewRegEx(`^(0|[1-9][0-9]*)$`, "num")

	t.Run("match", func(t *testing.T) {
		res := numeral.Match(strings.Fields("100500 windows"))
		require.Equal(t, rules.NewParseResult([]string{"windows"}, tree.Node{
			Name:     "num",
			Children: []tree.Node{{Name: "100500"}},
		}), res)
	})

	t.Run("empty", func(t *testing.T) {
		res := numeral.Match([]string{})
		require.Nil(t, res)
	})

	t.Run("single word", func(t *testing.T) {
		res := numeral.Match([]string{"1"})
		require.Equal(t, rules.NewParseResult([]string{}, tree.Node{
			Name:     "num",
			Children: []tree.Node{{Name: "1"}},
		}), res)
	})
}

////`^(0|[1-9][0-9]*)$`
