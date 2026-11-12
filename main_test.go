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
