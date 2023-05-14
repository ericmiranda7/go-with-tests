package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hi to a specific person", func(t *testing.T) {
		name := "Eric"
		got := Hello(name)
		want := "Hello, Eric"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hi to the world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Hello(): got %q, want %q", got, want)
	}
}
