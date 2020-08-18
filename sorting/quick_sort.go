package main

import "fmt"

func quickSort(arr []int, l, r int) {
	if l < r {
		pi := partition(arr, l, r)

		quickSort(arr, l, pi-1)
		quickSort(arr, pi+1, r)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}

func quickSortIter(arr []int, l, h int) {
	stack := make([]int, h-l+1)

	top := -1
	top++
	stack[top] = l
	top++
	stack[top] = h

	for top >= 0 {
		h = stack[top]
		top--
		l = stack[top]
		top--

		pi := partition(arr, l, h)

		if pi-1 > l {
			top++
			stack[top] = l
			top++
			stack[top] = pi - 1
		}

		if pi+1 < h {
			top++
			stack[top] = pi + 1
			top++
			stack[top] = h
		}
	}
}

func main() {
	arr := []int{6, 4, 7, 1, 2, 3, 0}

	// quickSort(arr, 0, len(arr)-1)
	quickSortIter(arr, 0, len(arr)-1)

	fmt.Println(arr)
}
