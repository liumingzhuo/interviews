package robot

import (
	"testing"
)

func TestRobot(t *testing.T) {
	x0, y0, _ := move("R2(LF)", 0, 0, 1)
	x, y := -1, 1
	if x != x0 || y != y0 {
		t.Errorf("error, x got %d y got %d", x0, y0)
	}
}
