package main

import "fmt"

func reverseWords(s string) string {
	ind := len(s) - 1
	start, end := 0, ind+1

	var res []byte
	for ind >= 0 {
		if s[ind] == ' ' {
			start = ind + 1
			for start != end {
				res = append(res, s[start])
				start++
			}
			res = append(res, ' ')
			end = ind
		}
		ind--
	}

	start = 0
	for start != end {
		res = append(res, s[start])
		start++
	}

	return string(res)
}

func main() {
	s := "it is a string"
	fmt.Println(reverseWords(s))
}
