/*
Package lcs implements methods to create diff table and use it to build string slice difference.
*/
package lcs

import (
	"fmt"
)

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

// Construct return the largest common subsequence
func Construct(C [][]byte, X []string, Y []string) []string {
	m := len(X)
	n := len(Y)

	lcsLastIndex := C[m][n]
	// uint8 or rune?
	var lcs = make([]string, lcsLastIndex)

	for m > 0 && n > 0 {
		if X[m-1] == Y[n-1] {
			lcs[lcsLastIndex-1] = X[m-1]
			lcsLastIndex--
			m--
			n--
		} else if C[m-1][n] >= C[m][n-1] {
			m--
		} else {
			n--
		}
	}

	return lcs
}

func PrintDiff(C [][]byte, X []string, Y []string) []string {
	m := len(X)
	n := len(Y)

	lenIndex := m + n - int(C[m][n])
	// uint8 or rune?
	var lcs = make([]string, lenIndex)

	for index := lenIndex; index > 0; index-- {
		if m > 0 && n > 0 && X[m-1] == Y[n-1] {
			lcs[index-1] = " " + X[m-1]
			m--
			n--
		} else if m > 0 && (n == 0 || C[m-1][n] >= C[m][n-1]) {
			lcs[index-1] = "-" + X[m-1]
			m--
		} else if n > 0 && (m == 0 || C[m-1][n] < C[m][n-1]) {
			lcs[index-1] = "+" + Y[n-1]
			n--
		} else {
			lcs[index-1] = ""
		}
	}

	return lcs
}

type Diff struct {
	// range
	first  [2]int
	second [2]int
}

//small or not?
func (d *Diff) initStartRange(f int, s int) {
	d.first[0] = f
	d.second[0] = s
}

func (d *Diff) initEndRange(f int, s int) {
	d.first[1] = f
	d.second[1] = s
}

func (d *Diff) PrintDiff(X []string, Y []string) []string {
	var diff []string
	var head string
	firstRange := false
	secondRange := false
	var mode string

	if d.first[1]-1-(d.first[0]+1) > 1 {
		firstRange = true
	}
	if d.second[1]-1-(d.second[0]+1) > 1 {
		secondRange = true
	}

	for i := d.first[0] + 1; i < d.first[1]; i++ {
		mode = "d"
		diff = append(diff, "< "+X[i-1])
	}

	if firstRange {
		diff = append(diff, "---")
	}

	for j := d.second[0] + 1; j < d.second[1]; j++ {
		if mode == "d" {
			mode = "c"
		} else {
			mode = "a"
		}
		diff = append(diff, "> "+Y[j-1])
	}

	if !firstRange {
		head = fmt.Sprintf("%d", d.first[0]+1)
	} else {
		head = fmt.Sprintf("%d,%d", d.first[0]+1, d.first[1]-1)
	}
	head = head + mode
	if !secondRange {
		head = fmt.Sprintf("%s%d", head, d.second[0]+1)
	} else {
		head = fmt.Sprintf("%s%d,%d", head, d.second[0]+1, d.second[1]-1)
	}

	return append([]string{head}, diff...)
}

func DiffChain(C [][]byte, X []string, Y []string) []Diff {
	m := len(X)
	n := len(Y)

	length := C[m][n] + 1
	// uint8 or rune?
	var chain = make([]Diff, length)

	// TODO: add comment
	// or add new struct for [2]int, start end?
	// when assign struct is it by value
	current := &chain[length-1]
	current.initEndRange(m+1, n+1)

	// TODO: m + n - int() is a step number add comment
	for index := m + n - int(C[m][n]); index > 0; index-- {
		if m > 0 && n > 0 && X[m-1] == Y[n-1] {
			current.initStartRange(m, n)
			length--
			current = &chain[length-1]
			current.initEndRange(m, n)
			m--
			n--
		} else if m > 0 && (n == 0 || C[m-1][n] >= C[m][n-1]) {
			m--
		} else if n > 0 && (m == 0 || C[m-1][n] < C[m][n-1]) {
			n--
		}
	}

	return chain
}
