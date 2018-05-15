package main

import (
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

	C := lcs.Table(file1, file2)
	//lenC := len(C)
	//for i := 0; i < lenC; i++ {
	//	fmt.Println(C[i])
	//}
	//fmt.Println(lcs.Construct(C, file1, file2))
	//
	//for _, s := range lcs.PrintDiff(C, file1, file2) {
	//	fmt.Println(s)
	//}
	//
	//fmt.Println(file1)
	//fmt.Println(file2)
	chain := lcs.DiffChain(C, file1, file2)
	for _, ch := range chain {
		lcs.Print(ch.PrintDiff(file1, file2))
	}
}
