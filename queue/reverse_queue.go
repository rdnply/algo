package main

import (
	"fmt"

	"../stack/stack"

	"./queue"
)

func reverseQueue(q *queue.Queue) {
	st := stack.New()

	for !q.Empty() {
		st.Push(q.Dequeue())
	}

	for !st.Empty() {
		q.Enqueue(st.Pop())
	}
}

func main() {
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	fmt.Println(q)

	reverseQueue(&q)

	fmt.Println(q)
}
