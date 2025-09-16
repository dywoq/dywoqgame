package tokenizer

import (
	"errors"
	"regexp"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/context"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
)

var (
	identifierRe = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*")
)

func Identifier(c context.Context) (*token.Token, error) {
	match := identifierRe.FindString(c.Input()[c.Position():])
	if match == "" {
		return nil, errors.New("not a valid identifier")
	}
	c.AdvanceBy(len(match))
	return &token.Token{Type: token.Identifier, Literal: match}, nil
}
