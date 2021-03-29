package diffregroupstr

import "testing"

func TestGroupStr(t *testing.T) {
	got := isGrouStr("abc", "bcae")
	want := true
	if got != want {
		t.Fatalf("exception  want:%t,but got:%t", want, got)
	}
}
