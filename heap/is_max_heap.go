package main

import "fmt"

func isHeap(arr []int, i, n int) bool {
	if i > (n-2)/2 {
		return true
	}

	if arr[i] >= arr[2*i+1] && arr[i] >= arr[2*i+2] &&
		isHeap(arr, 2*i+1, n) && isHeap(arr, 2*i+2, n) {
		return true
	}

	return false
}

func isHeapIter(arr []int) bool {
	n := len(arr)

	for i := 0; i < (n-2)/2; i++ {
		if 2*i+1 < n && arr[2*i+1] > arr[i] {
			return false
		}

		if 2*i+2 < n && arr[2*i+2] > arr[i] {
			return false
		}
	}

	return true
}

func main() {
	arr := []int{90, 15, 10, 7, 12, 2, 7, 3}

	res := isHeapIter(arr)

	fmt.Println(res)
}
