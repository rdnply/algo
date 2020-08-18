package main

import (
	"fmt"

	"./stack"
)

func deleteMiddle(st *stack.Stack, n int, curr int) {
	if st.Empty() || curr == n {
		return
	}

	temp := st.Pop()
	deleteMiddle(st, n, curr+1)

	if curr != n/2 {
		st.Push(temp)
	}
}

func main() {
	st := stack.New()

	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)

	fmt.Println(st)

	deleteMiddle(&st, len(st), 0)

	fmt.Println(st)
}
