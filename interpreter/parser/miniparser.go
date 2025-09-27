package parser

import (
	"github.com/dywoq/dywoqgame/interpreter/token"
)

// MiniParser is for parsing tokens into AST nodes.
// If it returns ErrNoMatch, it means the parser will try other mini-parsers.
type MiniParser func(c Context, t *token.Token) (Node, error)

// ParseInstruction parses the instruction into the AST node.
func ParseInstruction(c Context, t *token.Token) (Node, error) {
	if c.Current().Kind != token.KIND_BASE_INSTRUCTION && !token.IsIdentifier(c.Current().Literal) {
		return nil, ErrNoMatch
	}
	identifier := c.Current().Literal
	c.Advance(1)

	var args []InstructionArgumentStatement

	for !c.Eof() {
		current := c.Current()
		if current.Kind == token.KIND_SEPARATOR && current.Literal == "," {
			c.Advance(1)
			continue
		}

		args = append(args, InstructionArgumentStatement{
			Type:  current.Kind,
			Value: current.Literal,
		})
		c.Advance(1)
	}

	return InstructionStatement{
		Identifier: identifier,
		Arguments:  args,
	}, nil
}
