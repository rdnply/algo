package main

import (
	"fmt"

	"./queue"
)

func reverseQueue(q *queue.Queue) {
	if q.Empty() {
		return
	}

	temp := q.Dequeue()
	reverseQueue(q)
	q.Enqueue(temp)
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
