package quickSort

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("[1 2 3 4 5]", func(t *testing.T) {
		slice := []int{5, 3, 4, 2, 1}
		slice = QuickSort(slice)

		assert.Equal(t, slice, []int{1, 2, 3, 4, 5})
	})

	t.Run("empty", func(t *testing.T) {
		slice := []int{}
		slice = QuickSort(slice)

		require.Nil(t, slice)
	})
}
