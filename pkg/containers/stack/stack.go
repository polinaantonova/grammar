package stack

type Stack[T any] struct {
	slice []T
	size  int
}

func (s *Stack[T]) Push(data T) {
	if s.size < len(s.slice) {
		s.slice[s.size] = data
	} else {
		s.slice = append(s.slice, data)
	}
	s.size++
}

func (s *Stack[T]) Pop() T {
	res := s.slice[s.size-1]
	s.size--
	return res
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Slice() []T {
	return s.slice
}

func NewStack[T any](cap int) *Stack[T] {
	x := Stack[T]{}
	x.slice = make([]T, cap)
	return &x
}
