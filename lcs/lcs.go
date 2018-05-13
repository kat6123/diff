package lcs

func LCS(X string, Y string) [][]byte {
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

func BuildLCS(C [][]byte, X string, Y string) string {
	m := len(X)
	n := len(Y)

	lcsLastIndex := C[m][n]
	// uint8 or rune?
	var lcs = make([]uint8, lcsLastIndex)

	for m > 0 && n > 0 {
		if X[m - 1] == Y[n - 1]	{
			lcs[lcsLastIndex - 1] = X[m - 1]
			lcsLastIndex--; m--; n--
		} else if C[m -1][n] >= C[m][n - 1]{
			m--
		} else {
			n--
		}
	}

	return string(lcs)
}