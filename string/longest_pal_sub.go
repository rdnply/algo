package main

import "fmt"

func longestPalSub(s string) string {
	start, maxLen := 0, 1
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ok := true

			for k := 0; k < (j-i+1)/2; k++ {
				if s[i+k] != s[j-k] {
					ok = false
					break
				}
			}

			if ok && (j-i+1) > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}

	return string(s[start : start+maxLen])
}

func longestPalindromSub(s string) string {
	start, maxLen := 0, 1
	n := len(s)

	table := make([][]bool, n)
	for i := range table {
		table[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		table[i][i] = true
	}

	for i := 0; i+1 < n; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = true
			maxLen = 2
			start = i
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {
			j := i + k - 1

			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true

				if k > maxLen {
					maxLen = k
					start = i
				}
			}
		}
	}

	return string(s[start : start+maxLen])
}

func lps(s string) string {
	var low, high int
	maxLen := 1
	start := 0
	n := len(s)

	for i := 1; i < n; i++ {
		low = i - 1
		high = i

		for low >= 0 && high < n && s[low] == s[high] {
			if (high - low + 1) > maxLen {
				maxLen = high - low + 1
				start = low
			}
			low--
			high++
		}

		low = i - 1
		high = i + 1

		for low >= 0 && high < n && s[low] == s[high] {
			if (high - low + 1) > maxLen {
				maxLen = high - low + 1
				start = low
			}
			low--
			high++
		}
	}

	return string(s[start : start+maxLen])
}

func main() {
	s := "ddabacc"

	fmt.Println(longestPalSub(s))
	fmt.Println(longestPalindromSub(s))
	fmt.Println(lps(s))
}
