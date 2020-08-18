package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func insertionSortRecur(arr []int, n int) {
	if n <= 1 {
		return
	}

	insertionSortRecur(arr, n-1)

	last := arr[n-1]
	j := n - 2

	for j >= 0 && last < arr[j] {
		arr[j+1] = arr[j]
		j--
	}

	arr[j+1] = last
}

func main() {
	arr := []int{5, 3, 1, 5, 2, 6, 0}

	// insertionSort(arr)
	insertionSortRecur(arr, len(arr))

	fmt.Println(arr)
}
