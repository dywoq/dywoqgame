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
		err := command("go", Dir(c), "get", pkg)
		if err != nil {
			return err
		}
	}
	err := command("go", Dir(c), "mod", "tidy")
	if err != nil {
		return err
	}
	return nil
}

func command(name string, dir string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	return cmd.Run()
}
