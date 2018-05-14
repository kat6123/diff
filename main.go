package main

import (
	"fmt"
	"github.com/Kat6123/diff/lcs"
)

func main() {
	file1, err := lcs.ReadFile("file1")
	if err != nil {
		return
	}
	file2, err := lcs.ReadFile("file2")
	if err != nil {
		return
	}

	//X := "MXMJYAUZ"
	//Y := "MMZJAWXU"

	C := lcs.LCS(file1, file2)
	lenC := len(C)
	for i := 0; i < lenC; i++ {
		fmt.Println(C[i])
	}
	fmt.Println(lcs.BuildLCS(C, file1, file2))
	fmt.Println(lcs.BuildDiff(C, file1, file2))
}
