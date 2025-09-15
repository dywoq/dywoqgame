package entry

import (
	"errors"
	"os"
	"os/exec"

	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
)

var pkgs = []string{
	"github.com/dywoq/dywoqgame/engine",
}

// Install instals any necessary packages required to build a entry program.
func Install(c *config.Config) error {
	if len(pkgs) == 0 {
		return errors.New("github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/entry: no packages to install")
	}
	for _, pkg := range pkgs {
		cmd := exec.Command("go", "get", pkg)
		cmd.Dir = Dir(c)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
