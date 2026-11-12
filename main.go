package main

import (
	"errors"
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
	colour := fs.String("colour", "", "colour to print the message in (red, blue or green)")
	if err := fs.Parse(args); err != nil {
		return err
	}
	var prefix, suffix string
	switch *colour {
	case "":
	case "red":
		prefix, suffix = "\033[31m", "\033[0m"
	case "blue":
		prefix, suffix = "\033[34m", "\033[0m"
	case "green":
		prefix, suffix = "\033[32m", "\033[0m"
	default:
		return errors.New("Specified colour not supported.")
	}
	_, err := fmt.Fprintln(w, prefix+*message+suffix)
	return err
}
