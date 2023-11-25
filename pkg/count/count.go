package count

import (
	"strings"
)

// Lines counts the number of lines in the given byte slice.
func Lines(data []byte) int {
	lineCount := 0
	for _, b := range data {
		if b == '\n' {
			lineCount++
		}
	}
	return lineCount
}

// Characters counts the number of characters in the given byte slice.
// It treats each byte as a character, which might not be accurate for multi-byte characters.
func Characters(data []byte) int {
	return len(string(data))
}

// Words counts the number of words in the given byte slice.
// It uses strings.Fields to split the string into words based on whitespace.
func Words(data []byte) int {
	str := string(data)
	fields := strings.Fields(str)
	return len(fields)
}

// Bytes returns the byte count of the given byte slice.
func Bytes(data []byte) int {
	return len(data)
}
