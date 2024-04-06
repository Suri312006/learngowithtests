package maps

import "errors"

type Dictionary map[string]string
type DictionaryErr string

var (
	ErrorNotFound       = errors.New("could not find the word you were looking for")
	ErrAlreadyExists    = errors.New("Word already exists in dictionary ")
	ErrWordDoesNotExist = errors.New("Word does not exist, cannot update")
)

func (d Dictionary) Search(word string) (string, error) {

	// can return two values, the second value is boolean whether or not the value exists in the dict
	defenition, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}

	return defenition, nil
}

// even though we are mutating the state inside map, we dont need a pointer to it

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorNotFound:
		d[key] = value

	case nil:
		return ErrAlreadyExists

	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrorNotFound:
		return ErrWordDoesNotExist

	case nil:
		d[key] = value

	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	// built in go function that works on map and key
	delete(d, key)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
