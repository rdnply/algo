package main

import "fmt"

type node struct {
	data                int
	left, right, parent *node
}

func insert(root *node, k int) *node {
	if root == nil {
		return &node{k, nil, nil, nil}
	}

	if k < root.data {
		root.left = insert(root.left, k)
		root.left.parent = root
	} else {
		root.right = insert(root.right, k)
		root.right.parent = root
	}

	return root
}

func depth(n *node) int {
	d := -1

	for n != nil {
		d++
		n = n.parent
	}

	return d
}

func LCA(n1 *node, n2 *node) *node {
	d1, d2 := depth(n1), depth(n2)

	diff := d1 - d2

	if d1 < 0 {
		n1, n2 = n2, n1
		diff = -diff
	}

	for diff > 0 {
		diff--
		n1 = n1.parent
	}

	for n1 != nil && n2 != nil {
		if n1 == n2 {
			return n1
		}

		n1 = n1.parent
		n2 = n2.parent
	}

	return nil
}

func main() {
	var root *node

	root = insert(root, 20)
	root = insert(root, 8)
	root = insert(root, 22)
	root = insert(root, 4)
	root = insert(root, 12)
	root = insert(root, 10)
	root = insert(root, 14)

	n1 := root.left.right.left
	n2 := root.right

	lca := LCA(n1, n2)
	fmt.Print(lca.data)
}
