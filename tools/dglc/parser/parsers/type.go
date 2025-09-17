package parsers

import (
	"github.com/dywoq/dywoqgame/tools/dglc/parser/context"
	"github.com/dywoq/dywoqgame/tools/dglc/parser/expression"
)

type Func func(c context.Context) (*expression.Expression, error)
