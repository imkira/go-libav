package avformat

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/imkira/go-libav/avutil"
)

func testInputFormatMatroska(t *testing.T, f *Input) {
	if f == nil {
		t.Fatalf("Expecting format")
	}
	names := f.Names()
	if !reflect.DeepEqual(names, []string{"matroska", "webm"}) {
		t.Fatalf("Expecting names but got %v", names)
	}
	longName, ok := f.LongNameOk()
	if !ok || longName != "Matroska / WebM" {
		t.Fatalf("Expecting name but got %s", longName)
	}
	mimeTypes := f.MimeTypes()
	if !reflect.DeepEqual(mimeTypes, []string{"audio/webm", "audio/x-matroska", "video/webm", "video/x-matroska"}) {
		t.Fatalf("Expecting mimeTypes but got %v", mimeTypes)
	}
	extensions := f.Extensions()
	if !reflect.DeepEqual(extensions, []string{"mkv", "mk3d", "mka", "mks"}) {
		t.Fatalf("Expecting extensions but got %v", extensions)
	}
}

func TestFindInputByShortName(t *testing.T) {
	shortNames := []string{
		"matroska",
	}
	for _, shortName := range shortNames {
		f := FindInputByShortName(shortName)
		testInputFormatMatroska(t, f)
	}
	if FindInputByShortName("maaaaatroska") != nil {
		t.Fatalf("Not expecting format")
	}
}

func TestInputFlags(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	fixture := fixturePath("sample_mpeg4.mp4")
	err := ctx.OpenInput(fixture, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	result := ctx.Input().Flags()
	if result != FlagNoByteSeek {
		t.Fatalf("[TestFlags] result = %v, NG, expected = %v", result, FlagNoByteSeek)
	}
}

func TestProbeInput(t *testing.T) {
	pd := NewProbeData()
	defer pd.Free()
	if ProbeInput(pd, true) != nil {
		t.Fatalf("Not expecting format")
	}

	pd = NewProbeData()
	defer pd.Free()
	pd.SetFileName("file.mkv")
	testInputFormatMatroska(t, ProbeInput(pd, true))

	pd = NewProbeData()
	defer pd.Free()
	pd.SetMimeType("video/x-matroska")
	testInputFormatMatroska(t, ProbeInput(pd, true))

	pd = NewProbeData()
	defer pd.Free()
	matroskaHeader := []byte{
		0x1a, 0x45, 0xdf, 0xa3, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0x42, 0x82, 0x89, 0x6d,
		0x61, 0x74, 0x72, 0x6f, 0x73, 0x6b, 0x61, 0x00, 0x42, 0x87, 0x81, 0x02, 0x42, 0x85, 0x81, 0x02,
	}
	pd.SetBuffer(matroskaHeader)
	testInputFormatMatroska(t, ProbeInput(pd, true))

	pd = NewProbeData()
	defer pd.Free()
	pd.SetBuffer(matroskaHeader)
	f, score := ProbeInputWithScore(pd, true, 0)
	testInputFormatMatroska(t, f)
	if score != 100 {
		t.Fatalf("Expecting score but got %d", score)
	}

	pd = NewProbeData()
	defer pd.Free()
	pd.SetBuffer(nil)
	if ProbeInput(pd, true) != nil {
		t.Fatalf("Not expecting format")
	}
}

func testOutputFormatMatroska(t *testing.T, f *Output) {
	if f == nil {
		t.Fatalf("Expecting format")
	}
	name, ok := f.NameOk()
	if !ok || name != "matroska" {
		t.Fatalf("Expecting name but got %s", name)
	}
	longName, ok := f.LongNameOk()
	if !ok || longName != "Matroska" {
		t.Fatalf("Expecting name but got %s", longName)
	}
	mimeType, ok := f.MimeTypeOk()
	if !ok || mimeType != "video/x-matroska" {
		t.Fatalf("Expecting mimetype but got %s", mimeType)
	}
	extensions := f.Extensions()
	if !reflect.DeepEqual(extensions, []string{"mkv"}) {
		t.Fatalf("Expecting extensions but got %v", extensions)
	}
}

func TestGuessOutputFromShortName(t *testing.T) {
	shortNames := []string{
		"matroska",
		"MATROSKA",
	}
	for _, shortName := range shortNames {
		f := GuessOutputFromShortName(shortName)
		testOutputFormatMatroska(t, f)
	}
	if GuessOutputFromShortName("maaaaatroska") != nil {
		t.Fatalf("Not expecting format")
	}
}

func TestGuessOutputFromFileName(t *testing.T) {
	fileNames := []string{
		"test.mkv",
		"test.MKV",
		"file://test.mkv",
		"http://example.com/test.mkv",
	}
	for _, fileName := range fileNames {
		f := GuessOutputFromFileName(fileName)
		testOutputFormatMatroska(t, f)
	}
	if GuessOutputFromFileName("maaaaatroska") != nil {
		t.Fatalf("Not expecting format")
	}
}

func TestGuessOutputFromMimeType(t *testing.T) {
	mimeTypes := []string{
		"video/x-matroska",
	}
	for _, mimeType := range mimeTypes {
		f := GuessOutputFromMimeType(mimeType)
		if f == nil {
			t.Fatalf("Expecting format")
		}
		testOutputFormatMatroska(t, f)
	}
	if GuessOutputFromMimeType("video/x-maaaaatroska") != nil {
		t.Fatalf("Not expecting format")
	}
}

func TestNewContextForInput(t *testing.T) {
	ctx, err := NewContextForInput()
	if err != nil || ctx == nil {
		t.Fatalf("Expecting context")
	}
	defer ctx.Free()
}

func TestNewContextForOutput(t *testing.T) {
	output := GuessOutputFromShortName("matroska")
	if output == nil {
		t.Fatalf("Expecting output")
	}
	ctx, err := NewContextForOutput(output)
	if err != nil || ctx == nil {
		t.Fatalf("Expecting context")
	}
	defer ctx.Free()
}

func TestContextOpenInputNonExistent(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	err := ctx.OpenInput("foobarnonexistent", nil, nil).(*avutil.Error)
	if err == nil {
		defer ctx.CloseInput()
		t.Fatal(err)
	}
	if err.Error() != "No such file or directory" {
		t.Fatal(err)
	}
}

func TestContextOpenInputExistent(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	fixture := fixturePath("sample_mpeg4.mp4")
	err := ctx.OpenInput(fixture, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.CloseInput()
}

func TestContextOpenInputWithOptions(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	fixture := fixturePath("sample_mpeg4.mp4")

	options := avutil.NewDictionary()
	defer options.Free()
	options.Set("foo", "1")
	options.Set("export_all", "1")
	options.Set("bar", "1")

	err := ctx.OpenInput(fixture, nil, options)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.CloseInput()

	// consumed options disappear from the dictionary
	m := options.Map()
	if !reflect.DeepEqual(m, map[string]string{"foo": "1", "bar": "1"}) {
		t.Fatalf("Expecting map but got %v", m)
	}
}

func fixturePath(elem ...string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir = filepath.Join(filepath.Dir(dir), "fixtures")
	path, err := filepath.Abs(filepath.Join(dir, filepath.Join(elem...)))
	if err != nil {
		panic(err)
	}
	return path
}

func TestSetFileName(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()

	var buff bytes.Buffer
	for i := 0; i < 1023; i++ {
		buff.WriteRune('a')
		ctx.SetFileName(buff.String())
		result := ctx.FileName()
		if result != buff.String() {
			t.Fatalf("[TestSetFileName] result = %s, NG, expected = %s", result, buff.String())
		}
	}
	buff.WriteRune('a')
	ctx.SetFileName(buff.String())
	result := ctx.FileName()
	if result != string(buff.Bytes()[:1023]) {
		t.Fatalf("[TestSetFileName] result = %s, NG, expected = %s", result, buff.String())
	}
}

func TestContextSeekToTimestamp(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	fixture := fixturePath("sample_mpeg4.mp4")
	err := ctx.OpenInput(fixture, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	start := ctx.StartTime()
	if err := ctx.SeekToTimestamp(-1, -9223372036854775808, start, start, SeekFlagNone); err != nil {
		t.Fatalf("[TestSeekToTimestamp] result(error) = %v, NG, expected no error", err)
	}
}

func TestSampleAspectRatio(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	stream, err := ctx.NewStream()
	if err != nil {
		t.Fatal(err)
	}
	expected := avutil.NewRational(1, 5)
	stream.SetSampleAspectRatio(expected)
	result := stream.SampleAspectRatio()
	if result.Numerator() != expected.Numerator() || result.Denominator() != expected.Denominator() {
		t.Fatalf("[TestSampleAspectRatio] result = %d/%d, NG, expected = %d/%d",
			result.Numerator(), result.Denominator(), expected.Numerator(), expected.Denominator())
	}
}

func TestRealFrameRate(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	stream, err := ctx.NewStream()
	if err != nil {
		t.Fatal(err)
	}
	expected := avutil.NewRational(30, 1)
	stream.SetRealFrameRate(expected)
	result := stream.RealFrameRate()
	if result.Numerator() != expected.Numerator() || result.Denominator() != expected.Denominator() {
		t.Fatalf("[TestRealFrameRate] result = %d/%d, NG, expected = %d/%d",
			result.Numerator(), result.Denominator(), expected.Numerator(), expected.Denominator())
	}
}

func TestStreamFirstDTSOK(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()
	stream, err := ctx.NewStream()
	if err != nil {
		t.Fatal(err)
	}
	expected := int64(1500000)
	stream.SetFirstDTS(expected)
	result := stream.FirstDTS()
	if result != expected {
		t.Fatalf("[TestStreamFirstDTSOK] result = %d, NG, expected = %d", result, expected)
	}
}

func TestGuessFrameRate(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()

	fixture := fixturePath("sample_mpeg4.mp4")
	if err := ctx.OpenInput(fixture, nil, nil); err != nil {
		t.Fatal(err)
	}
	defer ctx.CloseInput()

	if err := ctx.FindStreamInfo(nil); err != nil {
		t.Fatal(err)
	}

	if ctx.BitRate() <= 0 {
		t.Fatalf("[TestGuessFrameRate] bitrate result = %d, NG, expected greater than 0", ctx.BitRate())
	}

	expected := [][]int{{0, 0}, {30, 1}}
	for i, stream := range ctx.Streams() {
		result := ctx.GuessFrameRate(stream, nil)
		if result.Numerator() != expected[i][0] || result.Denominator() != expected[i][1] {
			t.Fatalf("[TestGuessFrameRate] result = %d/%d, NG, expected = %d/%d",
				result.Numerator(), result.Denominator(), expected[i][0], expected[i][1])
		}
	}
}

func TestContextDuration(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()

	ctx.SetDuration(1000000)
	result := ctx.Duration()
	if result != 1000000 {
		t.Fatalf("[TestContextDuration] result = %d, NG, expected = %d", result, 1000000)
	}
}

func TestContextMaxDelay(t *testing.T) {
	ctx, _ := NewContextForInput()
	defer ctx.Free()

	ctx.SetMaxDelay(500000)
	result := ctx.MaxDelay()
	if result != 500000 {
		t.Fatalf("[TestContextMaxDelay] result = %d, NG, expected = %d", result, 500000)
	}
}

func TestContextMetaData(t *testing.T) {
	fmtCtx, err := NewContextForInput()
	if err != nil {
		t.Fatal(err)
	}
	defer fmtCtx.Free()
	metadata := fmtCtx.MetaData()
	if count := metadata.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if err := metadata.Set("foo", "foo"); err != nil {
		t.Fatal(err)
	}
	if count := metadata.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if count := fmtCtx.MetaData().Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if err := metadata.Delete("foo"); err != nil {
		t.Fatal(err)
	}
	if count := metadata.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if err := metadata.Set("bar", "bar"); err != nil {
		t.Fatal(err)
	}
	if count := metadata.Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if count := fmtCtx.MetaData().Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
	if err := metadata.Delete("bar"); err != nil {
		t.Fatal(err)
	}
	if count := metadata.Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
}

func TestContextSetMetaData(t *testing.T) {
	fmtCtx, err := NewContextForInput()
	if err != nil {
		t.Fatal(err)
	}
	defer fmtCtx.Free()
	if count := fmtCtx.MetaData().Count(); count != 0 {
		t.Fatalf("Expecting count but got %d", count)
	}
	metadata := avutil.NewDictionary()
	if err := metadata.Set("foo", "foo"); err != nil {
		t.Fatal(err)
	}
	fmtCtx.SetMetaData(metadata)
	if count := fmtCtx.MetaData().Count(); count != 1 {
		t.Fatalf("Expecting count but got %d", count)
	}
}
