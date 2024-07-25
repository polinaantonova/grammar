package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		mySlice := []int{2, 3, 1}
		tree := BuildTree(mySlice)
		res := tree.Search(2)
		require.Equal(t, res, true)
	})

	t.Run("empty slice", func(t *testing.T) {
		var mySlice []int
		tree := BuildTree(mySlice)
		res := tree.Search(2)
		require.Equal(t, res, false)
	})

	t.Run("two child nodes -- ok", func(t *testing.T) {
		mySlice := []int{5, 4, 6, 3, 7}
		tree := BuildTree(mySlice)
		res := tree.Search(5)
		require.Equal(t, res, true)
	})

	t.Run("two child nodes -- no", func(t *testing.T) {
		mySlice := []int{5, 4, 6, 3, 7}
		tree := BuildTree(mySlice)
		res := tree.Search(10)
		require.Equal(t, res, false)
	})

	t.Run("only right nodes", func(t *testing.T) {
		mySlice := []int{1, 2, 3}
		tree := BuildTree(mySlice)
		res := tree.Search(2)
		require.True(t, res)
	})

	t.Run("left nodes", func(t *testing.T) {
		mySlice := []int{3, 2, 1}
		tree := BuildTree(mySlice)

		t.Run("found", func(t *testing.T) {
			res := tree.Search(1)
			require.Equal(t, res, true)
		})

		t.Run("not found", func(t *testing.T) {
			res := tree.Search(4)
			require.Equal(t, res, false)
		})
	})

}
