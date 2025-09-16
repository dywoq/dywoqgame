package panicking

var state any

// State returns the current panic state.
// If the panic state is not set, it would return nil and false.
func State() (any, bool) {
	if state == nil {
		return nil, false
	}
	return state, true
}

func (i *Implementation) Panic(v any) {
	if _, ok := State(); ok {
		return
	}
	state = v
}

func (i *Implementation) Recover() (any, bool) {
	val, ok := State()
	if ok {
		state = nil
		return val, ok
	}
	return nil, false
}
