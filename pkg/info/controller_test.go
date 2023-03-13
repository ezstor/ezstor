package info

import (
	"testing"
)

func TestControllers(t *testing.T) {
	c := Controllers()
	if len(c) == 0 {
		t.Fatalf("Controllers got nothing.")
	}
	t.Logf("Test Controllers success.")
}
