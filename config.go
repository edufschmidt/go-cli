package cli

import (
	"io"
	"os"
)

// CLI configurations
type Config struct {
	// Name of the CLI app
	Name string

	// Version of the CLI app
	Version string

	// Commands is a mapping of subcommand names to a Command implementation.
	// If there is a command with a blank string "", then it will be used as
	// the default command if no subcommand is specified.
	//
	// If the key has a space in it, this will create a nested subcommand.
	// For example, if the key is "foo bar", then to access it our CLI
	// must be accessed with "./cli foo bar".
	Commands map[string]Command

	// HelpFunc is the function called to generate the global help text.
	HelpFunc HelpFunc

	// HelpWriter is the Writer where the help text is outputted to. If
	// not specified, it will default to Stderr
	HelpWriter io.Writer
}

// Default CLI configurations
func DefaultConfig() *Config {
	config := &Config{
		Name:       "app",
		Commands:   map[string]Command{},
		HelpFunc:   DefaultHelpFunc("app"),
		HelpWriter: os.Stderr,
		Version:    "",
	}
	return config
}

func (c *Config) Merge(b *Config) *Config {

	if b == nil {
		return c
	}

	result := *c

	if b.Name != "" {
		result.Name = b.Name
	}
	if b.Version != "" {
		result.Version = b.Version
	}
	if b.Commands != nil {
		result.Commands = b.Commands
	}
	if b.HelpFunc != nil {
		result.HelpFunc = b.HelpFunc
	} else {
		if b.Name != "" {
			result.HelpFunc = DefaultHelpFunc(b.Name)
		}
	}
	if b.HelpWriter != nil {
		result.HelpWriter = b.HelpWriter
	}

	return &result
}
