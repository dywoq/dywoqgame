package scanner

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/dywoq/dywoqgame/interpreter/token"
)

// Tokenizer is a alias for the functions,
// what turn the characters into the tokens.
//
// May return the error ErrNoMatch and nil instead of token if the character doesn't matches the tokenizer, then,
// scanner should try to use the other tokenizer instead.
type Tokenizer func(c Context, r rune) (*token.Token, error)

// TokenizeNumber tokenizes r into the number token.
// If there's a number after the base (23.60), the tokenizer will count it as a float.
func TokenizeNumber(c Context, r rune) (*token.Token, error) {
	if !unicode.IsDigit(r) {
		return nil, ErrNoMatch
	}
	c.Advance(1)

	start := c.Position().Position
	for !c.Eof() && unicode.IsDigit(c.Current()) {
		c.Advance(1)
	}

	if c.Eof() || c.Current() != '.' {
		str, err := c.Slice(start, c.Position().Position)
		if err != nil {
			return nil, err
		}
		return c.New(str, token.KIND_INTEGER), nil
	}

	c.Advance(1)

	for !c.Eof() && unicode.IsDigit(c.Current()) {
		c.Advance(1)
	}
	str, err := c.Slice(start, c.Position().Position)
	if err != nil {
		return nil, err
	}
	return c.New(str, token.KIND_FLOAT), nil
}

// TokenizeString tokenizers r into the string token.
// Returns an error if the string is not unterminated.
func TokenizeString(c Context, r rune) (*token.Token, error) {
	if r != '"' {
		return nil, ErrNoMatch
	}
	startPos := c.Position().Position
	c.Advance(1)

	for !c.Eof() {
		char := c.Current()
		if char == '"' {
			endPos := c.Position().Position + 1
			c.Advance(1)
			str, err := c.Slice(startPos, endPos)
			if err != nil {
				return nil, err
			}
			return c.New(str, token.KIND_STRING), nil
		}
		if char == '\\' {
			c.Advance(1)
			if c.Eof() {
				return nil, errors.New("unterminated escape sequence")
			}
		}
		c.Advance(1)
	}
	return nil, errors.New("unterminated string literal")
}

// TokenizeKeywords tokenizes r into the keyword token.
func TokenizeKeyword(c Context, r rune) (*token.Token, error) {
	if !unicode.IsLetter(r) {
		return nil, ErrNoMatch
	}
	startPos := c.Position().Position
	for unicode.IsLetter(c.Current()) {
		c.Advance(1)
	}
	str, err := c.Slice(startPos, c.Position().Position)
	if err != nil {
		return nil, err
	}
	if !token.Keywords.Is(str) {
		c.Position().Position = startPos
		return nil, ErrNoMatch
	}
	return c.New(str, token.KIND_KEYWORD), nil
}

// TokenizeSeparator tokenizes r into the separator token.
func TokenizeSeparator(c Context, r rune) (*token.Token, error) {
	if !token.Separators.Is(string(r)) {
		return nil, ErrNoMatch
	}
	c.Advance(1)
	return c.New(string(r), token.KIND_SEPARATOR), nil
}

// TokenizeType tokenizes r into the type token.
func TokenizeType(c Context, r rune) (*token.Token, error) {
	if !unicode.IsLetter(r) {
		return nil, ErrNoMatch
	}

	startPos := c.Position().Position

	for !c.Eof() && unicode.IsLetter(c.Current()) {
		c.Advance(1)
	}

	str, err := c.Slice(startPos, c.Position().Position)
	if err != nil {
		return nil, err
	}

	switch str {
	case "str", "bool", "void":
		return c.New(str, token.KIND_TYPE), nil
	case "i", "u":
		for !c.Eof() && unicode.IsDigit(c.Current()) {
			c.Advance(1)
		}
		str, err := c.Slice(startPos, c.Position().Position)
		if err != nil {
			return nil, err
		}
		if !token.Types.Is(str) {
			return nil, fmt.Errorf("wrong integer type: %s", str)
		}
		return c.New(str, token.KIND_TYPE), nil
	default:
		if token.Types.Is(str) {
			return c.New(str, token.KIND_TYPE), nil
		}
		c.Position().Position = startPos
		return nil, ErrNoMatch
	}
}

// TokenizeIdentifier tokenizes r into the identifier token.
func TokenizeIdentifier(c Context, r rune) (*token.Token, error) {
	if !unicode.IsLetter(r) && r != '_' {
		return nil, ErrNoMatch
	}
	start := c.Position().Position

	for !c.Eof() {
		c.Advance(1)
		if unicode.IsSpace(c.Current()) || token.Separators.Is(string(c.Current())) {
			break
		}
	}

	str, err := c.Slice(start, c.Position().Position)
	if err != nil {
		return nil, err
	}

	if !token.IsIdentifier(str) {
		return nil, fmt.Errorf("wrong identifier: %s", str)
	}

	return c.New(str, token.KIND_IDENTIFIER), nil
}

// TokenizeBaseInstruction tokenizes r into the base instruction.
func TokenizeBaseInstruction(c Context, r rune) (*token.Token, error) {
	if !unicode.IsLetter(r) {
		return nil, ErrNoMatch
	}
	start := c.Position().Position

	for !c.Eof() {
		c.Advance(1)
		if unicode.IsSpace(c.Current()) {
			break
		}
	}

	str, err := c.Slice(start, c.Position().Position)
	if err != nil {
		return nil, err
	}

	if !token.BaseInstructions.Is(str) {
		c.Position().Position = start
		return nil, ErrNoMatch
	}

	return c.New(str, token.KIND_BASE_INSTRUCTION), nil
}
