package main

import (
	"flag"
	"github.com/Kat6123/diff"
	"log"
)

var (
	unifiedPtr = flag.Bool("u", false, "output Diff in unified format")
	commonPtr  = flag.Bool("c", false, "output common part of two files")
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatal("two file paths are required")
		return
	}

	path1 := flag.Arg(0)
	path2 := flag.Arg(1)

	// TODO: transform to full path
	file1, err := diff.ReadFile(path1)
	if err != nil {
		return
	}
	file2, err := diff.ReadFile(path2)
	if err != nil {
		return
	}

	var result []string
	// Where I should define var diff here or at the top?
	// What about mutually exclusive flags
	switch {
	case *commonPtr:
		result = diff.Common(file1, file2)
	case *unifiedPtr:
		result = diff.Unified(file1, file2)
	default:
		result = diff.Normal(file1, file2)
	}

	diff.Print(result)
}
