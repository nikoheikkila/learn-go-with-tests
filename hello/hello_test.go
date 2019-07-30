package main

import "testing"

func TestHello(t *testing.T) {

	assertSame := func(t *testing.T, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q, expected %q", got, want)
		}
	}

	t.Run("should execute main function", func(t *testing.T) {
		main()
	})

	t.Run("should say hello to people", func(t *testing.T) {
		got := hello("Chris", "English")
		want := "Hello, Chris!"

		assertSame(t, got, want)
	})

	t.Run("should say 'Hello World' when name left blank", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, world!"

		assertSame(t, got, want)
	})

	t.Run("should greet in Spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertSame(t, got, want)
	})

	t.Run("should greet in French", func(t *testing.T) {
		got := hello("Michel", "French")
		want := "Bonjour, Michel!"

		assertSame(t, got, want)
	})
}
