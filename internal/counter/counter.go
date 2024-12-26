package counter

import (
	"bufio"
	"io"
	"unicode"
)

// Counter tracks lines, words, bytes, and characters in a text stream.
type Counter struct {
	lines      int
	words      int
	bytes      int
	characters int
}

// New initializes and returns a new Counter instance.
func New() *Counter {
	return &Counter{}
}

// Process reads from an io.Reader and updates the counter's statistics.
func (c *Counter) Process(r io.Reader) error {
	reader := bufio.NewReader(r)
	inWord := false

	for {
		char, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		c.bytes += size
		if char == unicode.ReplacementChar && size == 1 {
			// Invalid UTF-8 byte sequence
			continue
		}
		c.characters++

		if char == '\n' {
			c.lines++
		}

		// Word counting
		if unicode.IsSpace(char) {
			inWord = false
		} else {
			if !inWord {
				c.words++
				inWord = true
			}
		}
	}

	return nil
}

// Lines returns the number of lines counted.
func (c *Counter) Lines() int {
	return c.lines
}

// Words returns the number of words counted.
func (c *Counter) Words() int {
	return c.words
}

// Bytes returns the number of bytes counted.
func (c *Counter) Bytes() int {
	return c.bytes
}

// Characters returns the number of characters counted.
func (c *Counter) Characters() int {
	return c.characters
}
