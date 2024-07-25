package linkedChain

import (
	"github.com/polina/grammar/pkg/containers/linkedList"
)

const bucketSize = 32

type Bucket[T any] struct {
	values [bucketSize]T
	start  int
	end    int
}

type LinkedChain[T any] struct {
	buckets *linkedList.LinkedList[*Bucket[T]]
}

func NewLinkedChain[T any]() *LinkedChain[T] {
	return &LinkedChain[T]{}
}

func NewBucket[T any](val T) *Bucket[T] {
	b := Bucket[T]{
		start: 0,
		end:   0,
	}
	b.values[0] = val
	b.end++
	return &b
}

func (b *Bucket[T]) Add(val T) {
	b.values[b.end] = val
	b.end++
}

func (b *Bucket[T]) Pop() T {
	val := b.values[b.start]
	b.start++

	return val
}

func (b *Bucket[T]) IsFull() bool {
	return b.end >= bucketSize
}

func (b *Bucket[T]) IsEmpty() bool {
	return b.start >= b.end
}

func (ch *LinkedChain[T]) AddLast(val T) {
	if ch.buckets == nil {
		ch.buckets = &linkedList.LinkedList[*Bucket[T]]{}
		ch.buckets.AddLast(NewBucket(val))
		return
	}
	lastBucket := ch.buckets.GetLast()
	if lastBucket.IsFull() {
		lastBucket = NewBucket(val)
		ch.buckets.AddLast(lastBucket)
		return
	}
	lastBucket.Add(val)
	return
}

func (ch *LinkedChain[T]) PopFirst() T {
	firstBucket := ch.buckets.GetFirst()

	if ch.buckets.IsEmpty() {
		panic("empty list")
	}

	if firstBucket.IsEmpty() {
		ch.buckets.PopFirst()
		firstBucket = ch.buckets.GetFirst()
	}
	poppedValue := firstBucket.Pop()
	return poppedValue
}
