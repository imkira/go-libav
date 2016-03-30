package avfilter

import (
	"strings"
	"testing"

	"github.com/imkira/go-libav/avutil"
)

func TestVersion(t *testing.T) {
	major, minor, micro := Version()
	if major < 6 || minor < 0 || micro < 0 {
		t.Fatalf("Invalid version")
	}
}

func TestNewFilterFromC(t *testing.T) {
	ctx := NewFilterFromC(nil)
	if ctx == nil {
		t.Fatalf("Expecting filter")
	}
}

func TestGraphRequestOldest(t *testing.T) {
	graph := testNewGraph(t)
	defer graph.Free()

	result := graph.RequestOldest()
	if result.(*avutil.Error).Code() != avutil.ErrorCodeEOF {
		t.Fatalf("[TestGraphRequestOldest] result = %d, NG, expected = %d",
			result.(*avutil.Error).Code(), avutil.ErrorCodeEOF)
	}
}

func TestFilterByNameOK(t *testing.T) {
	filter := testFilterByName(t, "displace")
	expectedDiscription := "Displace pixels."
	if filter.Description() != expectedDiscription {
		t.Errorf("[TestFilterByNameOK] description expected: %s, got: %s", expectedDiscription, filter.Description())
	}
	if filter.Flags() != FlagSupportTimelineInternal {
		t.Errorf("[TestFilterByNameOK] flags expected: 0x%x, got: 0x%x", FlagSupportTimelineInternal, filter.Flags())
	}
}

func TestFilterByNameRequiredNameParam(t *testing.T) {
	filter := FindFilterByName("")
	if filter != nil {
		t.Fatalf("Expecting filter is nil")
	}
}

func TestFilterByNameInvalidNameParam(t *testing.T) {
	filter := FindFilterByName("aaa")
	if filter != nil {
		t.Fatalf("Expecting filter is nil")
	}
}

func TestAddFilterOK(t *testing.T) {
	graph := testNewGraph(t)
	defer graph.Free()
	testAddFilter(t, graph, "scale", "rescale")
	testAddFilter(t, graph, "fps", "fps1")
	if graph.NumberOfFilters() != 2 {
		t.Fatalf("[TestAddFilterOK] number of filters expected: 2, got: %d", graph.NumberOfFilters())
	}
}

func testFilterByName(t *testing.T, name string) *Filter {
	filter := FindFilterByName(name)
	if filter == nil {
		t.Fatalf("Expecting filter")
	}
	if !strings.EqualFold(name, filter.Name()) {
		t.Fatalf("[testFilterByName] name expected: %s, got: %s", name, filter.Name())
	}
	return filter
}

func testAddFilter(t *testing.T, graph *Graph, name, id string) *Context {
	filter := testFilterByName(t, name)
	ctx, err := graph.AddFilter(filter, id)
	if err != nil {
		t.Fatal(err)
	}
	if ctx == nil {
		t.Fatalf("Expecting filter context")
	}
	if !strings.EqualFold(id, ctx.Name()) {
		t.Fatalf("[testAddFilter] context name expected: %s, got: %s", id, ctx.Name())
	}
	if ctx.Filter() == nil {
		t.Fatalf("Expecting filter")
	}
	if ctx.NumberOfInputs() <= 0 {
		t.Fatalf("[testAddFilter] number of inputs expected: greater than 0, got: %d", ctx.NumberOfInputs())
	}
	if ctx.NumberOfOutputs() <= 0 {
		t.Fatalf("[testAddFilter] number of outputs expected: greater than 0, got: %d", ctx.NumberOfOutputs())
	}
	return ctx
}

func testNewGraph(t *testing.T) *Graph {
	graph, err := NewGraph()
	if err != nil {
		t.Fatal(err)
	}
	if graph == nil {
		t.Fatalf("Expecting filter graph")
	}
	return graph
}
