package main

import "fmt"

type node struct {
	data        int
	left, right *node
}

func NewNode(d int) *node {
	return &node{d, nil, nil}
}

func findPath(root *node, path *[]int, k int) bool {
	if root == nil {
		return false
	}

	*path = append(*path, root.data)

	if root.data == k {
		return true
	}

	if (root.left != nil && findPath(root.left, path, k)) ||
		(root.right != nil && findPath(root.right, path, k)) {
		return true
	}

	*path = (*path)[:len(*path)-1]
	return false
}

func lca(root *node, n1, n2 int) int {
	var path1, path2 []int

	if !findPath(root, &path1, n1) || !findPath(root, &path2, n2) {
		return -1
	}

	var ind int
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] != path2[i] {
			ind = i
			break
		}
	}

	return path1[ind-1]
}

func findLCAUtil(root *node, n1, n2 int, v1, v2 *bool) *node {
	if root == nil {
		return nil
	}

	if root.data == n1 {
		*v1 = true
		return root
	}

	if root.data == n2 {
		*v2 = true
		return root
	}

	leftLCA := findLCAUtil(root.left, n1, n2, v1, v2)
	rightLCA := findLCAUtil(root.right, n1, n2, v1, v2)

	if leftLCA != nil && rightLCA != nil {
		return root
	}

	if leftLCA != nil {
		return leftLCA
	} else {
		return rightLCA
	}
}

func find(root *node, k int) bool {
	if root == nil {
		return false
	}

	if root.data == k || find(root.left, k) || find(root.right, k) {
		return true
	}

	return false
}

func findLCA(root *node, n1, n2 int) *node {
	var v1, v2 bool

	lca := findLCAUtil(root, n1, n2, &v1, &v2)

	if v1 && v2 || v1 && find(lca, n2) || v2 && find(lca, n1) {
		return lca
	}

	return nil
}

func main() {
	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	root.right.left = NewNode(6)
	root.right.right = NewNode(7)

	fmt.Println("lca(4,5)", lca(root, 4, 5))
	fmt.Println("lca(4,6)", lca(root, 4, 6))
	fmt.Println("lca(3,4)", lca(root, 3, 4))
	// fmt.Println("lca(2,4)", lca(root, 2, 4))

	fmt.Println("lca(4,5)", findLCA(root, 4, 5).data)
	fmt.Println("lca(4,6)", findLCA(root, 4, 6).data)
	fmt.Println("lca(3,4)", findLCA(root, 3, 4).data)
	fmt.Println("lca(2,4)", findLCA(root, 2, 4).data)
}
