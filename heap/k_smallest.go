package main

import (
	"fmt"

	"./heap"
)

func makeHeap(arr []int) heap.Heap {
	h := heap.New()
	for i := range arr {
		h.InsertKey(arr[i])
	}

	return h
}

func findKSmallest(arr []int, k int) []int {
	res := make([]int, 0)

	h := makeHeap(arr)
	for i := 0; i < k; i++ {
		min := h.ExtractMin()
		res = append(res, min)
	}

	return res
}

func main() {
	arr := []int{5, 1, 3, 5, 6, 7, 4, 2}

	res := findKSmallest(arr, 3)

	fmt.Println(res)
}
