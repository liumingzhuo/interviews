package reversestr

import "testing"

func TestReverseStr(t *testing.T) {
	got, ok := reverseString("abcdefg")
	want := "gfedcba"
	if !ok || got != want {
		t.Fatalf("excption: want:%s,but got:%s", want, got)
	}
}
