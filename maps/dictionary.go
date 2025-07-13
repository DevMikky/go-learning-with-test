package maps

import (
	"errors"
)

const (
	ErrNotFound           = DictionaryErr("the word is not in the dictionary")
	ErrAlreadyExists      = DictionaryErr("the word already exists")
	ErrWordDoesNotExist   = DictionaryErr("cannot perform operation on word because it does not exist")
	ErrWordDeleteNotExist = DictionaryErr("cannot perform operation on word because not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {

	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {

		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {

	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDeleteNotExist
	case err == nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
