package token

import "unicode"

// Kind represents the token kind.
type Kind string

// Map represents the map of tokens.
type Map map[string]Kind

// Position is a token position.
type Position struct {
	Line     int `json:"line"`
	Column   int `json:"column"`
	Position int `json:"position"`
}

type Token struct {
	Literal  string    `json:"literal"`
	Kind     Kind      `json:"kind"`
	Position *Position `json:"position"`
}

const (
	KIND_ILLEGAL          Kind = "illegal"
	KIND_KEYWORD          Kind = "keyword"
	KIND_TYPE             Kind = "type"
	KIND_SEPARATOR        Kind = "separator"
	KIND_IDENTIFIER       Kind = "identifier"
	KIND_FLOAT            Kind = "float"
	KIND_INTEGER          Kind = "integer"
	KIND_STRING           Kind = "string"
	KIND_BASE_INSTRUCTION Kind = "base_instruction"
	KIND_EOF              Kind = "eof"
)

var (
	Keywords = Map{
		"export":  KIND_KEYWORD,
		"module":  KIND_KEYWORD,
		"import":  KIND_KEYWORD,
		"nil":     KIND_KEYWORD,
		"declare": KIND_KEYWORD,
	}

	Separators = Map{
		",": KIND_SEPARATOR,
		"{": KIND_SEPARATOR,
		"}": KIND_SEPARATOR,
	}

	Types = Map{
		"str":  KIND_TYPE,
		"bool": KIND_TYPE,
		"i8":   KIND_TYPE,
		"i16":  KIND_TYPE,
		"i32":  KIND_TYPE,
		"i64":  KIND_TYPE,
		"u8":   KIND_TYPE,
		"u16":  KIND_TYPE,
		"u32":  KIND_TYPE,
		"u64":  KIND_TYPE,
		"void": KIND_TYPE,
	}

	BaseInstructions = Map{
		"add":   KIND_BASE_INSTRUCTION,
		"sub":   KIND_BASE_INSTRUCTION,
		"mul":   KIND_BASE_INSTRUCTION,
		"div":   KIND_BASE_INSTRUCTION,
		"write": KIND_BASE_INSTRUCTION,
		"store": KIND_BASE_INSTRUCTION,
	}
)

// IsIdentifier reports whether value is a valid identifier,
// meaning value can't be keyword, separator, type,
// contain hash, left and right paren, slash or start with the digit.
func IsIdentifier(value string) bool {
	switch {
	case Keywords.Is(value), Separators.Is(value), Types.Is(value):
		return false
	}
	for i, r := range value {
		// immediately check if first rune is digit
		if i == 0 && unicode.IsDigit(r) {
			return false
		}
		switch r {
		case '#':
			return false
		case '/':
			return false
		case '(':
			return false
		case ')':
			return false
		}
	}
	return true
}

// NewPosition returns a pointer to new token position.
func NewPosition(line, column, position int) *Position {
	return &Position{line, column, position}
}

// NewToken returns a pointer to new token.
func NewToken(literal string, kind Kind, position *Position) *Token {
	return &Token{Literal: literal, Kind: kind, Position: position}
}

// Equal reports whether x and y are equal by their literal and kind.
// If they're nil, it returns false.
func Equal(x *Token, y *Token) bool {
	if x == nil || y == nil {
		return false
	}
	if x.Literal == y.Literal && x.Kind == y.Kind {
		return true
	}
	return false
}

// Is reports whether the value exists in the tokens map.
func (m Map) Is(value string) bool {
	_, ok := m[value]
	return ok
}
