package main

import (
	"fmt"

	"github.com/dywoq/dywoqgame/interpreter/parser"
	"github.com/dywoq/dywoqgame/interpreter/scanner"
)

func main() {
	input := "tok 222.322, 3"

	s := scanner.NewScanner()
	tokens, err := s.Scan(input)
	if err != nil {
		panic(err)
	}

	p := parser.NewParser()
	tree, err := p.Parse(tokens)
	if err != nil {
		panic(err)
	}

	fmt.Printf("tree: %v\n", tree)
}
