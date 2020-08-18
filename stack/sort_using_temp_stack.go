package main

import (
	"fmt"

	"./stack"
)

func sortStack(st *stack.Stack) stack.Stack {
	tempStack := stack.New()

	for !st.Empty() {
		temp := st.Pop()

		for !tempStack.Empty() && tempStack.Top() > temp {
			st.Push(tempStack.Pop())
		}

		tempStack.Push(temp)
	}

	return tempStack
}

func main() {
	st := stack.New()
	st.Push(34)
	st.Push(3)
	st.Push(31)
	st.Push(98)
	st.Push(23)

	fmt.Println(st)

	st = sortStack(&st)

	fmt.Println(st)
}
