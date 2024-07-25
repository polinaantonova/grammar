package linkedList

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type LinkedList[T any] struct {
	first *Node[T]
	last  *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{value: data}
}
func (l *LinkedList[T]) AddLast(data T) {
	if l.first == nil {
		l.first = NewNode(data)
		l.last = l.first
		return
	}
	currentNode := l.last
	l.last = NewNode(data)
	currentNode.next = l.last
	l.last.prev = currentNode
}

func (l *LinkedList[T]) GetLast() T {
	return l.last.value
}

func (l *LinkedList[T]) GetFirst() T {
	return l.first.value
}

func (l *LinkedList[T]) PopFirst() T {
	poppedNode := l.first
	if poppedNode.next == nil {
		l.first = nil
		l.last = nil
		return poppedNode.value
	}
	l.first = poppedNode.next
	l.first.prev = nil
	return poppedNode.value
}

func (l *LinkedList[T]) IsEmpty() bool {
	if l.first == nil {
		return true
	}
	return false
}
