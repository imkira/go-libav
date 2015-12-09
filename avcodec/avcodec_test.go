package avcodec

import (
	"testing"

	"github.com/imkira/go-libav/avutil"
)

func TestNewContextFromC(t *testing.T) {
	ctx := NewContextFromC(nil)
	if ctx == nil {
		t.Fatalf("Expecting context")
	}
}

func TestNewCodecDescriptorFromC(t *testing.T) {
	ctx := NewCodecDescriptorFromC(nil)
	if ctx == nil {
		t.Fatalf("Expecting context")
	}
}

func TestCodecDescriptor_Params(t *testing.T) {
	desc := CodecDescriptorByName("gif")
	if desc.ID() != 98 {
		t.Fatal("not match ID")
	}
	if desc.CodecType() != avutil.MediaTypeVideo {
		t.Fatal("not match CodecType")
	}
	if desc.Name() != "gif" {
		t.Fatal("not match Name")
	}
	if desc.LongName() != "GIF (Graphics Interchange Format)" {
		t.Fatal("not match LongName")
	}
	if desc.Props() != CodecPropLossless {
		t.Fatal("not match Props")
	}
	if desc.MimeTypes()[0] != "image/gif" {
		t.Fatal("not match MimeTypes")
	}

	desc = CodecDescriptorByName("sonic")
	if desc.Props() > 0 {
		t.Fatal("not match Props")
	}
	if len(desc.MimeTypes()) > 0 {
		t.Fatal("not match MimeTypes")
	}
}

func TestCodecDescriptorByID(t *testing.T) {
	found := CodecDescriptorByID(CodecID(28))
	if found == nil {
		t.Fatal("not found")
	}
	notFound := CodecDescriptorByID(CodecID(0))
	if notFound != nil {
		t.Fatal("found")
	}
}

func TestCodecDescriptorByName(t *testing.T) {
	found := CodecDescriptorByName("h264")
	if found == nil {
		t.Fatal("not found")
	}
	notFound := CodecDescriptorByName("notfound")
	if notFound != nil {
		t.Fatal("found")
	}
}

func TestCodecDescriptors(t *testing.T) {
	descriptors := CodecDescriptors()
	if len(descriptors) == 0 {
		t.Fatal("not found")
	}
}
