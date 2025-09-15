package args

import "flag"

type Args struct {
	Config  string // -config in CLI
	Command string // -command in CLI (build, run)
}

// NewArgs creates new Args instance.
func NewArgs() *Args {
	a := &Args{}
	flag.StringVar(&a.Config, "config", "./config.json", "Config path")
	flag.StringVar(&a.Command, "command", "build", "Command to run")
	return a
}
