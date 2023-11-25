package count

import (
	"testing"
)

func TestLines(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected int
	}{
		{"Empty string", []byte(""), 0},
		{"One line", []byte("Hello World"), 0},
		{"Two lines", []byte("Hello\nWorld"), 1},
		{"Three lines", []byte("Line 1\nLine 2\nLine 3"), 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Lines(test.data); got != test.expected {
				t.Errorf("Lines() = %d; want %d", got, test.expected)
			}
		})
	}
}

func TestCharacters(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected int
	}{
		{"Empty string", []byte(""), 0},
		{"English text", []byte("Hello"), 5},
		{"Multibyte characters", []byte("こんにちは"), 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Characters(test.data); got != test.expected {
				t.Errorf("Characters() = %d; want %d", got, test.expected)
			}
		})
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected int
	}{
		{"Empty string", []byte(""), 0},
		{"One word", []byte("Hello"), 1},
		{"Two words", []byte("Hello World"), 2},
		{"Sentence", []byte("The quick brown fox"), 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Words(test.data); got != test.expected {
				t.Errorf("Words() = %d; want %d", got, test.expected)
			}
		})
	}
}

func TestBytes(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected int
	}{
		{"Empty string", []byte(""), 0},
		{"English text", []byte("Hello"), 5},
		{"Multibyte characters", []byte("こんにちは"), 15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Bytes(test.data); got != test.expected {
				t.Errorf("Bytes() = %d; want %d", got, test.expected)
			}
		})
	}
}
