package entry

import "github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"

// Build builds the entry game program.
func Build(c *config.Config) error {
	return command("go", Dir(c), "build", ".")
}
