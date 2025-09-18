package main

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/tokenizer"
	"github.com/dywoq/dywoqgame/tools/dglc/parser"
	"github.com/dywoq/dywoqgame/tools/dglc/parser/parsers"
)

func main() {
	src := "make int32 x, 2\nconst int32 y, 5"
	l := lexer.New(src, []tokenizer.Func{
		tokenizer.Separate,
		tokenizer.VarType,
		tokenizer.Instruction,

		tokenizer.Identifier,

		tokenizer.Number,
		tokenizer.String,
	})

	result, err := l.Tokenize()
	if err != nil {
		panic(err)
	}

	p := parser.NewParser(result, []parsers.Func{
		parsers.MakeInstruction,
		parsers.ConstInstruction,
	})
	tree, err := p.Parse()
	if err != nil {
		panic(err)
	}
	for _, expression := range tree {
		if expression != nil {
			fmt.Printf("expression: %v\n", expression)
		}
	}
}
