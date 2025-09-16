package tokenizer

import (
	"errors"
	"regexp"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/context"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
)

var (
	integerRe = regexp.MustCompile(`^[0-9]+`)
	stringRe  = regexp.MustCompile(`^"[^\"]*"`)
)

func Number(c context.Context) (*token.Token, error) {
	match := integerRe.FindString(c.Input()[c.Position():])
	if match == "" {
		return &token.Token{}, errors.New("not an integer")
	}
	c.AdvanceBy(len(match))
	return &token.Token{Type: token.Integer, Literal: match}, nil
}

func String(c context.Context) (*token.Token, error) {
	match := stringRe.FindString(c.Input()[c.Position():])
	if match == "" {
		return &token.Token{}, errors.New("not a valid string literal")
	}
	c.AdvanceBy(len(match))
	literal := match[1 : len(match)-1]
	return &token.Token{Type: token.String, Literal: literal}, nil
}
