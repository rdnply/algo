package main

import "fmt"

type node struct {
	data        int
	left, right *node
}

func NewNode(d int) *node {
	return &node{d, nil, nil}
}

func findLevel(root *node, k, lvl int) int {
	if root == nil {
		return -1
	}

	if root.data == k {
		return lvl
	}

	left := findLevel(root.left, k, lvl+1)
	if left == -1 {
		return findLevel(root.right, k, lvl+1)
	}

	return left
}

func lca(root *node, n1, n2 int) *node {
	if root == nil {
		return nil
	}

	if root.data == n1 || root.data == n2 {
		return root
	}

	leftLCA := lca(root.left, n1, n2)
	rightLCA := lca(root.right, n1, n2)

	if leftLCA != nil && rightLCA != nil {
		return root
	}

	if leftLCA != nil {
		return lca(root.left, n1, n2)
	}

	return lca(root.right, n1, n2)
}

func calcDist(root *node, a, b int) int {
	lca := lca(root, a, b)

	d1 := findLevel(lca, a, 0)
	d2 := findLevel(lca, b, 0)

	return d1 + d2
}

func main() {
	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	root.right.left = NewNode(6)
	root.right.right = NewNode(7)
	root.right.left.right = NewNode(8)

	fmt.Println("dist(4, 5)", calcDist(root, 4, 5))
	fmt.Println("dist(8, 5)", calcDist(root, 8, 5))
}
