package scanner

import (
	"errors"
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

	if c.Eof() || c.Peek() != '.' {
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
		if char == 0 {
			break
		}
		if char == '\\' {
			c.Advance(1)
			if c.Eof() {
				break
			}
			c.Advance(1)
			continue
		}

		if char == '"' {
			endPos := c.Position().Position + 1

			c.Advance(1)

			str, err := c.Slice(startPos, endPos)
			if err != nil {
				return nil, err
			}

			return c.New(str, token.KIND_STRING), nil
		}
		c.Advance(1)
	}
	return nil, errors.New("unterminated string literal")
}
