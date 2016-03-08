package avcodec

import (
	"testing"

	"github.com/imkira/go-libav/avutil"
)

func TestNewPacket(t *testing.T) {
	pkt, err := NewPacket()
	if err != nil {
		t.Fatal(err)
	}
	defer pkt.Free()
	if pkt == nil {
		t.Fatalf("Expecting packet")
	}
}

func TestPacketFree(t *testing.T) {
	pkt, _ := NewPacket()
	if pkt.CAVPacket == nil {
		t.Fatalf("Expecting packet")
	}
	for i := 0; i < 3; i++ {
		pkt.Free()
		if pkt.CAVPacket != nil {
			t.Fatalf("Not expecting packet")
		}
	}
}

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

func TestContextStatInOutOK(t *testing.T) {
	codec := FindEncoderByName("mpeg4")
	if codec == nil {
		t.Error("error")
	}
	ctx, err := NewContextWithCodec(codec)
	if err != nil {
		t.Error("error")
	}
	defer ctx.Free()

	expected := "test"
	if err := ctx.SetStatsIn(avutil.String(expected)); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected not error", err)
	}
	result, ok := ctx.StatsIn()
	if !ok {
		t.Fatalf("[TestContextStatInOutOK] ok=%t, NG, expected ok=true", ok)
	}
	if result != expected {
		t.Fatalf("[TestContextStatInOutOK] result=%s NG, expected=%s", result, expected)
	}

	if err := ctx.SetStatsIn(nil); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected not error", err)
	}
	result, ok = ctx.StatsIn()
	if ok {
		t.Fatalf("[TestContextStatInOutOK] ok=%t, NG, expected ok=false", ok)
	}
	if result != "" {
		t.Fatalf("[TestContextStatInOutOK] result=%s NG, expected=\"\"(empty)", result)
	}

	if err := ctx.SetStatsOut(avutil.String(expected)); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected not error", err)
	}
	result, ok = ctx.StatsOut()
	if !ok {
		t.Fatalf("[TestContextStatInOutOK] ok=%t, NG, expected ok=true", ok)
	}
	if result != expected {
		t.Fatalf("[TestContextStatInOutOK] result=%s NG, expected=%s", result, expected)
	}

	if err := ctx.SetStatsOut(nil); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected not error", err)
	}
	result, ok = ctx.StatsOut()
	if ok {
		t.Fatalf("[TestContextStatInOutOK] ok=%t, NG, expected ok=false", ok)
	}
	if result != "" {
		t.Fatalf("[TestContextStatInOutOK] result=%s NG, expected=\"\"(empty)", result)
	}
}
