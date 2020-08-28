package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func lcs(a, b string, n, m int) int {
//     if n == 0 || m == 0 {
//         return 0
//     }
//
//     if a[n-1] == b[m-1] {
//         return 1 + lcs(a, b, n-1, m-1)
//     }
//
//     return max(lcs(a, b, n-1, m), lcs(a, b, n, m-1))
// }

func lcs(a, b string) string {
	n, m := len(a), len(b)
	lcs := make([][]int, n+1)
	for i := range lcs {
		lcs[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 || j == 0 {
				lcs[i][j] = 0
			} else if a[i-1] == b[j-1] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}

		}
	}

	// construct lcs
	ind := lcs[n][m]
	str := make([]byte, ind)
	i, j := n, m

	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			str[ind-1] = a[i-1]
			i--
			j--
			ind--
		} else if lcs[i-1][j] > lcs[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(str)
}

func main() {
	a := "AGGTAB"
	b := "GXTXAYB"

	fmt.Println(lcs(a, b))

}
