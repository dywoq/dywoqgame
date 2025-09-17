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
	tokenType := isVarTypeKeyword(literal)
	if tokenType == token.Illegal {
		return nil, nil
	}
	return &token.Token{Type: tokenType, Literal: literal, Line: c.Line(), Column: c.Column()}, nil
}

func isVarTypeKeyword(literal string) token.Type {
	switch literal {
	case "int8":
		return token.TypeSignedInteger8bit
	case "int16":
		return token.TypeSignedInteger16Bit
	case "int32":
		return token.TypeSignedInteger32bit
	case "int64":
		return token.TypeSignedInteger64Bit
	case "uint8":
		return token.TypeUnsignedInteger8bit
	case "uint16":
		return token.TypeUnsignedInteger16Bit
	case "uint32":
		return token.TypeUnsignedInteger32bit
	case "uint64":
		return token.TypeUnsignedInteger64Bit
	case "string":
		return token.TypeString
	default:
		return token.Illegal
	}
}
