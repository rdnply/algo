package main

import "fmt"

func fix(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] != -1 && arr[i] != i {
			x := arr[i]

			for arr[x] != -1 && arr[x] != x {
				y := arr[x]
				arr[x] = x
				x = y
			}

			arr[x] = x

			if arr[i] != i {
				arr[i] = -1
			}
		}
	}

	return arr
}

func main() {
	arr := []int{-1, -1, 6, 1, 9, 3, 2, -1, 4, -1}
	arr = fix(arr)
	fmt.Println(arr)
}
