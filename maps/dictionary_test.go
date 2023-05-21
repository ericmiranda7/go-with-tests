package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just a test"}
		got, err := dictionary.Search("test")
		want := "just a test"

		if err != nil {
			t.Fatal("wasn't expecting an error")
		}

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		_, got := dictionary.Search("noword")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("add a nonexisting word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just a test"}
		word := "max"
		meaning := "the highest value"

		dictionary.Add(word, meaning)

		got, err := dictionary.Search(word)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, got)
	})

	t.Run("adds an existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "just a test"}
		word := "test"
		meaning := "changed meaning"

		err := dictionary.Add(word, meaning)

		assertError(t, err, ErrAlreadyExists)
	})
}

func TestUpdate(t *testing.T) {
	d := Dictionary{"test": "just a test"}
	t.Run("existing word", func(t *testing.T) {
		err := d.Update("test", "trying something out")

		assertError(t, err, nil)
		assertDefinition(t, d, "test", "trying something out")
	})

	t.Run("new word", func(t *testing.T) {
		err := d.Update("newword", "new word")

		assertError(t, err, ErrWordAlreadyExists)
	})

}

func TestDelete(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		d := Dictionary{"test": "just a test"}
		d.Delete("test")
		_, err := d.Search("test")

		assertError(t, err, ErrNotFound)
	})

	//t.Run("non existing word", func(t *testing.T) {
	//	d := Dictionary{"test": "just a test"}
	//	d.Delete("tests")
	//})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Fatal("was expecting an error")
	}
}

func assertDefinition(t *testing.T, dict Dictionary, word, def string) {
	t.Helper()

	got := dict[word]

	assertStrings(t, got, def)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
