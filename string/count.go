package main

import "fmt"

const (
	OUT = iota
	IN
)

func countWords(s string) int {
	cw := 0
	state := OUT
	for _, c := range s {
		if c == '\n' || c == '\t' || c == ' ' {
			state = OUT
		} else if state == OUT {
			state = IN
			cw++
		}
	}

	return cw
}

func main() {
	s := "one two\nthree\tfour "
	cnt := countWords(s)
	fmt.Println(cnt)
}
