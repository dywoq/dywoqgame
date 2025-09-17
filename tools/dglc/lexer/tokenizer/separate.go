package tokenizer

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/context"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
)

func Separate(c context.Context) (*token.Token, error) {
	ch := c.Current()
	var t token.Type
	switch ch {
	case ',':
		t = token.SeparatorComma
	case ';':
		t = token.SeparatorSemicolon
	case '(':
		t = token.SeparatorLParen
	case ')':
		t = token.SeparatorRParen
	case '{':
		t = token.SeparatorLBrace
	case '}':
		t = token.SeparatorRBrace
	default:
		return nil, fmt.Errorf("unknown separator: %q", ch)
	}
	c.Advance()
	return &token.Token{Type: t, Literal: string(ch), Line: c.Line(), Column: c.Column()}, nil
}
