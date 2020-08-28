package main

import "fmt"

func leftRotate(arr []int, d int) []int {
	n := len(arr)

	for i := 0; i < d; i++ {
		temp := arr[0]
		for i := 1; i < n; i++ {
			arr[i-1] = arr[i]
		}
		arr[n-1] = temp
	}

	return arr
}

// func gcd(a, b int) int {
//     if b == 0 {
//         return a
//     }
//
//     return gcd(b, a%b)
// }

func gcd(a, b int) int {
	for b != 0 {
		a %= b
		a, b = b, a
	}

	return a
}

func leftRotate2(arr []int, d int) []int {
	n := len(arr)
	d = d % n
	gcd := gcd(d, n)
	for i := 0; i < gcd; i++ {
		temp := arr[i]
		j := i

		for {
			k := j + d
			if k >= n {
				k -= n
			}

			if k == i {
				break
			}

			arr[j] = arr[k]
			j = k
		}

		arr[j] = temp
	}

	return arr
}

func reverse(arr []int, s, e int) []int {
	for i, j := s, e; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func leftRotate3(arr []int, d int) []int {
	if d == 0 {
		return nil
	}

	n := len(arr)
	arr = reverse(arr, 0, d-1)
	arr = reverse(arr, d, n-1)
	arr = reverse(arr, 0, n-1)

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	arr = leftRotate(arr, 2)
	fmt.Println(arr)
	arr = leftRotate2(arr, 3)
	fmt.Println(arr)
	arr = leftRotate3(arr, 3)
	fmt.Println(arr)
}
