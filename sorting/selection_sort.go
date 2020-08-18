package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minInd := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minInd] {
				minInd = j
			}
		}

		arr[minInd], arr[i] = arr[i], arr[minInd]
	}
}

func main() {
	arr := []int{6, 4, 2, 1, 2, 3, 5, 2}

	selectionSort(arr)

	fmt.Println(arr)
}
