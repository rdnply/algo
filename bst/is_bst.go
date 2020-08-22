package main

import "fmt"

type node struct {
	data        int
	left, right *node
}

func NewNode(d int) *node {
	return &node{d, nil, nil}
}

// func maxVal(root *node) int {
//     curr := root
//     for curr != nil && curr.right != nil {
//         curr = curr.right
//     }
//
//     return curr.data
// }
//
// func minVal(root *node) int {
//     curr := root
//     for curr != nil && curr.left != nil {
//         curr = curr.left
//     }
//
//     return curr.data
// }

// func isBST(root *node) bool {
//     if root == nil {
//         return false
//     }
//
//     if root.left != nil && maxVal(root.left) > root.data {
//         return false
//     }
//
//     if root.right != nil && minVal(root.right) < root.data {
//         return false
//     }
//
//     if !isBST(root.left) || !isBST(root.right) {
//         return false
//     }
//
//     return true
// }

const MaxInt = 1000000000
const MinInt = -MaxInt

func isBST(root *node) bool {
	return isBSTUtil(root, MinInt, MaxInt)
}

func isBSTUtil(root *node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.data < min || root.data > max {
		return false
	}

	return isBSTUtil(root.left, min, root.data-1) &&
		isBSTUtil(root.right, root.data+1, max)
}

func main() {
	root := NewNode(4)
	root.left = NewNode(2)
	root.right = NewNode(5)
	root.left.left = NewNode(1)
	root.left.right = NewNode(3)

	fmt.Println(isBST(root))

}
