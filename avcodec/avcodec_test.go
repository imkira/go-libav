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

func TestCodecProfileName(t *testing.T) {
	codec := FindDecoderByName("h264")
	if codec == nil {
		t.Fatal("codec not found")
	}
	name := codec.ProfileName(100)
	if name != "High" {
		t.Errorf("profile name expected:High, got:%s", name)
	}
	name = codec.ProfileName(1)
	if name != "" {
		t.Errorf("unexpected profile name, got:%s", name)
	}
}

func TestCodecProfiles(t *testing.T) {
	type data struct {
		id   int
		name string
	}
	datas := []*data{
		&data{id: 66, name: "Baseline"},
		&data{id: 578, name: "Constrained Baseline"},
		&data{id: 77, name: "Main"},
		&data{id: 88, name: "Extended"},
		&data{id: 100, name: "High"},
		&data{id: 110, name: "High 10"},
		&data{id: 2158, name: "High 10 Intra"},
		&data{id: 122, name: "High 4:2:2"},
		&data{id: 2170, name: "High 4:2:2 Intra"},
		&data{id: 144, name: "High 4:4:4"},
		&data{id: 244, name: "High 4:4:4 Predictive"},
		&data{id: 2292, name: "High 4:4:4 Intra"},
		&data{id: 44, name: "CAVLC 4:4:4"},
	}
	codec := FindDecoderByName("h264")
	if codec == nil {
		t.Fatal("codec not found")
	}
	profiles := codec.Profiles()
	if len(datas) != len(profiles) {
		t.Fatalf("profiles count expected:%d, got:%d", len(datas), len(profiles))
	}
	for i, profile := range profiles {
		if datas[i].id != profile.ID() {
			t.Errorf("profile id expected:%d, got:%d", datas[i].id, profile.ID())
		}
		if datas[i].name != profile.Name() {
			t.Errorf("profile name expected:%s, got:%s", datas[i].name, profile.Name())
		}
	}
}
