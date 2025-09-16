package stream

import (
	"reflect"

	"github.com/traefik/yaegi/interp"
)

type Implementation struct{}

var mapStreams = map[Type]*Stream{
	ErrorType:   Error,
	WarningType: Warning,
	FatalType:   Fatal,
	OutType:     Out,
	DebugType:   Debug,
}

func (i *Implementation) Write(t Type, p []byte) (int, error) {
	if !i.correctKind(t) {
		return 0, ErrWrongType
	}
	return mapStreams[t].Write(p)
}

func (i *Implementation) Flush(t Type) ([]byte, error) {
	if !i.correctKind(t) {
		return []byte{}, ErrWrongType
	}
	return mapStreams[t].Flush()
}

func (i *Implementation) Clean(t Type) error {
	if !i.correctKind(t) {
		return ErrWrongType
	}
	return mapStreams[t].Clean()
}

func (i *Implementation) Get(t Type) ([]byte, error) {
	if !i.correctKind(t) {
		return []byte{}, ErrWrongType
	}
	return mapStreams[t].Get()
}

func (i *Implementation) Capacity(t Type) (int, error) {
	if !i.correctKind(t) {
		return 0, ErrWrongType
	}
	return mapStreams[t].Capacity()
}

func (i *Implementation) Length(t Type) (int, error) {
	if !i.correctKind(t) {
		return 0, ErrWrongType
	}
	return mapStreams[t].Length()
}

func (i *Implementation) Register(in *interp.Interpreter) error {
	return in.Use(interp.Exports{
		"github.com/dywoq/dywoqgame/api/runtime/stream": {
			"Write":    reflect.ValueOf(i.Write),
			"Flush":    reflect.ValueOf(i.Flush),
			"Clean":    reflect.ValueOf(i.Clean),
			"Get":      reflect.ValueOf(i.Get),
			"Capacity": reflect.ValueOf(i.Capacity),
			"Length":   reflect.ValueOf(i.Length),
		},
	})
}

func (i *Implementation) correctKind(t Type) bool {
	_, ok := mapStreams[t]
	return ok
}
