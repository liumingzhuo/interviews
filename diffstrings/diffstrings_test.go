package diffstrings

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	input := "abcd123"
	got := DiffStrings(input)
	want := true
	if got != want {
		fmt.Printf("error, want: %t, but got %t", want, got)
	}
}
