package context

import (
	"go/token"

	"github.com/dywoq/dywoqgame/tools/dglc/parser/expression"
)

type Converters interface {
	String(expr *expression.Expression) (string, bool)
	Int8(expr *expression.Expression) (int8, bool)
	Int16(expr *expression.Expression) (int16, bool)
	Int32(expr *expression.Expression) (int32, bool)
	Int64(expr *expression.Expression) (int64, bool)
	UInt8(expr *expression.Expression) (uint8, bool)
	UInt16(expr *expression.Expression) (uint16, bool)
	UInt32(expr *expression.Expression) (uint32, bool)
	UInt64(expr *expression.Expression) (uint64, bool)
}

type Peek interface {
	// Current returns the current token that is processed by parser.
	// It can return nil if there's no processing tokens.
	Current() *token.Token

	// Past returns the past expression that was already processed.
	// It can return nil if there's no processed expressions.
	Past() *expression.Expression

	// Next returns the next token.
	// It can return nil if the parser reached the end.
	Next() *token.Token
}

type Context interface {
	Converters
	Peek

	// Advance advances the parser to the next token, skipping the current token.
	// Returns false if there were no left tokens.
	Advance() bool
}
