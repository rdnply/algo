package main

import "fmt"

type node struct {
	data        int
	left, right *node
}

func NewNode(d int) *node {
	return &node{d, nil, nil}
}

func search(root *node, k int) bool {
	for root != nil {
		if k > root.data {
			root = root.right
		} else if k < root.data {
			root = root.left
		} else {
			return true
		}
	}

	return false
}

func insert(root *node, k int) *node {
	if root == nil {
		return NewNode(k)
	}

	if k > root.data {
		root.right = insert(root.right, k)
	} else if k < root.data {
		root.left = insert(root.left, k)
	}

	return root
}

func main() {
	root := NewNode(50)
	insert(root, 20)
	insert(root, 40)
	insert(root, 80)
	insert(root, 10)
	insert(root, 2)

	fmt.Println(search(root, 3))
}
