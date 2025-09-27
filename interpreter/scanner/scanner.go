package scanner

import (
	"errors"

	"github.com/dywoq/dywoqgame/interpreter/token"
)

// Scanner is responsible for tokenizing DywoqGame Lang code.
type Scanner struct {
	currentInput string
	position     *token.Position
	tokenizers   []Tokenizer
	setupOn      bool
}

// NewScanner returns a new pointer to Scanner structure.
func NewScanner() *Scanner {
	return &Scanner{"", token.NewPosition(1, 1, 0), []Tokenizer{}, false}
}

// setup sets the default tokenizers.
// Skips if s.setupOn is true.
func (s *Scanner) setup() {
	if s.setupOn {
		return
	}
	s.tokenizers = []Tokenizer{
		TokenizeType,
		TokenizeKeyword,
		TokenizeSeparator,
		TokenizeBaseInstruction,
		TokenizeIdentifier,
		TokenizeNumber,
		TokenizeString,
	}
	s.setupOn = true
}

// resets resets the scanner position, and sets the current input to currentInput.
func (s *Scanner) reset(currentInput string) {
	s.position.Position = 0
	s.position.Line = 1
	s.position.Column = 1
	s.currentInput = currentInput
}

// Scan sets the current input to input, setups the tokenizers,
// and scans and tokenizes the input. Returns nil, nil if input is empty.
// Scan also skips whitespaces and comments.
//
// If scanning was successful, the function will return a slice of tokens with appended EOF,
// and nil.
func (s *Scanner) Scan(input string) ([]*token.Token, error) {
	if input == "" {
		return nil, nil
	}
	s.setup()
	s.reset(input)
	result := []*token.Token{}
	for !s.Eof() {
		r := s.Current()
		matched := false
		for _, tok := range s.tokenizers {
			tkn, err := tok(s, r)
			if err == ErrNoMatch {
				continue
			}
			if err != nil {
				return nil, err
			}
			result = append(result, tkn)
			matched = true
			break
		}
		if !matched {
			s.Advance(1)
		}
	}
	result = append(result, token.NewToken("", token.KIND_EOF, s.position))
	return result, nil
}

// tokenize tokenizes the current character into the token,
// if the current input is not empty.
func (s *Scanner) tokenize() (*token.Token, error) {
	if len(s.currentInput) == 0 {
		return nil, errors.New("the current input is empty")
	}
	for _, tokenizer := range s.tokenizers {
		t, err := tokenizer(s, s.Current())
		if err == nil {
			return t, nil
		}
		if err != ErrNoMatch {
			return nil, err
		}
	}
	return token.NewToken("", token.KIND_ILLEGAL, s.position), nil
}

// Eof reports whether the end of file (EOF) is reached by the scanner.
func (s *Scanner) Eof() bool {
	return s.position.Position < len(s.currentInput) == false
}

// Peek returns the future character,
// or zero if there's no future character, or the scanner position
// is out of bounds of the input.
func (s *Scanner) Peek() rune {
	switch {
	case s.position.Position+1 > len(s.currentInput), s.Eof():
		return 0
	default:
		return rune(s.currentInput[s.position.Position+1])
	}
}

// Current returns the current character,
// or zero if the scanner position
// is out of bounds of the input.
func (s *Scanner) Current() rune {
	switch {
	case s.Eof():
		return 0
	default:
		return rune(s.currentInput[s.position.Position])
	}
}

// Position returns a pointer to token.Position.
func (s *Scanner) Position() *token.Position {
	return s.position
}

// Advances to the next position by n.
// If the current position+n will be out of bounds of the input,
// or the current position is out of bounds, it will do nothing.
//
// If the character is new line, the function will increase the current line and
// reset the current column to 1.
func (s *Scanner) Advance(n int) {
	if s.Eof() || n <= 0 {
		return
	}
	s.position.Position += n
	ch := s.Current()
	if ch == '\n' {
		s.position.Line++
		s.position.Column = 1
	} else {
		s.position.Column++
	}
}

// New creates a new pointer to token.Token and returns the pointer.
// The only difference from token.NewToken is that scanner automatically
// sets the position - making the code more short.
func (s *Scanner) New(literal string, kind token.Kind) *token.Token {
	return token.NewToken(literal, kind, s.position)
}

// Slice returns the substring of the input with start and end.
// If the start is negative, start is greater than end,
// end is greater than the length of the input,
// or the current input is empty,
// sends an empty string and an error.
func (s *Scanner) Slice(start, end int) (string, error) {
	switch {
	case start < 0:
		return "", errors.New("start is less than zero")
	case end > s.position.Position:
		return "", errors.New("end is greater than the current position")
	case len(s.currentInput) == 0:
		return "", errors.New("the current input is empty")
	}
	return s.currentInput[start:end], nil
}
