package maps

type Dictionary map[string]string

type DictionaryErr string

func (de DictionaryErr) Error() string {
	return string(de)
}

var (
	ErrNotFound          = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExists     = DictionaryErr("world already exists in dictionary")
	ErrWordAlreadyExists = DictionaryErr("there already is a word in the dictionary")
)

func (d Dictionary) Search(k string) (string, error) {
	v, ok := d[k]

	if !ok {
		return "", ErrNotFound
	}

	return v, nil
}

func (d Dictionary) Add(k, m string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		d[k] = m
		return nil
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}
}

func (d Dictionary) Update(k, def string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		return ErrWordAlreadyExists
	case nil:
		d[k] = def
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}
