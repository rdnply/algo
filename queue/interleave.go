package main

import (
	"fmt"

	"../stack/stack"
	"./queue"
)

func interleave(q *queue.Queue) {
	st := stack.New()

	half := q.Size() / 2

	for i := 0; i < half; i++ {
		st.Push(q.Dequeue())
	}

	for !st.Empty() {
		q.Enqueue(st.Pop())
	}

	for i := 0; i < half; i++ {
		q.Enqueue(q.Dequeue())
	}

	for i := 0; i < half; i++ {
		st.Push(q.Dequeue())
	}

	for !st.Empty() {
		q.Enqueue(st.Pop())
		q.Enqueue(q.Dequeue())
	}
}

func main() {
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)

	fmt.Println(q)

	interleave(&q)

	fmt.Println(q)
}
