package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunDefaultMessage(t *testing.T) {
	var buf bytes.Buffer
	if err := run(nil, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "Hello World!\n" {
		t.Fatalf("expected %q, got %q", "Hello World!\n", got)
	}
}

func TestRunCustomMessage(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--message", "My custom Hello World message!"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "My custom Hello World message!\n" {
		t.Fatalf("expected %q, got %q", "My custom Hello World message!\n", got)
	}
}

func TestRunColourRed(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--colour", "red"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[31mHello World!\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[31mHello World!\033[0m\n", got)
	}
}

func TestRunColourBlue(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--colour", "blue"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[34mHello World!\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[34mHello World!\033[0m\n", got)
	}
}

func TestRunColourGreen(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--colour", "green"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[32mHello World!\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[32mHello World!\033[0m\n", got)
	}
}

func TestRunColourWithCustomMessage(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--message", "My custom Hello World message!", "--colour", "blue"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[34mMy custom Hello World message!\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[34mMy custom Hello World message!\033[0m\n", got)
	}
}

func TestRunUnsupportedColour(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--colour", "yellow"}
	err := run(args, &buf)
	if err == nil {
		t.Fatalf("run() returned nil error, want error for unsupported colour")
	}
	if got := err.Error(); got != "Specified colour not supported." {
		t.Fatalf("expected error %q, got %q", "Specified colour not supported.", got)
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no output, got %q", buf.String())
	}
}

func TestRunColourWrongCase(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--colour", "Red"}
	err := run(args, &buf)
	if err == nil {
		t.Fatalf("run() returned nil error, want error for unsupported colour")
	}
	if got := err.Error(); got != "Specified colour not supported." {
		t.Fatalf("expected error %q, got %q", "Specified colour not supported.", got)
	}
}

func TestRunColourHex(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--colour", "#FF0000"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[38;2;255;0;0mHello World!\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[38;2;255;0;0mHello World!\033[0m\n", got)
	}
}

func TestRunColourHexLowercase(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--colour", "#ff0000"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[38;2;255;0;0mHello World!\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[38;2;255;0;0mHello World!\033[0m\n", got)
	}
}

func TestRunColourHexWithCustomMessage(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--message", "hi", "--colour", "#00FF00"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[38;2;0;255;0mhi\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[38;2;0;255;0mhi\033[0m\n", got)
	}
}

func TestRunJSONMessageAndColour(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `{"message":"My Message","colour":"#00FF00"}`}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "\033[38;2;0;255;0mMy Message\033[0m\n" {
		t.Fatalf("expected %q, got %q", "\033[38;2;0;255;0mMy Message\033[0m\n", got)
	}
}

func TestRunJSONMessageOnly(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `{"message":"hi"}`}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "hi\n" {
		t.Fatalf("expected %q, got %q", "hi\n", got)
	}
}

func TestRunJSONEmptyObject(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `{}`}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != "Hello World!\n" {
		t.Fatalf("expected %q, got %q", "Hello World!\n", got)
	}
}

func TestRunJSONInvalid(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `not-json`}
	err := run(args, &buf)
	if err == nil {
		t.Fatalf("run() returned nil error, want error for invalid JSON")
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no output, got %q", buf.String())
	}
}

func TestRunJSONNotObject(t *testing.T) {
	for _, value := range []string{`null`, `"hello"`, `123`} {
		var buf bytes.Buffer
		args := []string{"--json", value}
		err := run(args, &buf)
		if err == nil {
			t.Fatalf("run() returned nil error, want error for non-object JSON %q", value)
		}
		if buf.Len() != 0 {
			t.Fatalf("expected no output, got %q", buf.String())
		}
	}
}

func TestRunJSONUnknownKey(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `{"bogus":"x"}`}
	err := run(args, &buf)
	if err == nil {
		t.Fatalf("run() returned nil error, want error for unknown key")
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no output, got %q", buf.String())
	}
}

func TestRunJSONInvalidMessage(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `{"message":5}`}
	err := run(args, &buf)
	if err == nil {
		t.Fatalf("run() returned nil error, want error for non-string message")
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no output, got %q", buf.String())
	}
}

func TestRunJSONWithMessageFlag(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `{"message":"x"}`, "--message", "y"}
	err := run(args, &buf)
	if err == nil {
		t.Fatalf("run() returned nil error, want error for --json combined with --message")
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no output, got %q", buf.String())
	}
}

func TestRunJSONWithColourFlag(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--json", `{"message":"x"}`, "--colour", "red"}
	err := run(args, &buf)
	if err == nil {
		t.Fatalf("run() returned nil error, want error for --json combined with --colour")
	}
	if buf.Len() != 0 {
		t.Fatalf("expected no output, got %q", buf.String())
	}
}

func TestRunHelp(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--help"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	out := buf.String()
	for _, want := range []string{
		"--message",
		"message to write to stdout",
		"--colour",
		"colour to print the message in",
		"--json",
		`--message "your message"`,
		`--colour "#FF0000"`,
		`--json '{"message": "My Message", "colour": "#00FF00"}'`,
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("help output does not contain %q:\n%s", want, out)
		}
	}
}

func TestRunHelpTakesPrecedence(t *testing.T) {
	var buf bytes.Buffer
	args := []string{"--help", "--message", "x"}
	if err := run(args, &buf); err != nil {
		t.Fatalf("run() returned error: %v", err)
	}
	if got := buf.String(); got != helpText {
		t.Fatalf("expected help text only, got %q", got)
	}
}

func TestRunColourHexInvalid(t *testing.T) {
	for _, colour := range []string{"#FFF", "#GGGGGG", "FF0000"} {
		var buf bytes.Buffer
		args := []string{"--colour", colour}
		err := run(args, &buf)
		if err == nil {
			t.Fatalf("run() returned nil error, want error for unsupported colour %q", colour)
		}
		if got := err.Error(); got != "Specified colour not supported." {
			t.Fatalf("expected error %q, got %q", "Specified colour not supported.", got)
		}
		if buf.Len() != 0 {
			t.Fatalf("expected no output, got %q", buf.String())
		}
	}
}
