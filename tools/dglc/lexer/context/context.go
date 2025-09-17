package context

type Context interface {
	// Input returns the input from the lexer.
	Input() string

	// Position returns the current position of the lexer.
	Position() int

	// Peek allows you to view the next character.
	Peek() byte

	// Advance advances the position to the next character.
	Advance()

	// AdvanceBy advances to the next position by n.
	AdvanceBy(n int) error

	// Slice allows you to select the value with the start and end of it.
	// The start will be immediately set to 0 if your start is negative.
	Slice(start, end int) (string, error)

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

	// SetPosition sets the current position to pos.
	SetPosition(pos int) error
}
