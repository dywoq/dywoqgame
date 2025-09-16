package globals

// Panic sets the panic state to v internally in the core of engine,
// and during the execution of game it will cause the immediate termination of the game.
func Panic(v ...any) {} // stub

// Recover returns the current panic state, otherwise, it would return zero value.
func Recover() any {
	var zero any
	return zero
} // stub
