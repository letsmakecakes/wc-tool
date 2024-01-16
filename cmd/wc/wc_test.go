package wc

import (
	"testing"
)

func TestNumberOfBytes(t *testing.T) {
	fileName := "test.txt"
	got, err := WC("-c", fileName)
	if err != nil {
		t.Error(err)
	}

	want := "342190 test.txt"
	if got != want {
		t.Errorf("Got: %s, want: %s", got, want)
	}
}

func TestNumberOfLines(t *testing.T) {
	fileName := "test.txt"
	got, err := WC("-l", fileName)
	if err != nil {
		t.Error(err)
	}

	want := "7145 test.txt"

	if got != want {
		t.Error("String don't match")
	}
}

func TestNumberOfWords(t *testing.T) {
	fileName := "test.txt"
	got, err := WC("-w", fileName)
	if err != nil {
		t.Error(err)
	}

	want := "58164 test.txt"
	if got != want {
		t.Error("String don't match")
	}
}

func TestNumberOfCharacters(t *testing.T) {
	fileName := "test.txt"
	got, err := WC("-m", fileName)
	if err != nil {
		t.Error(err)
	}

	want := "339292 test.txt"
	if got != want {
		t.Errorf("Strings don't match: %s", got)
	}
}

func TestAllOptions(t *testing.T) {
	fileName := "test.txt"
	got, err := WC("", fileName)
	if err != nil {
		t.Error(err)
	}

	want := "7145   58164  342190 test.txt"
	if got != want {
		t.Errorf("No match: %s", got)
	}
}
