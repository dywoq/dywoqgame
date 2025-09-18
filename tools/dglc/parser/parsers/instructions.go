package parsers

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
	"github.com/dywoq/dywoqgame/tools/dglc/parser/context"
	"github.com/dywoq/dywoqgame/tools/dglc/parser/expression"
)

func MakeInstruction(c context.Context) (*expression.Expression, error) {
	if c.Current().Literal != "make" {
		return nil, ErrNoMatch
	}
	c.Advance()

	varType := c.Current()
	if varType == nil || !token.IsVarType(varType.Type) {
		return nil, fmt.Errorf("expected a type name, got nothing or non-identifier: %v", varType)
	}
	c.Advance()

	varName := c.Current()
	if varName == nil || varName.Type != token.Identifier {
		return nil, fmt.Errorf("expected a variable name, got nothing or non-identifier: %v", varName.Literal)
	}
	c.Advance()

	if c.Current().Literal != "," {
		return nil, fmt.Errorf("expected a comma, got %q", c.Current().Literal)
	}
	c.Advance()

	val := c.Current()
	if val == nil {
		return nil, fmt.Errorf("expected a value, got nothing")
	}
	c.Advance()

	declExpr := expression.New(
		varType.Literal,
		expression.New(varName.Literal, nil, nil),
		expression.New(val.Literal, nil, nil),
	)
	return expression.New("make", declExpr, nil), nil
}

func ConstInstruction(c context.Context) (*expression.Expression, error) {
	if c.Current().Literal != "const" {
		return nil, ErrNoMatch
	}
	c.Advance()

	varType := c.Current()
	if varType == nil || !token.IsVarType(varType.Type) {
		return nil, fmt.Errorf("expected a type name, got nothing or non-identifier: %v", varType)
	}
	c.Advance()

	varName := c.Current()
	if varName == nil || varName.Type != token.Identifier {
		return nil, fmt.Errorf("expected a variable name, got nothing or non-identifier: %v", varName.Literal)
	}
	c.Advance()

	if c.Current().Literal != "," {
		return nil, fmt.Errorf("expected a comma, got %q", c.Current().Literal)
	}
	c.Advance()

	val := c.Current()
	if val == nil {
		return nil, fmt.Errorf("expected a value, got nothing")
	}
	c.Advance()

	declExpr := expression.New(
		varType.Literal,
		expression.New(varName.Literal, nil, nil),
		expression.New(val.Literal, nil, nil),
	)
	return expression.New("const", declExpr, nil), nil
}
