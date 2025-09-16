package stream

type Type string

const (
	ErrorType   Type = "error"
	WarningType Type = "warning"
	FatalType   Type = "fatal"
	OutType     Type = "out"
	DebugType   Type = "debug"
)
