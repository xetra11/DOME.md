package main

import (
	"fmt"
	"github.com/xetra11/DOME.md/src/parser"
	"os"
)

func main() {
	args := os.Args[1:]
	fmt.Print(args)
	parser.Parse(args[0])
}
