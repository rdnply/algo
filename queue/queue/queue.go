package queue

type Queue []int

func New() Queue {
	return make([]int, 0)
}

func (q *Queue) Enqueue(x int) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() int {
	if q.Empty() {
		return -1
	}

	popped := (*q)[0]
	*q = (*q)[1:]

	return popped
}

func (q Queue) Empty() bool {
	return len(q) == 0
}

func (q Queue) Front() int {
	if q.Empty() {
		return -1
	}

	return q[0]
}

func (q Queue) Rear() int {
	if q.Empty() {
		return -1
	}

	return q[len(q)-1]
}

func (q Queue) Size() int {
	return len(q)
}
