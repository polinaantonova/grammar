package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLinkedList_AddLast(t *testing.T) {
	t.Run("add to empty list", func(t *testing.T) {
		l := LinkedList[int]{}
		fillList(3, &l)
		res := dumpList(&l)
		assert.Equal(t, LinkedList[int]{}, l)
		assert.Equal(t, []int{0, 1, 2}, res)
	})
	t.Run("refill list", func(t *testing.T) {
		l := LinkedList[int]{}
		fillList(2, &l)
		l.PopFirst()
		fillList(3, &l)
		res := dumpList(&l)
		require.Equal(t, []int{1, 0, 1, 2}, res)
	})
	t.Run("pop from empty list", func(t *testing.T) {
		l := LinkedList[int]{}
		require.Panics(t, func() {
			l.PopFirst()
		})
	})
}

func dumpList(l *LinkedList[int]) []int {
	res := make([]int, 0, 3)
	for !l.IsEmpty() {
		res = append(res, l.PopFirst())
	}
	return res
}

func fillList(count int, l *LinkedList[int]) {
	for i := 0; i < count; i++ {
		l.AddLast(i)
	}
}
