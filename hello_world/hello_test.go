package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hi to a specific person", func(t *testing.T) {
		name := "Eric"
		got := Hello(name, "")
		want := "Hello, Eric"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hi to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("No language defaults to english", func(t *testing.T) {
		got := Hello("Eric", "")
		want := "Hello, Eric"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Spanish says hola", func(t *testing.T) {
		got := Hello("Eric", "spanish")
		want := "Hola, Eric"

		assertCorrectMessage(t, got, want)
	})

	t.Run("English says hello", func(t *testing.T) {
		got := Hello("Eric", "english")
		want := "Hello, Eric"

		assertCorrectMessage(t, got, want)
	})

	t.Run("French says bonjour", func(t *testing.T) {
		got := Hello("Eric", "french")
		want := "Bonjour, Eric"

		assertCorrectMessage(t, got, want)
	})

	t.Run("No name french", func(t *testing.T) {
		got := Hello("", "french")
		want := "Bonjour, world"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Hello(): got %q, want %q", got, want)
	}
}
