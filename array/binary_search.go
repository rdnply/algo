package main

import "fmt"

func binSearch(arr []int, x, l, r int) int {
	if r >= l {
		mid := (l + r) / 2

		if arr[mid] == x {
			return mid
		}
		if arr[mid] > x {
			return binSearch(arr, x, l, mid-1)
		}

		return binSearch(arr, x, mid+1, r)
	}

	return -1
}

func binarySearch(arr []int, x, l, r int) int {
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 6, 9, 13, 33}

	res := binSearch(arr, 14, 0, len(arr)-1)
	fmt.Println(res)
	res = binarySearch(arr, 13, 0, len(arr)-1)
	fmt.Println(res)
}
