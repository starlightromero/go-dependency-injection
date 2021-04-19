package main

import (
	"bytes"
	"testing"
)

func TestGeet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Star")

	got := buffer.String()
	want := "Hello, Star"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
