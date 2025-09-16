package implementation

import "github.com/traefik/yaegi/interp"

type Registrable interface {
	Register(i *interp.Interpreter) error
}

type Registry struct {
	Interpreter *interp.Interpreter
	registry    []Registrable
}

func (reg *Registry) Add(r Registrable) {
	reg.registry = append(reg.registry, r)
}

func (reg *Registry) Init() error {
	for _, r := range reg.registry {
		err := r.Register(reg.Interpreter)
		if err != nil {
			return err
		}
	}
	return nil
}
