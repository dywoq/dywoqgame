package entry

import "fmt"

func debug(c *Core, args ...any) {
	if c.debug {
		fmt.Print(args)
	}
}

func debugf(c *Core, format string, args ...any) {
	if c.debug {
		fmt.Printf(format, args)
	}
}
