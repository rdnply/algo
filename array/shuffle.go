package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(arr []int) []int {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	arr = shuffle(arr)
	fmt.Println(arr)
}
