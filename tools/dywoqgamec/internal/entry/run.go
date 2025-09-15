package entry

import (
	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
)

// Run runs the compiled executable.
func Run(c *config.Config) error {
	exec := module(c)
	return command(exec, Dir(c))
}
