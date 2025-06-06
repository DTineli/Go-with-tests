package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		want := "Hello, Tineli"
		got := Hello("Tineli", "")
		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return Hello, world as default", func(t *testing.T) {
		want := "Hello, world"
		got := Hello("", "")
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", "Spanish")
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		want := "Bonjour, Elodie"
		got := Hello("Elodie", "French")
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q want %q", got, want)
	}
}
