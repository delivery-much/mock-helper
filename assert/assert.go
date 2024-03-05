package assert

import "testing"

func True(t *testing.T, v bool) {
	if !v {
		t.Error("Expected value to be true, but it was false")
	}
}

func False(t *testing.T, v bool) {
	if v {
		t.Error("Expected value to be false, but it was true")
	}
}
