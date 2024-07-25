package binarySearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinarySearch3(t *testing.T) {
	t.Run("middle", func(t *testing.T) {
		mySlice := []int{1, 2, 3}
		res := BinarySearch3(2, mySlice)
		require.Equal(t, res, 1)
	})

	t.Run("less present", func(t *testing.T) {
		mySlice := []int{1, 2, 3, 4, 5}
		res := BinarySearch3(2, mySlice)
		require.Equal(t, res, 1)
	})

	t.Run("less not present", func(t *testing.T) {
		mySlice := []int{2, 3, 4}
		res := BinarySearch3(1, mySlice)
		require.Equal(t, res, 0)
	})

	t.Run("more", func(t *testing.T) {
		mySlice := []int{1, 2, 3, 4, 5}
		res := BinarySearch3(4, mySlice)
		require.Equal(t, res, 3)
	})

	t.Run("more not present", func(t *testing.T) {
		mySlice := []int{1, 2, 3}
		res := BinarySearch3(4, mySlice)
		require.Equal(t, res, 3)
	})

	t.Run("all slice check", func(t *testing.T) {
		mySlice := []int{1, 3, 5, 7}
		t.Run("0", func(t *testing.T) {
			res := BinarySearch3(0, mySlice)
			require.Equal(t, res, 0)
		})
		t.Run("1", func(t *testing.T) {
			res := BinarySearch3(1, mySlice)
			require.Equal(t, res, 0)
		})
		t.Run("2", func(t *testing.T) {
			res := BinarySearch3(2, mySlice)
			require.Equal(t, res, 1)
		})
		t.Run("3", func(t *testing.T) {
			res := BinarySearch3(3, mySlice)
			require.Equal(t, res, 1)
		})
		t.Run("4", func(t *testing.T) {
			res := BinarySearch3(4, mySlice)
			require.Equal(t, res, 2)
		})
		t.Run("5", func(t *testing.T) {
			res := BinarySearch3(5, mySlice)
			require.Equal(t, res, 2)
		})
		t.Run("6", func(t *testing.T) {
			res := BinarySearch3(6, mySlice)
			require.Equal(t, res, 3)
		})
		t.Run("7", func(t *testing.T) {
			res := BinarySearch3(7, mySlice)
			require.Equal(t, res, 3)
		})
		t.Run("8", func(t *testing.T) {
			res := BinarySearch3(8, mySlice)
			require.Equal(t, res, 4)
		})
		t.Run("100", func(t *testing.T) {
			res := BinarySearch3(100, mySlice)
			require.Equal(t, res, 4)
		})

	})

}
