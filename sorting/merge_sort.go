package main

import "fmt"

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := (l + r) / 2

		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	a1 := make([]int, n1)
	a2 := make([]int, n2)

	for i := 0; i < n1; i++ {
		a1[i] = arr[l+i]
	}

	for i := 0; i < n2; i++ {
		a2[i] = arr[m+i+1]
	}

	var i, j int
	k := l

	for i < n1 && j < n2 {
		if a1[i] < a2[j] {
			arr[k] = a1[i]
			i++
		} else {
			arr[k] = a2[j]
			j++
		}

		k++
	}

	for i < n1 {
		arr[k] = a1[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = a2[j]
		j++
		k++
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func mergeSortIter(arr []int) {
	n := len(arr)

	for currSize := 1; currSize <= n-1; currSize = 2 * currSize {
		for leftStart := 0; leftStart < n-1; leftStart += 2 * currSize {
			mid := min(leftStart+currSize-1, n-1)
			rightStart := min(leftStart+2*currSize-1, n-1)

			merge(arr, leftStart, mid, rightStart)
		}
	}
}

func main() {
	arr := []int{5, 1, 5, 3, 0, 2, 4, 6}

	// mergeSort(arr, 0, len(arr)-1)
	mergeSortIter(arr)

	fmt.Println(arr)
}
