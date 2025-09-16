package panicking

// Panic sets the panic state to v internally in the core of game engine,
// and during the execution of game it will cause the immediate termination of the game.
//
// Note that the function doesn't cleanup resources.
//
// The function does nothing if the panic set was already set.
func Panic(v any) {} // stub

// Recover returns the current panic state, otherwise,
// it would return nil and false.
func Recover() (any, bool) {
	var zero any
	return zero, false
} // stub
