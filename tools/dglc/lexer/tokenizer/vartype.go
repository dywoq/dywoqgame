package tokenizer

import (


	"github.com/dywoq/dywoqgame/tools/dglc/lexer/context"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
)

func VarType(c context.Context) (*token.Token, error) {
	startPos := c.Position()
	char := c.Peek()
	if !isLetter(char) {
		return nil, nil
	}
	for isLetter(c.Peek()) || isDigit(c.Peek()) {
		c.Advance()
	}
	literal, err := c.Slice(startPos, c.Position())
	if err != nil {
		return nil, err
	}
	tokenType := isKeyword(literal)
	if tokenType == token.Illegal {
		return nil, nil
	}
	return &token.Token{Type: tokenType, Literal: literal}, nil
}
