package context

import "github.com/dywoq/dywoqgame/tools/dglc/parser/expression"

type Converters interface {
	String(expr *expression.Expression) (string, bool)
	Int8(expr *expression.Expression) (int8, bool)
	Int16(expr *expression.Expression) (int16, bool)
	Int32(expr *expression.Expression) (int32, bool)
	Int64(expr *expression.Expression) (int64, bool)
	UInt8(expr *expression.Expression) (uint8, bool)
	UInt16(expr *expression.Expression) (uint16, bool)
	UInt32(expr *expression.Expression) (uint32, bool)
	UInt64(expr *expression.Expression) (uint64, bool)
}

type Context interface {
	Converters
}
