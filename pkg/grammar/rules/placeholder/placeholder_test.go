package placeholder

import (
	"strings"
	"testing"

	"github.com/polina/grammar/pkg/grammar/rules"
	"github.com/polina/grammar/pkg/grammar/rules/dict"
	"github.com/polina/grammar/pkg/grammar/tree"
	"github.com/stretchr/testify/require"
)

func TestPlaceholder_Match(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		w := NewPlaceholder()
		w.SetRule(dict.NewDict([]string{"mum"}, "w"))

		t.Run("match", func(t *testing.T) {
			res := w.Match(strings.Fields("mum washed window"))
			require.Equal(t, rules.NewParseResult([]string{"washed", "window"}, tree.Node{Name: "mum"}), res)
		})

		t.Run("empty", func(t *testing.T) {
			res := w.Match([]string{})
			require.Nil(t, res)
		})
	})

	t.Run("panic", func(t *testing.T) {
		w := NewPlaceholder()
		require.Panics(t, func() {
			_ = w.Match([]string{})
		})
	})

}
