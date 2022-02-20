package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a just Test"}
	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a just Test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defi := "this is just a test"

		err := dictionary.Add(word, defi)

		assertError(t, err, nil)
		assertDefi(t, dictionary, word, defi)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defi := "this is just a test"
		dictionary := Dictionary{word: defi}

		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefi(t, dictionary, word, defi)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("exsisting word", func(t *testing.T) {
		word := "test"
		defi := "this is test"
		newDefi := "new Definition"
		dictionary := Dictionary{word: defi}

		err := dictionary.Update(word, newDefi)

		assertError(t, err, nil)
		assertDefi(t, dictionary, word, newDefi)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		defi := "this is test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, defi)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test Definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", err)
	}
}

func assertDefi(t testing.TB, dictionary Dictionary, word, defi string) {
	t.Helper()
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if defi != got {
		t.Errorf("got %q want %q", got, defi)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
