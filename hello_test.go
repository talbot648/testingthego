package main

import "testing"

// rules
//Write a test
// make the compiler passs
// run the test, see if it fails and check the error message is meaningful
//write enough code to make the test pass
//refactor
//test

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})
	t.Run("in Spanish,", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French,", func(t *testing.T) {
		got := Hello("Dominique", "French")
		want := "Bonjour, Dominique"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Polish,", func(t *testing.T) {
		got := Hello("Bogdi", "Polish")
		want := "Witam, Bogdi"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)

	}
}
