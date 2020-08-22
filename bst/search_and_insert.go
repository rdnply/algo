package main

import "fmt"

type node struct {
	data        int
	left, right *node
}

func NewNode(d int) *node {
	return &node{d, nil, nil}
}

func insert(root *node, k int) *node {
	if root == nil {
		return NewNode(k)
	}

	if k < root.data {
		root.left = insert(root.left, k)
	} else {
		root.right = insert(root.right, k)
	}

	return root
}

func search(root *node, k int) *node {
	if root == nil || root.data == k {
		return root
	}

	if root.data < k {
		return search(root.right, k)
	}

	return search(root.left, k)
}

func inOrder(root *node) {
	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Print(root.data, " ")
	inOrder(root.right)
}

func minValueNode(root *node) *node {
	curr := root
	for curr != nil && curr.left != nil {
		curr = curr.left
	}

	return curr
}

func delete(root *node, k int) *node {
	if root == nil {
		return root
	}

	if root.data > k {
		root.left = delete(root.left, k)
	} else if root.data < k {
		root.right = delete(root.right, k)
	} else {
		if root.left == nil {
			temp := root.right
			root.right = nil
			return temp
		} else if root.right == nil {
			temp := root.left
			root.left = nil
			return temp
		}

		temp := minValueNode(root.right)

		root.data = temp.data

		root.right = delete(root.right, temp.data)
	}

	return root
}

func main() {
	root := NewNode(50)
	insert(root, 30)
	insert(root, 20)
	insert(root, 40)
	insert(root, 70)
	insert(root, 60)
	insert(root, 80)

	inOrder(root)
	fmt.Println(search(root, 90))

	delete(root, 50)
	inOrder(root)
}
