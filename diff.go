package diff

import (
	"fmt"
	"github.com/Kat6123/diff/lcs"
)

// Construct return the largest common subsequence
func Common(X []string, Y []string) []string {
	// Is it a normal place to create a table?
	C := lcs.Table(X, Y)

	m := len(X)
	n := len(Y)

	lcsLastIndex := C[m][n]
	// uint8 or rune?
	var common = make([]string, lcsLastIndex)

	for m > 0 && n > 0 {
		if X[m-1] == Y[n-1] {
			common[lcsLastIndex-1] = X[m-1]
			lcsLastIndex--
			m--
			n--
		} else if C[m-1][n] >= C[m][n-1] {
			m--
		} else {
			n--
		}
	}

	return common
}

func Unified(X []string, Y []string) []string {
	C := lcs.Table(X, Y)

	m := len(X)
	n := len(Y)

	lenIndex := m + n - int(C[m][n])
	// uint8 or rune?
	var diffRes = make([]string, lenIndex)

	for index := lenIndex; index > 0; index-- {
		if m > 0 && n > 0 && X[m-1] == Y[n-1] {
			diffRes[index-1] = " " + X[m-1]
			m--
			n--
		} else if m > 0 && (n == 0 || C[m-1][n] >= C[m][n-1]) {
			diffRes[index-1] = "-" + X[m-1]
			m--
		} else if n > 0 && (m == 0 || C[m-1][n] < C[m][n-1]) {
			diffRes[index-1] = "+" + Y[n-1]
			n--
		} else {
			diffRes[index-1] = ""
		}
	}

	return diffRes
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

func (d *Diff) printDiff(X []string, Y []string) []string {
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

func buildChain(C [][]byte, X []string, Y []string) []Diff {
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

func Normal(X []string, Y []string) []string {
	var result []string
	C := lcs.Table(X, Y)

	chain := buildChain(C, X, Y)

	for _, ch := range chain {
		// What is ellipsis? unpacking?
		result = append(result, ch.printDiff(X, Y)...)
	}
	return result
}
