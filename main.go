package main

import (
	"encoding/json"
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
  --json string      JSON object with optional "message" and "colour" keys
  --help             show this help

Examples:
  fs-scratch-0 --message "your message"
  fs-scratch-0 --colour "#FF0000"
  fs-scratch-0 --json '{"message": "My Message", "colour": "#00FF00"}'
`

func run(args []string, w io.Writer) error {
	fs := flag.NewFlagSet("fs-scratch-0", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	message := fs.String("message", "Hello World!", "message to write to stdout")
	colour := fs.String("colour", "", "colour to print the message in (red, blue, green or a #RRGGBB HTML colour hash)")
	jsonVal := fs.String("json", "", "JSON object with optional \"message\" and \"colour\" keys")
	if err := fs.Parse(args); err != nil {
		if err != flag.ErrHelp {
			return err
		}
		_, err = fmt.Fprint(w, helpText)
		return err
	}
	if *jsonVal != "" {
		setFlags := map[string]bool{}
		fs.Visit(func(f *flag.Flag) { setFlags[f.Name] = true })
		if setFlags["message"] || setFlags["colour"] {
			return errors.New("--json cannot be combined with --message or --colour.")
		}
		var fields map[string]json.RawMessage
		if err := json.Unmarshal([]byte(*jsonVal), &fields); err != nil {
			return fmt.Errorf("invalid --json value: %w", err)
		}
		if fields == nil {
			return errors.New("--json value must be a JSON object.")
		}
		for key := range fields {
			if key != "message" && key != "colour" {
				return fmt.Errorf("unknown key %q in --json value", key)
			}
		}
		if raw, ok := fields["message"]; ok {
			if err := json.Unmarshal(raw, message); err != nil {
				return fmt.Errorf("invalid --json value: %w", err)
			}
		}
		if raw, ok := fields["colour"]; ok {
			if err := json.Unmarshal(raw, colour); err != nil {
				return fmt.Errorf("invalid --json value: %w", err)
			}
		}
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
