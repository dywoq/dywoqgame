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
