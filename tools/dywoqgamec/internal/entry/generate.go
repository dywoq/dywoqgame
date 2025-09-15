package entry

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dywoq/dywoqgame/tools/dywoqgamec/internal/config"
)

// Generate creates a cache folder .dywoqgame, and generates files
// inside the folder named c.Title.  
func Generate(c *config.Config) error {
	dir := fmt.Sprintf(".dywoqgame/%s", c.Title)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	m := contentModuleFile(c)
	err = file(filepath.Join(dir, "go.mod"), []byte(m))
	if err != nil {
		return err
	}

	t := contentEntry()
	err = file(filepath.Join(dir, "main.go"), []byte(t))
	if err != nil {
		return err
	}
	return nil
}

func contentEntry() string {
	var b strings.Builder
	b.WriteString("package main\n")
	b.WriteString("\n")
	b.WriteString("import (\n")
	b.WriteString("   \"github.com/dywoq/dywoqgame/engine/core\n")
	b.WriteString(")\n")
	b.WriteString("\n")
	b.WriteString("func main() {\n")
	b.WriteString("   c := core.Game{}\n")
	b.WriteString("   if err := c.Loop(); err != nil {\n")
	b.WriteString("      panic(err)\n")
	b.WriteString("   }\n")
	b.WriteString("}\n")
	b.WriteString("\n")
	return b.String()
}

func contentModuleFile(c *config.Config) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("module %s\n", c.Title))
	b.WriteString("\n")
	b.WriteString("go 1.25.0\n")
	b.WriteString("\n")
	return b.String()
}

func file(path string, content []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	return err
}
