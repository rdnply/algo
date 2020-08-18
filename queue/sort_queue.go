package main

import (
	"fmt"

	"./queue"
)

const MaxInt = 1000000000

func findMinInd(q *queue.Queue, sortedInd int) int {
	minInd := -1
	minVal := MaxInt

	n := q.Size()

	for i := 0; i < n; i++ {
		if q.Front() < minVal && i <= sortedInd {
			minVal = q.Front()
			minInd = i
		}
		q.Enqueue(q.Dequeue())
	}

	return minInd
}

func addMinToRear(q *queue.Queue, minInd int) {
	n := q.Size()
	var minVal int

	for i := 0; i < n; i++ {
		if i != minInd {
			q.Enqueue(q.Dequeue())
		} else {
			minVal = q.Dequeue()
		}
	}

	q.Enqueue(minVal)
}

func sort(q *queue.Queue) {
	for i := 1; i <= q.Size(); i++ {
		minInd := findMinInd(q, q.Size()-i)
		addMinToRear(q, minInd)
	}
}

func main() {
	q := queue.New()
	q.Enqueue(5)
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(1)

	fmt.Println(q)

	sort(&q)

	fmt.Println(q)
}
