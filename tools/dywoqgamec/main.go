package main

import (
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/args"
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/entry"
)

func main() {
	a := args.NewArgs()
	
	c := &config.Config{}
	err := c.Fill(a.Config)
	if err != nil {
		panic(err)
	}

	err = entry.Generate(c)
	if err != nil {
		panic(err)
	}

	err = entry.Install(c)
	if err != nil {
		panic(err)
	}

	err = entry.Build(c)
	if err != nil {
		panic(err)
	}
}
