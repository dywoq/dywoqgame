package lexer

type Lexer struct {
	input             string
	pos, line, column int
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
