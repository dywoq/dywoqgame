package stream

import "errors"

var (
	ErrDataNil   = errors.New("github.com/dywoq/dywoqgame/engine/api/implementation/runtime/stream: data is nil")
	ErrDataFull  = errors.New("github.com/dywoq/dywoqgame/engine/api/implementation/runtime/stream: data is full")
	ErrWrongType = errors.New("github.com/dywoq/dywoqgame/engine/api/implementation/runtime/stream: wrong type")
)
