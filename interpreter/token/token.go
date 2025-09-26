package token

import "unicode"

// Kind represents the token kind.
type Kind string

// Map represents the map of tokens.
type Map map[string]Kind

const (
	KIND_ILLEGAL    Kind = "illegal"
	KIND_KEYWORD    Kind = "keyword"
	KIND_TYPE       Kind = "kind"
	KIND_SEPARATOR  Kind = "separator"
	KIND_IDENTIFIER Kind = "identifier"
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
		case '#', '/', '(', ')':
			return false
		}
	}
	return true
}

// Is reports whether the value exists in the tokens map.
func (m Map) Is(value string) bool {
	_, ok := m[value]
	return ok
}
