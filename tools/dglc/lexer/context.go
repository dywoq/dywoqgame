package lexer

type Context interface {
	// Input returns the input from the lexer.
	Input() string

	// Position returns the current position of the lexer.
	Position() int

	// Peek allows you to view the next character.
	Peek() byte

	// Advance advances the position to the next character.
	Advance()

	// Rewind moves the position back by -1, otherwise, if the position is zero,
	// the function does nothing.
	Rewind()

	// Slice allows you to select the value with the start and end of it.
	// The start will be immediately set to 0, not your start number, if the current position is zero.
	Slice(start, end int) string

	// Current returns the current character.
	Current() byte

	// Eof checks if the lexer reached the end of file.
	Eof() bool

	// Column returns the column of the current position.
	Column() int

	// Line returns the line of the current position.
	Line() int

	// Read returns the current character, and advances to the next position.
	Read() byte
}
