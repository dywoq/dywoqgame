package context

import (
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
	"github.com/dywoq/dywoqgame/tools/dglc/parser/expression"
)

type Peek interface {
	// Current returns the current token that is processed by parser.
	// It can return nil if there's no processing tokens.
	Current() *token.Token

	// Past returns the past expression that was already processed.
	// It can return nil if current position is zero.
	Past() *expression.Expression

	// Next returns the next token.
	// It can return nil if the parser is not working or the parser reached the end.
	Next() *token.Token
}

type Context interface {
	Peek

	// Advance advances the parser to the next token, skipping the current token.
	// Returns false if the parser is not processing or the parser reached the end.
	Advance() bool
}
