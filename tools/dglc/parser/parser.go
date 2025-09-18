package parser

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
	"github.com/dywoq/dywoqgame/tools/dglc/parser/expression"
	"github.com/dywoq/dywoqgame/tools/dglc/parser/parsers"
)

type Parser struct {
	Parsers []parsers.Func
	Tokens  []*token.Token

	ast []*expression.Expression
	pos int
}

func NewParser(tokens []*token.Token, parsers []parsers.Func) *Parser {
	p := &Parser{Tokens: tokens, Parsers: parsers}
	return p
}

func (p *Parser) Parse() ([]*expression.Expression, error) {
	p.pos = 0
	p.ast = nil
	for p.pos < len(p.Tokens) {
		if p.Tokens[p.pos].Type == token.Eof {
			break
		}
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		p.ast = append(p.ast, expr)
	}
	return p.ast, nil
}

func (p *Parser) Current() *token.Token {
	if p.pos >= len(p.Tokens) {
		return nil
	}
	return p.Tokens[p.pos]
}

func (p *Parser) Past() *expression.Expression {
	if p.pos == 0 {
		return nil
	}
	return p.ast[p.pos-1]
}

func (p *Parser) Next() *token.Token {
	if p.pos+1 >= len(p.Tokens) {
		return nil
	}
	return p.Tokens[p.pos+1]
}

func (p *Parser) Advance() bool {
	p.pos++
	return p.pos < len(p.Tokens)
}

func (p *Parser) parseExpression() (*expression.Expression, error) {
	for _, parser := range p.Parsers {
		expr, err := parser(p)
		if err == nil {
			return expr, nil
		}
		if err != parsers.ErrNoMatch {
			return nil, err
		}
	}
	return nil, fmt.Errorf("unexpected token during parsing: %v", p.Tokens[p.pos])
}
