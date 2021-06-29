//go:generate yarn --cwd ./ui
//go:generate yarn --cwd ./ui build
//go:generate touch ./ui/build/.keep

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"

	"github.com/edufschmidt/go-cli"
	command "github.com/edufschmidt/go-cli/examples/commands"
)

func panicHandler() {

	if panicPayload := recover(); panicPayload != nil {

		stack := string(debug.Stack())
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "================================================================================")
		fmt.Fprintln(os.Stderr, "This app has encountered a fatal error. This is a bug!")
		fmt.Fprintln(os.Stderr, "We would appreciate a report: https://github.com/username/repo/issues/")
		fmt.Fprintln(os.Stderr, "Please provide all of the below text in your report.")
		fmt.Fprintln(os.Stderr, "================================================================================")

		fmt.Fprintf(os.Stderr, "App Version:         %s\n", "1.0.0")
		fmt.Fprintf(os.Stderr, "Go Version:          %s\n", runtime.Version())
		fmt.Fprintf(os.Stderr, "Go Compiler:         %s\n", runtime.Compiler)
		fmt.Fprintf(os.Stderr, "Architecture:        %s\n", runtime.GOARCH)
		fmt.Fprintf(os.Stderr, "Operating System:    %s\n", runtime.GOOS)
		fmt.Fprintf(os.Stderr, "Panic:               %s\n\n", panicPayload)
		fmt.Fprintln(os.Stderr, stack)
	}
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {

	cli := setupCLI()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, os.Interrupt)

	go func() {
		<-signalCh
		fmt.Fprintf(os.Stderr, "Received signal. Interrupting...\n")
		cancel()
	}()

	defer panicHandler()

	code, err := cli.Run(ctx, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return code
}

func setupCLI() *cli.CLI {

	ui := &cli.SimpleUI{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	cli := cli.New(&cli.Config{
		Name: "app",
		Commands: map[string]cli.Command{
			"foo": &command.FooCommand{UI: ui},
		},
		Version: "1.0.0",
	})

	return cli
}
