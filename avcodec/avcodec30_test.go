// +build ffmpeg30

package avcodec

import (
	"reflect"
	"testing"

	"github.com/baohavan/go-libav/avutil"
)

func TestNewBitStreamFilterContextFromName(t *testing.T) {
	ctx, err := NewBitStreamFilterContextFromName("invalid")
	if err != ErrBitStreamFilterNotFound {
		t.Fatalf("[NewBitStreamFilterContextFromName] err=%v, NG expected=%v", err, ErrBitStreamFilterNotFound)
	}
	if ctx != nil {
		t.Fatalf("[NewBitStreamFilterContextFromName] ctx=%v, NG expected is nil", ctx)
	}
	ctx, err = NewBitStreamFilterContextFromName("h264_mp4toannexb")
	if err != nil {
		t.Fatalf("[NewBitStreamFilterContextFromName] err=%v, NG expected not error", err)
	}
	if ctx == nil {
		t.Fatalf("[NewBitStreamFilterContextFromName] ctx is nil, NG expected is not nil")
	}
	ctx.Close()
}

func TestBitStreamFilterContext_Next(t *testing.T) {
	ctx, err := NewBitStreamFilterContextFromName("h264_mp4toannexb")
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()
	result := ctx.Next()
	if result != nil {
		t.Fatalf("[TestBitStreamFilterContext_Next] result=%v, NG expected nil", result)
	}

	next, err := NewBitStreamFilterContextFromName("mjpeg2jpeg")
	if err != nil {
		t.Fatal(err)
	}
	defer next.Close()
	ctx.SetNext(next)
	result = ctx.Next()
	if !reflect.DeepEqual(next, result) {
		t.Fatalf("[TestBitStreamFilterContext_Next] next=%p, getNext=%p, NG expected same", next, result)
	}
}

func TestBitStreamFilterContext_Args(t *testing.T) {
	ctx, err := NewBitStreamFilterContextFromName("h264_mp4toannexb")
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Close()

	_, ok := ctx.ArgsOK()
	if ok {
		t.Fatalf("[TestBitStreamFilterContext_Args] ok=%t, NG expected=%t", ok, false)
	}
	result := ctx.Args()
	if result != "" {
		t.Fatalf("[TestBitStreamFilterContext_Args] result=%s, NG expected blank", result)
	}

	input := avutil.String("argstest")
	if err := ctx.SetArgs(input); err != nil {
		t.Fatalf("[TestBitStreamFilterContext_Args] err=%v, NG expected not error", err)
	}
	_, ok = ctx.ArgsOK()
	if !ok {
		t.Fatalf("[TestBitStreamFilterContext_Args] ok=%t, NG expected=%t", ok, true)
	}
	result = ctx.Args()
	if result != *input {
		t.Fatalf("[TestBitStreamFilterContext_Args] result=%s, NG expected=%s", result, *input)
	}

	if err := ctx.SetArgs(nil); err != nil {
		t.Fatalf("[TestBitStreamFilterContext_Args] err=%v, NG expected not error", err)
	}
	_, ok = ctx.ArgsOK()
	if ok {
		t.Fatalf("[TestBitStreamFilterContext_Args] ok=%t, NG expected=%t", ok, false)
	}
	result = ctx.Args()
	if result != "" {
		t.Fatalf("[TestBitStreamFilterContext_Args] result=%s, NG expected blank", result)
	}
}

func TestBitStreamFilterContext_CloseAll1M(t *testing.T) {
	before := testMemoryUsed(t)
	for i := 0; i < 1000000; i++ {
		ctx, err := NewBitStreamFilterContextFromName("h264_mp4toannexb")
		if err != nil {
			t.Fatal(err)
		}
		ctx.Close()
	}
	testMemoryLeak(t, before, 50*1024*1024)
}
