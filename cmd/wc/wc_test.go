package wc

import (
	"testing"
)

func TestNumberOfBytes(t *testing.T) {
	fileName := "test.txt"
	got, err := wc("-c", fileName)
	if err != nil {
		t.Error(err)
	}

	want := "342190 test.txt"
	if got != want {
		t.Error("Strings don't match")
	}
}

func TestNumberOfLines(t *testing.T) {
	fileName := "test.txt"
	got, err := wc("-l", fileName)
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
	got, err := wc("-w", fileName)
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
	got, err := wc("-m", fileName)
	if err != nil {
		t.Error(err)
	}

	want := "339292 test.txt"
	if got != want {
		t.Error("Strings don't match")
	}
}
