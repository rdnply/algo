package main

import (
	"fmt"

	"../stack/stack"
	"./queue"
)

func countSize(q queue.Queue) int {
	cnt := 0

	for !q.Empty() {
		q.Dequeue()
		cnt++
	}

	return cnt
}

func reverseFirstKElems(q *queue.Queue, k int) {
	n := countSize(*q)
	if q.Empty() || k > n {
		return
	}

	st := stack.New()
	for i := 0; i < k; i++ {
		st.Push(q.Dequeue())
	}

	for !st.Empty() {
		q.Enqueue(st.Pop())
	}

	for i := 0; i < n-k; i++ {
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

	fmt.Println(q)

	reverseFirstKElems(&q, 3)

	fmt.Println(q)
}
