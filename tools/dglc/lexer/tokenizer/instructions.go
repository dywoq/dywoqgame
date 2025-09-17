package tokenizer

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/context"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
)

func MovInstruction(c context.Context) (*token.Token, error) {
	startPos := c.Position()
	char := c.Peek()
	if !isLetter(char) {
		return nil, nil
	}
	for isLetter(c.Peek()) {
		c.Advance()
	}
	literal, err := c.Slice(startPos, c.Position())
	if err != nil {
		return nil, err
	}
	if literal != "mov" {
		c.SetPosition(startPos)
		return nil, fmt.Errorf("expected mov instruction")
	}
	return &token.Token{Type: token.InstructionMov, Literal: literal, Line: c.Line(), Column: c.Column()}, nil
}
