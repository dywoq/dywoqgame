package stream

// Type represents stream type.
type Type string

const (
	Error   Type = "error"
	Warning Type = "warning"
	Fatal   Type = "fatal"
	Out     Type = "out"
	Debug   Type = "debug"
)

// Write writes p to the t stream, returning amount of allocations.
// May return a error if t is wrong type.
//
// If stream data is out of maximum length, it removes the last message from the stream
// for your new message.
func Write(t Type, p []byte) (int, error) { return 0, nil } // stub

// Flush flushes the t stream and returns its data.
// May return a error if t is wrong type.
func Flush(t Type) ([]byte, error) { return []byte{}, nil } // stub

// Clean cleans the t stream without returning it.
// May return a error if t is wrong type.
func Clean(t Type) error { return nil } // stub

// Get gets the data from t stream.
// May return a error if t is wrong type
func Get(t Type) ([]byte, error) { return []byte{}, nil } // stub

// Capacity returns the capacity of the t stream data.
// May return a error if t is wrong type.
func Capacity(t Type) (int, error) { return 0, nil } // stub

// Length returns the length of the t stream data.
// May return a error if t is wrong type.
func Length(t Type) (int, error) { return 0, nil } // stub
