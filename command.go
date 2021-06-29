package cli

import "context"

const (
	CommandReturnCodeHelp = -1234567
)

// A Command is a runnable sub-command of a CLI
type Command interface {
	// Help returns detailed information about a command
	Help() string

	// Synopsis returns a one-line synopsis of the command
	// This should be less than 50 characters ideally.
	Synopsis() string

	// Run runs the actual command with the given CLI instance and
	// command-line arguments, returning an exit status
	Run(ctx context.Context, args []string) int
}

// A NamedCommand is a runnable sub-command of a CLI with a name
type NamedCommand interface {
	Command
	// Name returns the name of the command
	Name() string
}

// MockCommand is an implementation of Command that can be used for
// testing and for automatically populating missing parent commands.
type MockCommand struct {
	HelpText      string
	SynopsisText  string
	RunReturnCode int

	// Set by the command
	RunArgs   []string
	RunCalled bool
}

func (c *MockCommand) Help() string {
	return c.HelpText
}

func (c *MockCommand) Run(ctx context.Context, args []string) int {
	c.RunCalled = true
	c.RunArgs = args
	return c.RunReturnCode
}

func (c *MockCommand) Synopsis() string {
	return c.SynopsisText
}
