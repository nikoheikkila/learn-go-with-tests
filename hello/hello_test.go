package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
	}

	t.Run("should say hello to people", func(t *testing.T) {
		got := hello("Chris")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should say 'Hello World' when name left blank", func(t *testing.T) {
		got := hello("")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}
