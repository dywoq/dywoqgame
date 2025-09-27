package parser

import (
	"slices"

	"github.com/dywoq/dywoqgame/interpreter/token"
)

// Parser is responsible parsing tokens into the AST nodes.
type Parser struct {
	currentTokens []*token.Token
	pos           int
	miniParsers   []MiniParser
	setupOn       bool
}

func NewParser() *Parser {
	return &Parser{currentTokens: []*token.Token{}, pos: 0, miniParsers: []MiniParser{}, setupOn: false}
}

// setups setups the default mini parsers,
// and sets p.setupOn to true.
func (p *Parser) setup() {
	if !p.setupOn {
		p.setupOn = true
		p.miniParsers = []MiniParser{
			ParseInstruction,
		}
	}
}

// reset resets the current position,
// and updates the current tokens to the new given ones.
func (p *Parser) reset(tokens []*token.Token) {
	p.pos = 0
	if !slices.Equal(p.currentTokens, tokens) {
		p.currentTokens = tokens
	}
}

// Current returns the current token.
// Returns nil if the current parser position reached EOF token.
func (p *Parser) Current() *token.Token {
	if p.Eof() {
		return nil
	}
	return p.currentTokens[p.pos]
}

// Peek returns the future character,
// if the current parser position+1 will be greater than the length of the tokens,
// the function will return nil.
func (p *Parser) Peek() *token.Token {
	switch {
	case p.Eof():
		return nil
	case p.pos+1 >= len(p.currentTokens):
		return nil
	}
	return p.currentTokens[p.pos+1]
}

// Advance goes to the next token by n.
// If the current parser position+n will be greater than the length of the tokens,
// or the parser position reached EOF token, the function will return nil.
func (p *Parser) Advance(n int) {
	if len(p.currentTokens) == 0 {
		return
	}
	p.pos += n
	if p.pos >= len(p.currentTokens) {
		p.pos = len(p.currentTokens) - 1
	}
}

// Eof reports whether the parser reached EOF token.
func (p *Parser) Eof() bool {
	return p.currentTokens[p.pos].Kind == token.KIND_EOF
}

func (p *Parser) Parse(tokens []*token.Token) ([]Node, error) {
	if len(tokens) == 0 {
		return nil, nil
	}
	p.reset(tokens)
	p.setup()
	result := []Node{}
	for !p.Eof() {
		tok := p.Current()
		matched := false
		for _, miniParser := range p.miniParsers {
			parsed, err := miniParser(p, tok)
			if err == ErrNoMatch {
				continue
			}
			if err != nil {
				return nil, err
			}
			result = append(result, parsed)
			matched = true
		}
		if !matched {
			p.Advance(1)
		}
	}
	return result, nil
}
