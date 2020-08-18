package main

import "fmt"

func max(arr []int) int {
	mx := 0

	for _, el := range arr {
		if el > mx {
			mx = el
		}
	}

	return mx
}

func countingSort(arr []int) {
	mx := max(arr)

	count := make([]int, mx+1)
	output := make([]int, len(arr))

	for _, e := range arr {
		count[e]++
	}

	for i := 1; i <= mx; i++ {
		count[i] += count[i-1]
	}

	for i := 0; i <= mx; i++ {
		output[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}

	for i := range output {
		arr[i] = output[i]
	}
}

func main() {
	arr := []int{5, 1, 6, 4, 3, 2, 1}

	countingSort(arr)

	fmt.Println(arr)
}
