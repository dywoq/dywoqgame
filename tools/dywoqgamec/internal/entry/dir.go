package entry

import (
	"fmt"

	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
)

func Dir(c *config.Config) string {
	return fmt.Sprintf(".dywoqgame/%s", c.Title)
}
