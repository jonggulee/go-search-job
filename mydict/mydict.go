package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("Cant delete non-existing word")
)

func (d Dictionary) Search(word string) (string, error) {
	word, err := d[word]
	if err {
		return word, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantUpdate
	} else {
		d[word] = def
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantDelete
	} else {
		delete(d, word)
	}
	return nil
}
