package command

import (
	"context"
	"strings"

	cli "github.com/edufschmidt/go-cli"
)

// FooCommand :
type FooCommand struct {
	UI cli.UI

	Command
}

// Name :
func (c *FooCommand) Name() string {
	return "foo"
}

// Synopsis :
func (c *FooCommand) Synopsis() string {
	return "Run the foo command"
}

// Run :
func (c *FooCommand) Run(ctx context.Context, args []string) int {

	c.UI.Output("Running foo()...")

	return 0
}

// Help :
func (c *FooCommand) Help() string {
	h := `
Usage: app foo [options] [args]

  Run the foo command.

General Options:
` + GlobalOptions() + `
`
	return strings.TrimSpace(h)
}
