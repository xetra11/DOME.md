package main

import (
	"github.com/xetra11/DOME.md/src/parser"
	"os"
)

func main() {
	args := os.Args[1:]
	parser.Parse(args[0])
}
