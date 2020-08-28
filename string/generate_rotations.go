package main

import "fmt"

func printRotations(s string) {
	n := len(s)
	s += s

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(string(s[i+j]))
		}
		fmt.Println()
	}
}

func main() {
	s := "abc"
	printRotations(s)
}
