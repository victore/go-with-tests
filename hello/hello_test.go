package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Victor")
		want := "Hello, Victor"
		assertCorrectMessage(t, got, want)
	})

	t.Run("soy 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// tell Go this is a helper function. If a test fails, the line number reported will be in our function call
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
