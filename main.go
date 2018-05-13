package main

import (
	"fmt"
	"github.com/Kat6123/diff/lcs"
)


func main() {
	X := "XMJYAUZ"
	Y := "MZJAWXU"

	C := lcs.LCS(X, Y)
	lenC := len(C)
	for i := 0; i < lenC; i++ {
		fmt.Println(C[i])
	}
	fmt.Println(lcs.BuildLCS(C, X, Y))
}

