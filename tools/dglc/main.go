package main

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/tokenizer"
)

func main() {
	l := lexer.New("# sdsdsd\n  something", []tokenizer.Func{
		tokenizer.Number,
		tokenizer.String,
		tokenizer.Identifier,
	})
	result, err := l.Tokenize()
	if err != nil {
		panic(err)
	}
	for _, t := range result {
		if t != nil {
			fmt.Println(*t)
		}
	}
}
