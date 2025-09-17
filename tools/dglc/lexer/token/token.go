package token

import "fmt"

type Token struct {
	Type    Type
	Literal string
	Line    int
	Column  int
}

func (t *Token) String() string {
	return fmt.Sprintf("{type: %s, literal: \"%s\", line: %d, column: %d}", t.Type, t.Literal, t.Line, t.Column)
}
