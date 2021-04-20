package replacestr

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	got, err := replaceStr("ab ca")
	if err != nil {
		t.Error(err)
	}
	want := "ab%20ca"
	if got != want {
		fmt.Printf("got: %s but want :%s", got, want)
	}
}
