package globals

// Comparable defines Compare method to make the types comparable.
type Comparable[T any] interface {
	// Compare compares v to the value which implements Comparable interface.
	//
	// - Returns 0 if they're equal.
	//
	// - Returns 1 if v is greater than the value.
	//
	// - Returns -1 if v is less than the value.
	Compare(v T) int
}
