package mergeSort

type Stack struct {
	slice []int
	size  int
}

func (s *Stack) Push(data int) {
	if s.size < len(s.slice) {
		s.slice[s.size] = data
	} else {
		s.slice = append(s.slice, data)
	}
	s.size++
}

func (s *Stack) Pop() int {
	res := s.slice[s.size-1]
	s.size--
	return res
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Slice() []int {
	return s.slice
}

func NewStack(cap int) *Stack {
	x := Stack{}
	x.slice = make([]int, cap)
	return &x
}
