package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search word in dictionary", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("search word not in dictionary", func(t *testing.T) {

		_, err := dictionary.Search("unknown")

		if err == nil {

			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	dictionary := Dictionary{}

	t.Run("add word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add word already exists", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrAlreadyExists)

		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		updatedDefinition := "this is updated definition"

		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, updatedDefinition)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, word, updatedDefinition)
	})

	t.Run("update word not exist", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {

	t.Run("delete word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})

	t.Run("delete word not exist", func(t *testing.T) {

		word := "test"

		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDeleteNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {

	t.Helper()

	if !errors.Is(got, want) {

		t.Errorf("got %q want %q", got, want)
	}
}
