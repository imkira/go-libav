package avfilter

import "testing"

func TestNewFilterFromC(t *testing.T) {
	ctx := NewFilterFromC(nil)
	if ctx == nil {
		t.Fatalf("Expecting filter")
	}
}
