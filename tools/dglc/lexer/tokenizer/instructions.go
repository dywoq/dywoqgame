package tokenizer

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/context"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
)

var instructionsMap = map[string]token.Type{
	"mov":       token.InstructionMov,
	"make":      token.InstructionMake,
	"const":     token.InstructionConst,
	"add":       token.InstructionAdd,
	"minus":     token.InstructionMinus,
	"div":       token.InstructionDiv,
	"mul":       token.InstructionMul,
	"ret":       token.InstructionRet,
	"terminate": token.InstructionTerminate,
}

func Instruction(c context.Context) (*token.Token, error) {
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
	val, ok := instructionsMap[literal]
	if !ok {
		c.SetPosition(startPos)
		return nil, fmt.Errorf("didn't find the corresponding instruction: %s", literal)
	}
	return &token.Token{Type: val, Literal: literal, Line: c.Line(), Column: c.Column()}, nil
}
