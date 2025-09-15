package args

import "flag"

type Args struct {
	Config string // -config in CLI
}

// NewArgs creates new Args instance.
func NewArgs() *Args {
	a := &Args{}
	flag.StringVar(&a.Config, "config", "./config.json", "Config path")
	return a
}
