package panicking

import (
	"reflect"

	"github.com/traefik/yaegi/interp"
)

type Implementation struct{}

func (i *Implementation) Register(in *interp.Interpreter) error {
	return in.Use(interp.Exports{
		"github.com/dywoq/dywoqgame/api/runtime/panicking": {
			"Panic":   reflect.ValueOf(i.Panic),
			"Recover": reflect.ValueOf(i.Recover),
		},
	})
}
