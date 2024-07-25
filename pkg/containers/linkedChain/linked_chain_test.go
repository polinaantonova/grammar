package linkedChain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/polina/grammar/pkg/containers/linkedList"
)

func TestLinkedChain_AddLast(t *testing.T) {
	t.Run("add to and pop from empty", func(t *testing.T) {
		ch := NewLinkedChain[int]()
		fillList(20, ch)
		res := dumpList(ch)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, res)
	})

	t.Run("refill", func(t *testing.T) {
		ch := NewLinkedChain[int]()
		fillList(17, ch)
		ch.PopFirst()
		fillList(3, ch)
		res := dumpList(ch)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 0, 1, 2}, res)
	})

	t.Run("pop from empty list", func(t *testing.T) {
		ch := NewLinkedChain[int]()
		require.Panics(t, func() {
			ch.PopFirst()
		})
	})
}

func fillList(count int, ch *LinkedChain[int]) {
	for i := 0; i < count; i++ {
		ch.AddLast(i)
	}
}

func dumpList(ch *LinkedChain[int]) []int {
	res := make([]int, 0, 3)
	for !ch.buckets.GetLast().IsEmpty() {
		res = append(res, ch.PopFirst())
	}
	return res
}

func BenchmarkLinkedChain_AddLast(b *testing.B) {
	b.Run("linked chain", func(b *testing.B) {
		ch := LinkedChain[int]{}
		for i := 0; i < b.N; i++ {
			ch.AddLast(i)
		}
		for i := 0; i < b.N; i++ {
			ch.PopFirst()
		}
	})
	b.Run("linked list", func(b *testing.B) {
		l := linkedList.LinkedList[int]{}
		for i := 0; i < b.N; i++ {
			l.AddLast(i)
		}
		for i := 0; i < b.N; i++ {
			l.PopFirst()
		}
	})
}
