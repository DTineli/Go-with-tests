package main

import "testing"

func TestHello(t *testing.T) {

	want := "Hello, Tineli"

	got := Hello("Tineli")

	if got != want {
		t.Errorf("got: %q want %q", got, want)
	}
}
