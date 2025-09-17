package lexer

import (
	"errors"
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/tokenizer"
)

type Lexer struct {
	Tokenizers        []tokenizer.Func
	input             string
	pos, line, column int
}

func New(input string, tokenizers []tokenizer.Func) *Lexer {
	l := &Lexer{Tokenizers: tokenizers, input: input}
	return l
}

func (l *Lexer) Input() string {
	return l.input
}

func (l *Lexer) Position() int {
	return l.pos
}

func (l *Lexer) Peek() byte {
	if l.Eof() {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) Advance() {
	if l.Eof() {
		return
	}
	char := l.input[l.pos]
	if char == '\n' {
		l.line++
		l.column = 1
	} else {
		l.column++
	}
	l.pos++
}

func (l *Lexer) AdvanceBy(n int) error {
	if l.Eof() {
		return ErrNegativeAdvance
	}
	for range n {
		if l.Eof() {
			return ErrEof
		}
		l.Advance()
	}
	return nil
}

func (l *Lexer) Slice(start, end int) (string, error) {
	switch {
	case start < 0:
		return "", ErrNegativeStart
	case end < start:
		return "", ErrEndLessThanStart
	case end > len(l.input):
		return "", ErrOutOfPosition
	}
	return l.input[start:end], nil
}

func (l *Lexer) Current() byte {
	return l.input[l.pos]
}

func (l *Lexer) Eof() bool {
	return l.pos >= len(l.input)
}

func (l *Lexer) Column() int {
	return l.column
}

func (l *Lexer) Line() int {
	return l.line
}

func (l *Lexer) Read() byte {
	if l.Eof() {
		return 0
	}
	char := l.input[l.pos]
	l.Advance()
	return char
}

func (l *Lexer) Tokenize() ([]*token.Token, error) {
	tokens := []*token.Token{}
	for {
		if l.Eof() {
			break
		}
		l.skipWhitespace()
		l.skipComments()
		foundToken := false
		for _, tokenizer := range l.Tokenizers {
			startPos := l.Position()
			tok, err := tokenizer(l)
			if err == nil && tok != nil {
				tokens = append(tokens, tok)
				foundToken = true
				break
			} else {
				err := l.SetPosition(startPos)
				if err != nil {
					return nil, err
				}
			}
		}
		if !foundToken {
			return nil, fmt.Errorf("unexpected token at line %d, column %d", l.Line(), l.Column())
		}
	}
	tokens = append(tokens, &token.Token{Type: token.Eof, Literal: ""})
	return tokens, nil
}

func (l *Lexer) SetPosition(pos int) error {
	if pos < 0 || pos > len(l.input) {
		return errors.New("position is out of bounds")
	}
	l.pos = 0
	l.line = 1
	l.column = 1
	for l.pos < pos {
		l.Advance()
	}
	return nil
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) {
		char := l.input[l.pos]
		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			l.Advance()
		} else {
			break
		}
	}
}

func (l *Lexer) skipComments() {
	if l.Peek() == '#' {
		for l.Peek() != '\n' && !l.Eof() {
			l.Advance()
		}
		if l.Peek() == '\n' {
			l.Advance()
		}
	}
}
