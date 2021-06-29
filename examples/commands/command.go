package command

import (
	"fmt"
	"strings"

	cli "github.com/edufschmidt/go-cli"
	"github.com/spf13/pflag"
)

// Command is the base command
type Command struct {
	globalFlag string
}

// FlagSet declares flags that are common to all commands,
// returning a pflag.FlagSet struct that will hold their values after
// pflag.Parse() is called by the command.
func (c *Command) FlagSet(name string) *pflag.FlagSet {

	flags := pflag.NewFlagSet(name, pflag.ContinueOnError)

	flags.Usage = func() {}

	flags.StringVar(&c.globalFlag, "global-flag", "", "")

	// TODO: direct output to UI
	flags.SetOutput(nil)

	return flags
}

// GlobalOptions returns the global usage options string.
func GlobalOptions() string {
	text := `
  --global-flag=<string>
    A string flag.
    Default = ""
`
	return text
}

// DefaultErrorMessage returns the default error message for this command
func DefaultErrorMessage(cmd cli.NamedCommand) string {
	return fmt.Sprintf("For additional help try 'appname %s --help'", cmd.Name())
}

// manyStrings
type manyStrings []string

func (s *manyStrings) Set(val string) error {
	*s = append(*s, val)
	return nil
}

func (s *manyStrings) String() string {
	return strings.Join(*s, ",")
}
