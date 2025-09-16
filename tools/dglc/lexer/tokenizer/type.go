package tokenizer

import (
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/context"
	"github.com/dywoq/dywoqgame/tools/dglc/lexer/token"
)

type Func func(c context.Context) (*token.Token, error)
