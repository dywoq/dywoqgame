package implementation


type Registrable interface {
	Register() error
}

type Registry struct {
	
}

func (reg *Registry) Add(r Registrable) {

}

func (reg *Registry) Init() error {
	return nil
}
