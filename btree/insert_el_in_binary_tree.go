package main

import "fmt"

type Queue []*node

func NewQueue() Queue {
	return make([]*node, 0)
}

func (q *Queue) Enqueue(x *node) {
	*q = append(*q, x)
}

func (q *Queue) Dequeue() *node {
	if q.Empty() {
		return nil
	}

	popped := (*q)[0]
	*q = (*q)[1:]

	return popped
}

func (q Queue) Empty() bool {
	return len(q) == 0
}

func (q Queue) Front() *node {
	if q.Empty() {
		return nil
	}

	return q[0]
}

func (q Queue) Rear() *node {
	if q.Empty() {
		return nil
	}

	return q[len(q)-1]
}

func (q Queue) Size() int {
	return len(q)
}

type node struct {
	data  int
	left  *node
	right *node
}

func NewNode(d int) *node {
	return &node{d, nil, nil}
}

func insertLevelOrder(root *node, d int) {
	newNode := NewNode(d)
	if root == nil {
		root = newNode
		return
	}

	q := NewQueue()
	q.Enqueue(root)

	for !q.Empty() {
		temp := q.Dequeue()

		if temp.left != nil {
			q.Enqueue(temp.left)
		} else {
			temp.left = newNode
			break
		}

		if temp.right != nil {
			q.Enqueue(temp.right)
		} else {
			temp.right = newNode
			break
		}
	}
}

func inOrder(root *node) {
	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Print(root.data, " ")
	inOrder(root.right)
}

func deleteDeepest(root *node, toDel *node) {
	q := NewQueue()
	q.Enqueue(root)

	for !q.Empty() {
		temp := q.Dequeue()

		if temp == toDel {
			temp = nil
			return
		}

		if temp.left == toDel {
			temp.left = nil
			return
		} else {
			q.Enqueue(temp.left)
		}

		if temp.right == toDel {
			temp.right = nil
			return
		} else {
			q.Enqueue(temp.right)
		}
	}
}

func delete(root *node, key int) {
	if root == nil {
		return
	}

	if root.left == nil && root.right == nil {
		if root.data == key {
			root = nil
		} else {
			return
		}
	}

	q := NewQueue()
	q.Enqueue(root)
	var temp, keyNode *node

	for !q.Empty() {
		temp = q.Dequeue()

		if temp.data == key {
			keyNode = temp
		}

		if temp.left != nil {
			q.Enqueue(temp.left)
		}

		if temp.right != nil {
			q.Enqueue(temp.right)
		}
	}

	if keyNode != nil {
		x := temp.data
		deleteDeepest(root, temp)
		keyNode.data = x
	}
}

func main() {
	root := NewNode(10)
	root.left = NewNode(11)
	root.left.left = NewNode(7)
	root.right = NewNode(9)
	root.right.left = NewNode(15)
	root.right.right = NewNode(8)

	inOrder(root)
	fmt.Println()
	insertLevelOrder(root, 12)
	inOrder(root)
	fmt.Println()
	delete(root, 12)
	inOrder(root)
}
