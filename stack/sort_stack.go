package main

import (
	"fmt"

	"./stack"
)

func sortStack(st *stack.Stack) {
	if !st.Empty() {
		temp := st.Pop()
		sortStack(st)
		sortedInsert(st, temp)
	}
}

func sortedInsert(st *stack.Stack, x int) {
	if st.Empty() || x > st.Top() {
		st.Push(x)
	} else {
		temp := st.Pop()
		sortedInsert(st, x)
		st.Push(temp)
	}
}

func main() {
	st := stack.New()
	st.Push(30)
	st.Push(-5)
	st.Push(18)
	st.Push(14)
	st.Push(-3)

	fmt.Println(st)

	sortStack(&st)

	fmt.Println(st)
}
