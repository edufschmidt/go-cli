package cli

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// HelpFunc is the type of the function that is responsible for generating
// the help output when the CLI must show the general help text
type HelpFunc func(map[string]Command) string

// DefaultHelpFunc is the default function for generating help output
func DefaultHelpFunc(app string) HelpFunc {
	return func(commands map[string]Command) string {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf(
			"Usage: %s [--version] [--help] <command> [<args>]\n\n",
			app))
		buf.WriteString("Available commands:\n")

		// Get the list of keys so we can sort them, and also get the maximum
		// key length so they can be aligned properly.
		keys := make([]string, 0, len(commands))
		maxKeyLen := 0
		for key := range commands {
			if len(key) > maxKeyLen {
				maxKeyLen = len(key)
			}
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			command, ok := commands[key]
			if !ok {
				panic("command not found: " + key)
			}

			key = fmt.Sprintf("%s%s", key, strings.Repeat(" ", maxKeyLen-len(key)))
			buf.WriteString(fmt.Sprintf("    %s    %s\n", key, command.Synopsis()))
		}

		return buf.String()
	}
}
