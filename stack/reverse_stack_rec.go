package main

import "fmt"

type Stack []int

func NewStack() Stack {
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

var stack = NewStack()

func insertAtBottom(x int) {
	if stack.Empty() {
		stack.Push(x)
	} else {
		top := stack.Pop()
		insertAtBottom(x)
		stack.Push(top)
	}
}

func reverseStack() {
	if !stack.Empty() {
		x := stack.Pop()
		reverseStack()
		insertAtBottom(x)
	}
}

func main() {
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	fmt.Println(stack)

	reverseStack()

	fmt.Println(stack)

}
