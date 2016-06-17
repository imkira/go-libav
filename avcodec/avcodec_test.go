package avcodec

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/imkira/go-libav/avutil"
	"github.com/shirou/gopsutil/process"
)

func TestVersion(t *testing.T) {
	major, minor, micro := Version()
	if major < 57 || minor < 0 || micro < 0 {
		t.Fatalf("Invalid version")
	}
}

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

func TestPacketDuration(t *testing.T) {
	pkt, _ := NewPacket()
	defer pkt.Free()
	data := int64(100000)
	pkt.SetDuration(data)
	if pkt.Duration() != data {
		t.Fatalf("packet duration expected:%d, got:%d", data, pkt.Duration())
	}
}

func TestPacketNewFreeLeak10M(t *testing.T) {
	before := testMemoryUsed(t)
	for i := 0; i < 10000000; i++ {
		pkt, err := NewPacket()
		if err != nil {
			t.Fatal(err)
		}
		pkt.Free()
	}
	testMemoryLeak(t, before, 50*1024*1024)
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
	ctx := testNewContextWithCodec(t, "mpeg4")
	codec := FindEncoderByName("mpeg4")
	if codec == nil {
		t.Error("error")
	}
	ctx, err := NewContextWithCodec(codec)
	if err != nil {
		t.Error("error")
	}
	defer ctx.Free()

	expected := []byte("stats_in")
	if err := ctx.SetStatsIn(expected); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected is not error", err)
	}
	result := ctx.StatsIn()
	if !bytes.Equal(result, expected) {
		t.Fatalf("[TestContextStatInOutOK] result=%s NG, expected=%s", result, expected)
	}
	expected = []byte{}
	if err := ctx.SetStatsIn(expected); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected is not error", err)
	}
	result = ctx.StatsIn()
	if !bytes.Equal(result, expected) {
		t.Fatalf("[TestContextStatInOutOK] result=%v NG, expected=%v", result, expected)
	}
	if err := ctx.SetStatsIn(nil); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected is not error", err)
	}
	result = ctx.StatsIn()
	if result != nil {
		t.Fatalf("[TestContextStatInOutOK] result=%v NG, expected=nil", result)
	}

	expected = []byte("stats_out")
	if err := ctx.SetStatsOut(expected); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected is not error", err)
	}
	result = ctx.StatsOut()
	if !bytes.Equal(result, expected) {
		t.Fatalf("[TestContextStatInOutOK] result=%s NG, expected=%s", result, expected)
	}
	expected = []byte{}
	if err := ctx.SetStatsOut(expected); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected is not error", err)
	}
	result = ctx.StatsOut()
	if !bytes.Equal(result, expected) {
		t.Fatalf("[TestContextStatInOutOK] result=%v NG, expected=%v", result, expected)
	}
	if err := ctx.SetStatsOut(nil); err != nil {
		t.Fatalf("[TestContextStatInOutOK] err=%v NG, expected is not error", err)
	}
	result = ctx.StatsOut()
	if result != nil {
		t.Fatalf("[TestContextStatInOutOK] result=%v NG, expected=nil", result)
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

func TestContextBitRate(t *testing.T) {
	ctx := testNewContextWithCodec(t, "h264")
	defer ctx.Free()
	data := int64(180)
	ctx.SetBitRate(data)
	if ctx.BitRate() != data {
		t.Fatalf("context bitrate expected:%d, got:%d", data, ctx.BitRate())
	}
}

func TestContextRCMaxRate(t *testing.T) {
	ctx := testNewContextWithCodec(t, "h264")
	defer ctx.Free()
	data := int64(200)
	ctx.SetRCMaxRate(data)
	if ctx.RCMaxRate() != data {
		t.Fatalf("context rc maxrate expected:%d, got:%d", data, ctx.RCMaxRate())
	}
}

func TestContextRCMinRate(t *testing.T) {
	ctx := testNewContextWithCodec(t, "h264")
	defer ctx.Free()
	data := int64(50)
	ctx.SetRCMinRate(data)
	if ctx.RCMinRate() != data {
		t.Fatalf("context rc minrate expected:%d, got:%d", data, ctx.RCMinRate())
	}
}

func TestContextNewFreeLeak1M(t *testing.T) {
	before := testMemoryUsed(t)
	for i := 0; i < 1000000; i++ {
		ctx := testNewContextWithCodec(t, "mpeg4")
		ctx.Free()
	}
	testMemoryLeak(t, before, 50*1024*1024)
}

func testNewContextWithCodec(t *testing.T, name string) *Context {
	codec := FindDecoderByName(name)
	if codec == nil {
		t.Fatalf("Expecting codec")
	}
	ctx, err := NewContextWithCodec(codec)
	if err != nil {
		t.Fatal(err)
	}
	if ctx == nil {
		t.Fatalf("Expecting context")
	}
	return ctx
}

func testMemoryUsed(t *testing.T) uint64 {
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		t.Fatal(err)
	}
	info, err := p.MemoryInfo()
	if err != nil {
		t.Fatal(err)
	}
	return info.RSS
}

func testMemoryLeak(t *testing.T, before uint64, diff uint64) {
	after := testMemoryUsed(t)
	if after > before && after-before > diff {
		t.Fatalf("memory leak detected: %d bytes", after-before)
	}
}

func TestFindBestPixelFormat(t *testing.T) {
	list := []avutil.PixelFormat{}
	src := findPixelFormatByName("rgb48be", t)
	expectedBest := avutil.PixelFormatNone
	loss := avutil.LossFlagAlpha
	best := FindBestPixelFormat(list, src, false)
	if best != expectedBest {
		t.Fatalf("[TestFindBestPixelFormat] best=%d, NG expected=%d", best, expectedBest)
	}

	list = []avutil.PixelFormat{
		findPixelFormatByName("yuv420p", t),
		findPixelFormatByName("yuv444p", t),
		findPixelFormatByName("yuvj420p", t),
	}
	expectedBest = findPixelFormatByName("yuv444p", t)
	best = FindBestPixelFormat(list, src, false)
	if best != expectedBest {
		t.Fatalf("[TestFindBestPixelFormat2] best=%d, NG expected=%d", best, expectedBest)
	}

	expectedBest = findPixelFormatByName("yuv420p", t)
	expectedLoss := avutil.LossFlagResolution + avutil.LossFlagDepth + avutil.LossFlagColorspace
	best, retLoss := FindBestPixelFormatWithLossFlags(list, src, false, avutil.LossFlagChroma)
	if best != expectedBest {
		t.Fatalf("[TestFindBestPixelFormat3] best=%d, NG expected=%d", best, expectedBest)
	}
	if retLoss != expectedLoss {
		t.Fatalf("[TestFindBestPixelFormat3] loss=%d, NG expected=%d", loss, expectedLoss)
	}
}

func findPixelFormatByName(name string, t *testing.T) avutil.PixelFormat {
	pixFmt, ok := avutil.FindPixelFormatByName(name)
	if !ok {
		t.Fatalf("pixel format not found")
	}
	return pixFmt
}

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
