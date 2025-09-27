package scanner

import (
	"github.com/dywoq/dywoqgame/interpreter/token"
)

type Reader interface {
	// Peek returns the future character,
	// or zero if there's no future character, or the scanner position
	// is out of bounds of the input.
	Peek() rune

	// Current returns the current character,
	// or zero if the scanner position
	// is out of bounds of the input.
	Current() rune
}

type PositionGetter interface {
	// Position returns a pointer to token.Position.
	Position() *token.Position
}

type Advancer interface {
	// Advances to the next position by n.
	// If the current position+n will be out of bounds of the input,
	// or the current position is out of bounds, it will do nothing.
	//
	// If the character is new line, the function will increase the current line and
	// reset the current column to 1.
	Advance(n int)
}

type Creator interface {
	// New creates a new pointer to token.Token and returns the pointer.
	// The only difference from token.NewToken is that scanner automatically
	// sets the position - making the code more short.
	New(literal string, kind token.Kind) *token.Token
}

type Slicer interface {
	// Slice returns the substring of the input with start and end.
	// If the start is negative, start is greater than end,
	// end is greater than the length of the input,
	// or the current input is empty,
	// sends an empty string and an error.
	Slice(start, end int) (string, error)
}

type EofChecker interface {
	// Eof reports whether the end of file (EOF) is reached by the scanner.	
	Eof() bool
}

// Context is a interface that allows you to manipulate scanner,
// such as advancing, manipulating, creating tokens, slicing, EOF checking and getting position.
type Context interface {
	Reader
	PositionGetter
	Advancer
	Creator
	Slicer
	EofChecker
}
