package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubbleSortImpr(arr []int) {
	n := len(arr)
	swapped := false

	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if swapped == false {
			break
		}
	}
}

func bubbleSortRecur(arr []int, n int) {
	if n == 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	bubbleSortRecur(arr, n-1)
}

func main() {
	arr := []int{6, 4, 2, 5, 4, 2, 1, 1}

	// bubbleSort(arr)
	// bubbleSortImpr(arr)
	bubbleSortRecur(arr, len(arr))

	fmt.Println(arr)
}
