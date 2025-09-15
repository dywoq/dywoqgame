package entry

import (
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/args"
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
)

func Setup(c *config.Config, a *args.Args) error {
	err := c.Fill(a.Config)
	if err != nil {
		return err
	}
	err = Generate(c)
	if err != nil {
		return err
	}
	err = Install(c)
	if err != nil {
		return err
	}
	err = Build(c)
	if err != nil {
		return err
	}
	return nil
}
