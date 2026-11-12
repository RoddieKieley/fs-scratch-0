package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, w io.Writer) error {
	fs := flag.NewFlagSet("fs-scratch-0", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	message := fs.String("message", "Hello World!", "message to write to stdout")
	if err := fs.Parse(args); err != nil {
		return err
	}
	_, err := fmt.Fprintln(w, *message)
	return err
}
