package helloworld

import (
	"testing"
)

// tests need to be in a file with a name like xxx_test.go
// test function must start with word "Test"
// test fuction only takes in one argument `t *testing.T`
// in order to use `t *testing.T` you need to import "testing"

//you can do `t.Fail()` when you want to fail

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied",
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"

			assertCorrectMessage(t, got, want)

		})

	t.Run("in Spanish", func(t *testing.T){
		got:= Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T){
		got:= Hello("Ludwig", "French")
		want := "Bonjour, Ludwig"

		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got, want string) {
    // throws error where the test failed instead of here
	t.Helper()
	if got != want {
		//f is like f string in python
		t.Errorf("Got %q want %q", got, want)
	}
}
