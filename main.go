package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const helpText = `Usage:
  fs-scratch-0 [flags]

Flags:
  --message string   message to write to stdout (default "Hello World!")
  --colour string    colour to print the message in (red, blue, green or a #RRGGBB HTML colour hash)
  --help             show this help

Examples:
  fs-scratch-0 --message "your message"
  fs-scratch-0 --colour "#FF0000"
`

func run(args []string, w io.Writer) error {
	fs := flag.NewFlagSet("fs-scratch-0", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	message := fs.String("message", "Hello World!", "message to write to stdout")
	colour := fs.String("colour", "", "colour to print the message in (red, blue, green or a #RRGGBB HTML colour hash)")
	if err := fs.Parse(args); err != nil {
		if err != flag.ErrHelp {
			return err
		}
		_, err = fmt.Fprint(w, helpText)
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
		r, g, b, ok := parseHexColour(*colour)
		if !ok {
			return errors.New("Specified colour not supported.")
		}
		prefix = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		suffix = "\033[0m"
	}
	_, err := fmt.Fprintln(w, prefix+*message+suffix)
	return err
}

// parseHexColour parses a #RRGGBB HTML colour hash (case-insensitive)
// into its red, green and blue components.
func parseHexColour(colour string) (r, g, b int, ok bool) {
	if len(colour) != 7 || colour[0] != '#' {
		return 0, 0, 0, false
	}
	var components [3]int
	for i := 0; i < 3; i++ {
		v, err := strconv.ParseUint(colour[1+i*2:3+i*2], 16, 8)
		if err != nil {
			return 0, 0, 0, false
		}
		components[i] = int(v)
	}
	return components[0], components[1], components[2], true
}
