package avfilter

import (
	"testing"

	"github.com/imkira/go-libav/avutil"
)

func TestNewFilterFromC(t *testing.T) {
	ctx := NewFilterFromC(nil)
	if ctx == nil {
		t.Fatalf("Expecting filter")
	}
}
func TestGraphRequestOldest(t *testing.T) {
	graph, err := NewGraph()
	if err != nil {
		t.Fatal(err)
	}
	defer graph.Free()

	result := graph.RequestOldest()
	if result.(*avutil.Error).Code() != avutil.ErrorCodeEOF {
		t.Fatalf("[TestGraphRequestOldest] result = %d, NG, expected = %d",
			result.(*avutil.Error).Code(), avutil.ErrorCodeEOF)
	}
}
