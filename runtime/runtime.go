package runtime

import (
	"github.com/dywoq/dywoqgame/config"
	"github.com/dywoq/dywoqgame/resource"
)

// Runtime provides all information about the game,
// and ways to interact with it.
type Runtime interface {
	ResourceManager() resource.Management
	Metadata() config.Metadata
	Window() config.Window
}
