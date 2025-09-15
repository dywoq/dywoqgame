package main

import (
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/args"
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/entry"
)

func main() {
	a := args.NewArgs()
	c := &config.Config{}
	switch a.Command {
	case "build":
		err := entry.Setup(c, a)
		if err != nil {
			panic(err)
		}

	case "run":
		err := entry.Setup(c, a)
		if err != nil {
			panic(err)
		}
		err = entry.Run(c)
		if err != nil {
			panic(err)
		}
	}
}
