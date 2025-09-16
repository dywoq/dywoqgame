package lexer

import "errors"

var (
	ErrOutOfPosition    = errors.New("github.com/dywoq/dywoqgame/tools/dglc/lexer: out of position")
	ErrEndLessThanStart = errors.New("github.com/dywoq/dywoqgame/tools/dglc/lexer: end is less than start")
	ErrNegativeStart    = errors.New("github.com/dywoq/dywoqgame/tools/dglc/lexer: negative start")
)
