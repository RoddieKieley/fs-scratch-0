package main

import (
	"bytes"
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
