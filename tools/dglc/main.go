package main

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/tokenizer"
)

func main() {
	src := "add int64 x, int64 y {\nmov result, x\nmov ret, result\n}"
	l := lexer.New(src, []tokenizer.Func{
		tokenizer.Separate,
		tokenizer.VarType,
		tokenizer.MovInstruction,

		tokenizer.Identifier,

		tokenizer.Number,
		tokenizer.String,
	})

	result, err := l.Tokenize()
	if err != nil {
		panic(err)
	}
	for _, t := range result {
		if t != nil {
			fmt.Println(t)
		}
	}
}
