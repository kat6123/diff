/*
Package lcs implements methods to create diff table and use it to build string slice difference.
*/
package lcs

// Table build graph
func Table(X []string, Y []string) [][]byte {
	m := len(X)
	n := len(Y)

	C := initArr(m+1, n+1)

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if X[i-1] == Y[j-1] {
				C[i][j] = C[i-1][j-1] + 1
			} else {
				C[i][j] = max(C[i-1][j], C[i][j-1])
			}
		}
	}

	return C
}

func initArr(m int, n int) [][]byte {
	C := make([][]byte, m)
	for i := 0; i < m; i++ {
		C[i] = make([]byte, n)
	}
	return C
}

func max(a byte, b byte) byte {
	if a >= b {
		return a
	}
	return b
}
