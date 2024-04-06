package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("searching for word in dictionary", func(t *testing.T) {

		got, _ := dict.Search("test")
		assertStrings(t, got, "this is just a test")

	})
	t.Run("searching for word not in the dictionary", func(t *testing.T) {
		_, err := dict.Search("unknown")

		assertError(t, err, ErrorNotFound)

	})
}

func TestAdd(t *testing.T) {
	t.Run("adding word", func(t *testing.T) {

		dictionary := Dictionary{}
		word := "test"
		defenition := "this is just a test"
		dictionary.Add(word, defenition)

		assertDefenition(t, dictionary, word, defenition)

	})

	t.Run("Adding existing word", func(t *testing.T) {

		word := "test"
		oldDefenition := "this is just a test"
		dict := Dictionary{word: oldDefenition}

		newDefenition := "this is so much more than just a dumb little test"

		err := dict.Add(word, newDefenition)
		assertError(t, err, ErrAlreadyExists)
		assertDefenition(t, dict, word, oldDefenition)

	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		defenition := "this is just a test"
		dict := Dictionary{word: defenition}
		newDefenition := "dwoahdiuahdiulwahaiudhaildhwi"

		dict.Update(word, newDefenition)
		assertDefenition(t, dict, word, newDefenition)

	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defenition := "this is just a test"
		dict := Dictionary{word: defenition}

		newWord := "gibberish"
		newDefenition := "dwoahdiuahdiulwahaiudhaildhwi"

		err := 	dict.Update(newWord, newDefenition)
		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestDelete(t *testing.T){
	t.Run("deleting existing word", func(t *testing.T) {

		word := "test"
		defenition := "this is just a test"
		dict := Dictionary{word: defenition}

		dict.Delete(word)

		_, err := dict.Search(word)
		if err != ErrorNotFound{
			t.Errorf("Expected %q to be deleted", word)
		}
		

	})
}

func assertDefenition(t testing.TB, dict Dictionary, word, defenition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, defenition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
