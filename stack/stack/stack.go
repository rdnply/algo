package stack

type Stack []int

func New() Stack {
	return make([]int, 0)
}

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	if s.Empty() {
		return -1
	}

	lastInd := len(*s) - 1

	popped := (*s)[lastInd]
	(*s) = (*s)[:lastInd]

	return popped
}

func (s Stack) Top() int {
	if s.Empty() {
		return -1
	}

	return s[len(s)-1]
}

func (s Stack) Size() int {
	return len(s)
}
