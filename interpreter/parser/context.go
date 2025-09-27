package parser

import "github.com/dywoq/dywoqgame/interpreter/token"

type Reader interface {
	// Current returns the current token.
	// Returns nil if the current parser position reached EOF token.
	Current() *token.Token

	// Peek returns the future character,
	// if the current parser position+1 will be greater than the length of the tokens,
	// the function will return nil.
	Peek() *token.Token
}

type Advancer interface {
	// Advance goes to the next token by n.
	// If the current parser position+n will be greater than the length of the tokens,
	// or the parser position reached EOF token, the function will return nil.
	Advance(n int)
}

type EofChecker interface {
	// Eof reports whether the parser reached EOF token.
	Eof() bool
}

type Context interface {
	Reader
	Advancer
	EofChecker
}
