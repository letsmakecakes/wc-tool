package counter

import (
	"strings"
	"testing"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantLines int
		wantWords int
		wantBytes int
		wantChars int
	}{
		{
			name:      "empty string",
			input:     "",
			wantLines: 0,
			wantWords: 0,
			wantBytes: 0,
			wantChars: 0,
		},
		{
			name:      "single word",
			input:     "hello",
			wantLines: 0,
			wantWords: 1,
			wantBytes: 5,
			wantChars: 5,
		},
		{
			name:      "multiple words with newline",
			input:     "hello world\n",
			wantLines: 1,
			wantWords: 2,
			wantBytes: 12,
			wantChars: 12,
		},
		{
			name:      "multiple lines",
			input:     "hello\nworld\n",
			wantLines: 2,
			wantWords: 2,
			wantBytes: 12,
			wantChars: 12,
		},
		{
			name:      "unicode characters",
			input:     "hello 世界\n",
			wantLines: 1,
			wantWords: 2,
			wantBytes: 13,
			wantChars: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			err := c.Process(strings.NewReader(tt.input))
			if err != nil {
				t.Errorf("Process() error = %v", err)
				return
			}

			if got := c.Lines(); got != tt.wantLines {
				t.Errorf("Lines() = %v, want %v", got, tt.wantLines)
			}
			if got := c.Words(); got != tt.wantWords {
				t.Errorf("Words() = %v, want %v", got, tt.wantWords)
			}
			if got := c.Bytes(); got != tt.wantBytes {
				t.Errorf("Bytes() = %v, want %v", got, tt.wantBytes)
			}
			if got := c.Characters(); got != tt.wantChars {
				t.Errorf("Characters() = %v, want %v", got, tt.wantChars)
			}
		})
	}
}
