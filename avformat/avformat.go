package avformat

//#include <libavutil/avutil.h>
//#include <libavutil/avstring.h>
//#include <libavcodec/avcodec.h>
//#include <libavformat/avformat.h>
//#include <stdlib.h>
//
//#ifdef AVFMT_FLAG_FAST_SEEK
//#define GO_AVFMT_FLAG_FAST_SEEK AVFMT_FLAG_FAST_SEEK
//#else
//#define GO_AVFMT_FLAG_FAST_SEEK 0
//#endif
//
//static const AVStream *go_av_streams_get(const AVStream **streams, unsigned int n)
//{
//  return streams[n];
//}
//
//static AVDictionary **go_av_alloc_dicts(int length)
//{
//  size_t size = sizeof(AVDictionary*) * length;
//  return (AVDictionary**)av_malloc(size);
//}
//
//static void go_av_dicts_set(AVDictionary** arr, unsigned int n, AVDictionary *val)
//{
//  arr[n] = val;
//}
//
// size_t sizeOfAVFormatContextFilename = sizeof(((AVFormatContext *)NULL)->filename);
//
// int GO_AVFORMAT_VERSION_MAJOR = LIBAVFORMAT_VERSION_MAJOR;
// int GO_AVFORMAT_VERSION_MINOR = LIBAVFORMAT_VERSION_MINOR;
// int GO_AVFORMAT_VERSION_MICRO = LIBAVFORMAT_VERSION_MICRO;
//
//typedef int (*AVFormatContextIOOpenCallback)(struct AVFormatContext *s, AVIOContext **pb, const char *url, int flags, AVDictionary **options);
//typedef void (*AVFormatContextIOCloseCallback)(struct AVFormatContext *s, AVIOContext *pb);
//
// int exec_cb(av_format_control_message fn, AVFormatContext *s, int type, void *data, size_t data_size) {
//   return fn(s, type, data, data_size);
// }
// int interrupt_cb(void* data) {
//	return data==NULL ? 0 : *((int*)data);
// }
// void set_interrupt_cb(AVFormatContext *c) {
//	  c->interrupt_callback.callback = interrupt_cb;
//	  c->interrupt_callback.opaque = 0;
//}
//
//
// #cgo LDFLAGS: -lavformat -lavutil
import "C"

import (
	"errors"
	"strings"
	"time"
	"unsafe"

	"github.com/SpalkLtd/go-libav/avcodec"
	"github.com/SpalkLtd/go-libav/avutil"
)

var (
	ErrAllocationError     = errors.New("allocation error")
	ErrInvalidArgumentSize = errors.New("invalid argument size")
)

type IOInterruptCallback struct {
	CAVIOInterruptCB *C.AVIOInterruptCB
}

func NewIOInterruptCallbackFromC(cb unsafe.Pointer) *IOInterruptCallback {
	return &IOInterruptCallback{CAVIOInterruptCB: (*C.AVIOInterruptCB)(cb)}
}

func (ctx *Context) SetInterruptCallback() {
	C.set_interrupt_cb(ctx.CAVFormatContext)
}

type Flags int

const (
	FlagNoFile       Flags = C.AVFMT_NOFILE
	FlagNeedNumber   Flags = C.AVFMT_NEEDNUMBER
	FlagShowIDs      Flags = C.AVFMT_SHOW_IDS
	FlagGlobalHeader Flags = C.AVFMT_GLOBALHEADER
	FlagNoTimestamps Flags = C.AVFMT_NOTIMESTAMPS
	FlagGenericIndex Flags = C.AVFMT_GENERIC_INDEX
	FlagTSDiscont    Flags = C.AVFMT_TS_DISCONT
	FlagVariableFPS  Flags = C.AVFMT_VARIABLE_FPS
	FlagNoDimensions Flags = C.AVFMT_NODIMENSIONS
	FlagNoStreams    Flags = C.AVFMT_NOSTREAMS
	FlagNoBinSearch  Flags = C.AVFMT_NOBINSEARCH
	FlagNoGenSearch  Flags = C.AVFMT_NOGENSEARCH
	FlagNoByteSeek   Flags = C.AVFMT_NO_BYTE_SEEK
	FlagAllowFlush   Flags = C.AVFMT_ALLOW_FLUSH
	FlagTSNonStrict  Flags = C.AVFMT_TS_NONSTRICT
	FlagTSNegative   Flags = C.AVFMT_TS_NEGATIVE
	FlagSeekToPTS    Flags = C.AVFMT_SEEK_TO_PTS
)

type ContextFlags int

const (
	ContextFlagGenPTS         ContextFlags = C.AVFMT_FLAG_GENPTS
	ContextFlagIgnoreIndex    ContextFlags = C.AVFMT_FLAG_IGNIDX
	ContextFlagNonBlock       ContextFlags = C.AVFMT_FLAG_NONBLOCK
	ContextFlagIgnoreDTS      ContextFlags = C.AVFMT_FLAG_IGNDTS
	ContextFlagNoFillin       ContextFlags = C.AVFMT_FLAG_NOFILLIN
	ContextFlagNoParse        ContextFlags = C.AVFMT_FLAG_NOPARSE
	ContextFlagNoBuffer       ContextFlags = C.AVFMT_FLAG_NOBUFFER
	ContextFlagCustomIO       ContextFlags = C.AVFMT_FLAG_CUSTOM_IO
	ContextFlagDiscardCorrupt ContextFlags = C.AVFMT_FLAG_DISCARD_CORRUPT
	ContextFlagFlushPackets   ContextFlags = C.AVFMT_FLAG_FLUSH_PACKETS
	ContextFlagBitExact       ContextFlags = C.AVFMT_FLAG_BITEXACT
	ContextFlagMP4ALATM       ContextFlags = C.AVFMT_FLAG_MP4A_LATM
	ContextFlagSortDTS        ContextFlags = C.AVFMT_FLAG_SORT_DTS
	ContextFlagPrivOpt        ContextFlags = C.AVFMT_FLAG_PRIV_OPT
	ContextFlagKeepSideData   ContextFlags = C.AVFMT_FLAG_KEEP_SIDE_DATA
	ContextFlagFastSeek       ContextFlags = C.GO_AVFMT_FLAG_FAST_SEEK
)

type ContextExtraFlags int

const (
	ContextExtraFlagNoHeader ContextExtraFlags = C.AVFMTCTX_NOHEADER
)

type AvoidFlags int

const (
	AvoidFlagNegTSAuto            AvoidFlags = C.AVFMT_AVOID_NEG_TS_AUTO
	AvoidFlagNegTSMakeNonNegative AvoidFlags = C.AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE
	AvoidFlagNegTSMakeZero        AvoidFlags = C.AVFMT_AVOID_NEG_TS_MAKE_ZERO
)

type Disposition int

const (
	DispositionDefault         Disposition = C.AV_DISPOSITION_DEFAULT
	DispositionDub             Disposition = C.AV_DISPOSITION_DUB
	DispositionOriginal        Disposition = C.AV_DISPOSITION_ORIGINAL
	DispositionComment         Disposition = C.AV_DISPOSITION_COMMENT
	DispositionLyrics          Disposition = C.AV_DISPOSITION_LYRICS
	DispositionKaraoke         Disposition = C.AV_DISPOSITION_KARAOKE
	DispositionForced          Disposition = C.AV_DISPOSITION_FORCED
	DispositionHearingImpaired Disposition = C.AV_DISPOSITION_HEARING_IMPAIRED
	DispositionVisualImpaired  Disposition = C.AV_DISPOSITION_VISUAL_IMPAIRED
	DispositionCleanEffects    Disposition = C.AV_DISPOSITION_CLEAN_EFFECTS
	DispositionAttachedPic     Disposition = C.AV_DISPOSITION_ATTACHED_PIC
	DispositionCaptions        Disposition = C.AV_DISPOSITION_CAPTIONS
	DispositionDescriptions    Disposition = C.AV_DISPOSITION_DESCRIPTIONS
	DispositionMetadata        Disposition = C.AV_DISPOSITION_METADATA
)

type EventFlags int

const (
	EventFlagMetadataUpdated EventFlags = C.AVFMT_EVENT_FLAG_METADATA_UPDATED
)

type IOFlags int

const (
	IOFlagRead      IOFlags = C.AVIO_FLAG_READ
	IOFlagWrite     IOFlags = C.AVIO_FLAG_WRITE
	IOFlagReadWrite IOFlags = C.AVIO_FLAG_READ_WRITE
	IOFlagNonblock  IOFlags = C.AVIO_FLAG_NONBLOCK
	IOFlagDirect    IOFlags = C.AVIO_FLAG_DIRECT
)

type SeekFlags int

const (
	SeekFlagNone     SeekFlags = 0
	SeekFlagBackward SeekFlags = C.AVSEEK_FLAG_BACKWARD
	SeekFlagByte     SeekFlags = C.AVSEEK_FLAG_BYTE
	SeekFlagAny      SeekFlags = C.AVSEEK_FLAG_ANY
	SeekFlagFrame    SeekFlags = C.AVSEEK_FLAG_FRAME
)

func init() {
	RegisterAll()
}

func Version() (int, int, int) {
	return int(C.GO_AVFORMAT_VERSION_MAJOR), int(C.GO_AVFORMAT_VERSION_MINOR), int(C.GO_AVFORMAT_VERSION_MICRO)
}

func RegisterAll() {
	C.av_register_all()
}

func NetworkInit() {
	C.avformat_network_init()
}

type CodecTagList struct {
	CAVCodecTag **C.struct_AVCodecTag
}

func NewCodecTagListFromC(cCodecTag unsafe.Pointer) *CodecTagList {
	return &CodecTagList{CAVCodecTag: (**C.struct_AVCodecTag)(cCodecTag)}
}

func (ctm *CodecTagList) IDByTag(tag uint) avcodec.CodecID {
	return (avcodec.CodecID)(C.av_codec_get_id(ctm.CAVCodecTag, (C.uint)(tag)))
}

func (ctm *CodecTagList) TagByID(id avcodec.CodecID) uint {
	return (uint)(C.av_codec_get_tag(ctm.CAVCodecTag, (C.enum_AVCodecID)(id)))
}

type Input struct {
	CAVInputFormat *C.AVInputFormat
}

func FindInputByShortName(shortName string) *Input {
	cShortName := C.CString(shortName)
	defer C.free(unsafe.Pointer(cShortName))
	cInput := C.av_find_input_format(cShortName)
	if cInput == nil {
		return nil
	}
	return NewInputFromC(unsafe.Pointer(cInput))
}

func NewInputFromC(cInput unsafe.Pointer) *Input {
	return &Input{CAVInputFormat: (*C.AVInputFormat)(cInput)}
}

func (f *Input) PrivateClass() *avutil.Class {
	if f.CAVInputFormat.priv_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(f.CAVInputFormat.priv_class))
}

func (f *Input) Names() []string {
	return cStringSplit(f.CAVInputFormat.name, ",")
}

func (f *Input) LongName() string {
	str, _ := f.LongNameOk()
	return str
}

func (f *Input) LongNameOk() (string, bool) {
	return cStringToStringOk(f.CAVInputFormat.long_name)
}

func (f *Input) MimeTypes() []string {
	return cStringSplit(f.CAVInputFormat.mime_type, ",")
}

func (f *Input) Extensions() []string {
	return cStringSplit(f.CAVInputFormat.extensions, ",")
}

func (f *Input) CodecTags() *CodecTagList {
	if f.CAVInputFormat.codec_tag == nil {
		return nil
	}
	return NewCodecTagListFromC(unsafe.Pointer(f.CAVInputFormat.codec_tag))
}

func (f *Input) Flags() Flags {
	return Flags(f.CAVInputFormat.flags)
}

type ProbeData struct {
	CAVProbeData C.AVProbeData
}

func NewProbeData() *ProbeData {
	return &ProbeData{}
}

func (pd *ProbeData) Free() {
	defer C.free(unsafe.Pointer(pd.CAVProbeData.filename))
	pd.CAVProbeData.filename = nil
	defer C.av_freep(unsafe.Pointer(&pd.CAVProbeData.buf))
	pd.CAVProbeData.buf_size = 0
	defer C.free(unsafe.Pointer(pd.CAVProbeData.mime_type))
	pd.CAVProbeData.mime_type = nil
}

func (pd *ProbeData) SetFileName(fileName *string) error {
	C.free(unsafe.Pointer(pd.CAVProbeData.filename))
	if fileName == nil {
		pd.CAVProbeData.filename = nil
		return nil
	}
	pd.CAVProbeData.filename = C.CString(*fileName)
	if pd.CAVProbeData.filename == nil {
		return ErrAllocationError
	}
	return nil
}

func (pd *ProbeData) SetBuffer(buffer []byte) error {
	pd.CAVProbeData.buf_size = 0
	C.av_freep(unsafe.Pointer(&pd.CAVProbeData.buf))
	size := C.size_t(len(buffer))
	extraSize := C.size_t(C.AVPROBE_PADDING_SIZE)
	buf := C.av_malloc(size + extraSize)
	if buf == nil {
		return ErrAllocationError
	}
	if size != 0 {
		C.memcpy(buf, unsafe.Pointer(&buffer[0]), size)
	}
	C.memset(unsafe.Pointer(uintptr(buf)+uintptr(size)), 0, extraSize)
	pd.CAVProbeData.buf = (*C.uchar)(buf)
	pd.CAVProbeData.buf_size = C.int(size)
	return nil
}

func (pd *ProbeData) SetMimeType(mimeType *string) error {
	C.free(unsafe.Pointer(pd.CAVProbeData.mime_type))
	if mimeType == nil {
		pd.CAVProbeData.mime_type = nil
		return nil
	}
	pd.CAVProbeData.mime_type = C.CString(*mimeType)
	if pd.CAVProbeData.mime_type == nil {
		return ErrAllocationError
	}
	return nil
}

func ProbeInput(pd *ProbeData, isOpened bool) *Input {
	cInput := C.av_probe_input_format(&pd.CAVProbeData, boolToCInt(isOpened))
	if cInput == nil {
		return nil
	}
	return NewInputFromC(unsafe.Pointer(cInput))
}

func ProbeInputWithScore(pd *ProbeData, isOpened bool, scoreMax int) (*Input, int) {
	cscoreMax := C.int(scoreMax)
	cInput := C.av_probe_input_format2(&pd.CAVProbeData, boolToCInt(isOpened), &cscoreMax)
	if cInput == nil {
		return nil, 0
	}
	return NewInputFromC(unsafe.Pointer(cInput)), int(cscoreMax)
}

type Output struct {
	CAVOutputFormat *C.AVOutputFormat
}

func NewOutputFromC(cOutput unsafe.Pointer) *Output {
	return &Output{CAVOutputFormat: (*C.AVOutputFormat)(cOutput)}
}

func (f *Output) QueryCodec(codecID avcodec.CodecID) (bool, bool) {
	return f.QueryCodecWithCompliance(codecID, avcodec.ComplianceNormal)
}

func (f *Output) QueryCodecWithCompliance(codecID avcodec.CodecID, compliance avcodec.Compliance) (bool, bool) {
	res := C.avformat_query_codec(f.CAVOutputFormat, (C.enum_AVCodecID)(codecID), (C.int)(compliance))
	switch res {
	case 0:
		return false, true
	case 1:
		return true, true
	default:
		return false, false
	}
}

func (f *Output) PrivateClass() *avutil.Class {
	if f.CAVOutputFormat.priv_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(f.CAVOutputFormat.priv_class))
}

func (f *Output) Name() string {
	str, _ := f.NameOk()
	return str
}

func (f *Output) NameOk() (string, bool) {
	return cStringToStringOk(f.CAVOutputFormat.name)
}

func (f *Output) LongName() string {
	str, _ := f.LongNameOk()
	return str
}

func (f *Output) LongNameOk() (string, bool) {
	return cStringToStringOk(f.CAVOutputFormat.long_name)
}

func (f *Output) MimeType() string {
	str, _ := f.MimeTypeOk()
	return str
}

func (f *Output) MimeTypeOk() (string, bool) {
	return cStringToStringOk(f.CAVOutputFormat.mime_type)
}

func (f *Output) Extensions() []string {
	return cStringSplit(f.CAVOutputFormat.extensions, ",")
}

func (f *Output) AudioCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(f.CAVOutputFormat.audio_codec)
}

func (f *Output) VideoCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(f.CAVOutputFormat.video_codec)
}

func (f *Output) SubtitleCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(f.CAVOutputFormat.subtitle_codec)
}

func (f *Output) Flags() Flags {
	return Flags(f.CAVOutputFormat.flags)
}

func (f *Output) CodecTags() *CodecTagList {
	if f.CAVOutputFormat.codec_tag == nil {
		return nil
	}
	return NewCodecTagListFromC(unsafe.Pointer(f.CAVOutputFormat.codec_tag))
}

func (f *Output) DataCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(f.CAVOutputFormat.data_codec)
}

func (f *Output) GuessCodecID(filename string, mediaType avutil.MediaType) avcodec.CodecID {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	return (avcodec.CodecID)(C.av_guess_codec(f.CAVOutputFormat, nil, cFilename, nil, C.enum_AVMediaType(mediaType)))
}

func GuessOutputFromShortName(shortName string) *Output {
	cShortName := C.CString(shortName)
	defer C.free(unsafe.Pointer(cShortName))
	cOutput := C.av_guess_format(cShortName, nil, nil)
	if cOutput == nil {
		return nil
	}
	return NewOutputFromC(unsafe.Pointer(cOutput))
}

func GuessOutputFromFileName(fileName string) *Output {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))
	cOutput := C.av_guess_format(nil, cFileName, nil)
	if cOutput == nil {
		return nil
	}
	return NewOutputFromC(unsafe.Pointer(cOutput))
}

func GuessOutputFromMimeType(mimeType string) *Output {
	cMimeType := C.CString(mimeType)
	defer C.free(unsafe.Pointer(cMimeType))
	cOutput := C.av_guess_format(nil, nil, cMimeType)
	if cOutput == nil {
		return nil
	}
	return NewOutputFromC(unsafe.Pointer(cOutput))
}

type Stream struct {
	CAVStream *C.AVStream
}

func NewStreamFromC(cStream unsafe.Pointer) *Stream {
	return &Stream{CAVStream: (*C.AVStream)(cStream)}
}

func (s *Stream) Index() int {
	return int(s.CAVStream.index)
}

func (s *Stream) ID() int {
	return int(s.CAVStream.id)
}

func (s *Stream) CodecContext() *avcodec.Context {
	if s.CAVStream.codec == nil {
		return nil
	}
	return avcodec.NewContextFromC(unsafe.Pointer(s.CAVStream.codec))
}

func (s *Stream) TimeBase() *avutil.Rational {
	tb := &s.CAVStream.time_base
	return avutil.NewRational(int(tb.num), int(tb.den))
}

func (s *Stream) SetTimeBase(timeBase *avutil.Rational) {
	s.CAVStream.time_base.num = (C.int)(timeBase.Numerator())
	s.CAVStream.time_base.den = (C.int)(timeBase.Denominator())
}

func (s *Stream) SampleAspectRatio() *avutil.Rational {
	sar := &s.CAVStream.sample_aspect_ratio
	return avutil.NewRational(int(sar.num), int(sar.den))
}

func (s *Stream) SetSampleAspectRatio(aspectRatio *avutil.Rational) {
	s.CAVStream.sample_aspect_ratio.num = (C.int)(aspectRatio.Numerator())
	s.CAVStream.sample_aspect_ratio.den = (C.int)(aspectRatio.Denominator())
}

func (s *Stream) StartTime() int64 {
	return int64(s.CAVStream.start_time)
}

func (s *Stream) RawDuration() int64 {
	return int64(s.CAVStream.duration)
}

func (s *Stream) Duration() time.Duration {
	timeBase := s.TimeBase().Float64()
	return time.Duration((timeBase * float64(s.RawDuration())) * 1000 * 1000 * 1000)
}

func (s *Stream) NumberOfFrames() int64 {
	return int64(s.CAVStream.nb_frames)
}

func (s *Stream) Disposition() Disposition {
	return Disposition(s.CAVStream.disposition)
}

func (s *Stream) SetDisposition(disposition Disposition) {
	s.CAVStream.disposition = C.int(disposition)
}

func (s *Stream) MetaData() *avutil.Dictionary {
	return avutil.NewDictionaryFromC(unsafe.Pointer(&s.CAVStream.metadata))
}

func (s *Stream) SetMetaData(metaData *avutil.Dictionary) {
	var cMetaData *C.AVDictionary
	if metaData != nil {
		cMetaData = (*C.AVDictionary)(metaData.Value())
	}
	s.CAVStream.metadata = cMetaData
}

func (s *Stream) AverageFrameRate() *avutil.Rational {
	return avutil.NewRationalFromC(unsafe.Pointer(&s.CAVStream.avg_frame_rate))
}

func (s *Stream) SetAverageFrameRate(frameRate *avutil.Rational) {
	s.CAVStream.avg_frame_rate.num = (C.int)(frameRate.Numerator())
	s.CAVStream.avg_frame_rate.den = (C.int)(frameRate.Denominator())
}

func (s *Stream) RealFrameRate() *avutil.Rational {
	r := C.av_stream_get_r_frame_rate(s.CAVStream)
	return avutil.NewRationalFromC(unsafe.Pointer(&r))
}

func (s *Stream) SetRealFrameRate(frameRate *avutil.Rational) {
	s.CAVStream.r_frame_rate.num = (C.int)(frameRate.Numerator())
	s.CAVStream.r_frame_rate.den = (C.int)(frameRate.Denominator())
}

func (s *Stream) SetFirstDTS(firstDTS int64) {
	s.CAVStream.first_dts = (C.int64_t)(firstDTS)
}

func (s *Stream) FirstDTS() int64 {
	return int64(s.CAVStream.first_dts)
}

func (s *Stream) EndPTS() int64 {
	return int64(C.av_stream_get_end_pts(s.CAVStream))
}

type Context struct {
	CAVFormatContext *C.AVFormatContext
}

func NewContextForInput() (*Context, error) {
	cCtx := C.avformat_alloc_context()
	if cCtx == nil {
		return nil, ErrAllocationError
	}
	return NewContextFromC(unsafe.Pointer(cCtx)), nil
}

func NewContextForOutput(output *Output) (*Context, error) {
	var cCtx *C.AVFormatContext
	code := C.avformat_alloc_output_context2(&cCtx, output.CAVOutputFormat, nil, nil)
	if code < 0 {
		return nil, avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return NewContextFromC(unsafe.Pointer(cCtx)), nil
}

func NewContextFromC(cCtx unsafe.Pointer) *Context {
	ctx := Context{
		CAVFormatContext: (*C.AVFormatContext)(cCtx),
	}
	ctx.SetInterruptCallback()
	return &ctx
}

func (ctx *Context) Free() {
	if ctx.CAVFormatContext != nil {
		defer C.avformat_free_context(ctx.CAVFormatContext)
		ctx.CAVFormatContext = nil
	}
}

func (ctx *Context) Class() *avutil.Class {
	if ctx.CAVFormatContext.av_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(ctx.CAVFormatContext.av_class))
}

func (ctx *Context) Input() *Input {
	if ctx.CAVFormatContext.iformat == nil {
		return nil
	}
	return NewInputFromC(unsafe.Pointer(ctx.CAVFormatContext.iformat))
}

func (ctx *Context) Output() *Output {
	if ctx.CAVFormatContext.oformat == nil {
		return nil
	}
	return NewOutputFromC(unsafe.Pointer(ctx.CAVFormatContext.oformat))
}

func (ctx *Context) IOContext() *IOContext {
	if ctx.CAVFormatContext.pb == nil {
		return nil
	}
	return NewIOContextFromC(unsafe.Pointer(ctx.CAVFormatContext.pb))
}

func (ctx *Context) SetIOContext(ioCtx *IOContext) {
	var cIOCtx *C.AVIOContext
	if ioCtx != nil {
		cIOCtx = ioCtx.CAVIOContext
	}
	ctx.CAVFormatContext.pb = cIOCtx
}

func (ctx *Context) NewStream() (*Stream, error) {
	return ctx.NewStreamWithCodec(nil)
}

func (ctx *Context) NewStreamWithCodec(codec *avcodec.Codec) (*Stream, error) {
	var cCodec *C.AVCodec
	if codec != nil {
		cCodec = (*C.AVCodec)(unsafe.Pointer(codec.CAVCodec))
	}
	cStream := C.avformat_new_stream(ctx.CAVFormatContext, cCodec)
	if cStream == nil {
		return nil, ErrAllocationError
	}
	return NewStreamFromC(unsafe.Pointer(cStream)), nil
}

func (ctx *Context) NumberOfStreams() uint {
	return uint(ctx.CAVFormatContext.nb_streams)
}

func (ctx *Context) WriteHeader(options *avutil.Dictionary) error {
	var cOptions **C.AVDictionary
	if options != nil {
		cOptions = (**C.AVDictionary)(options.Pointer())
	}
	code := C.avformat_write_header(ctx.CAVFormatContext, cOptions)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) WriteTrailer() error {
	code := C.av_write_trailer(ctx.CAVFormatContext)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) ReadFrame() (*avcodec.Packet, error) {
	pkt, err := avcodec.NewPacket()
	if err != nil {
		return nil, err
	}
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	code := C.av_read_frame(ctx.CAVFormatContext, cPkt)
	if code < 0 {
		// if avutil.ErrorCode(code) == avutil.ErrorCodeEOF {
		// 	return nil, nil
		// }
		return nil, avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return pkt, nil
}

/**
* Start playing a network-based stream (e.g. RTSP stream) at the
* current position.
 */
func (ctx *Context) ReadPlay() error {

	code := C.av_read_play(ctx.CAVFormatContext)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

/**
* Pause playing a network-based stream (e.g. RTSP stream) at the
* current position.
 */
func (ctx *Context) ReadPause() error {

	code := C.av_read_pause(ctx.CAVFormatContext)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) WriteFrame(pkt *avcodec.Packet) error {
	var cPkt *C.AVPacket
	if cPkt != nil {
		cPkt = (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	}
	code := C.av_write_frame(ctx.CAVFormatContext, cPkt)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) InterleavedWriteFrame(pkt *avcodec.Packet) error {
	var cPkt *C.AVPacket
	if pkt != nil {
		cPkt = (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	}
	code := C.av_interleaved_write_frame(ctx.CAVFormatContext, cPkt)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) Streams() []*Stream {
	count := ctx.NumberOfStreams()
	if count <= 0 {
		return nil
	}
	streams := make([]*Stream, 0, count)
	for i := uint(0); i < count; i++ {
		cStream := C.go_av_streams_get(ctx.CAVFormatContext.streams, C.uint(i))
		stream := NewStreamFromC(unsafe.Pointer(cStream))
		streams = append(streams, stream)
	}
	return streams
}

func (ctx *Context) FileName() string {
	return C.GoString(&ctx.CAVFormatContext.filename[0])
}

func (ctx *Context) SetFileName(fileName string) {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))
	C.av_strlcpy(&ctx.CAVFormatContext.filename[0], cFileName, C.sizeOfAVFormatContextFilename)
}

func (ctx *Context) StartTime() int64 {
	return int64(ctx.CAVFormatContext.start_time)
}

func (ctx *Context) Duration() int64 {
	return int64(ctx.CAVFormatContext.duration)
}

func (ctx *Context) SetDuration(duration int64) {
	ctx.CAVFormatContext.duration = (C.int64_t)(duration)
}

func (ctx *Context) BitRate() int64 {
	return int64(ctx.CAVFormatContext.bit_rate)
}

func (ctx *Context) MaxDelay() int {
	return int(ctx.CAVFormatContext.max_delay)
}

func (ctx *Context) SetMaxDelay(maxDelay int) {
	ctx.CAVFormatContext.max_delay = (C.int)(maxDelay)
}

func (ctx *Context) Flags() ContextFlags {
	return ContextFlags(ctx.CAVFormatContext.flags)
}

func (ctx *Context) SetFlags(flags ContextFlags) {
	ctx.CAVFormatContext.flags = (C.int)(flags)
}

func (ctx *Context) ExtraFlags() ContextExtraFlags {
	return ContextExtraFlags(ctx.CAVFormatContext.ctx_flags)
}

func (ctx *Context) AudioCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(ctx.CAVFormatContext.audio_codec_id)
}

func (ctx *Context) VideoCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(ctx.CAVFormatContext.video_codec_id)
}

func (ctx *Context) SubtitleCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(ctx.CAVFormatContext.subtitle_codec_id)
}

func (ctx *Context) MetaData() *avutil.Dictionary {
	return avutil.NewDictionaryFromC(unsafe.Pointer(&ctx.CAVFormatContext.metadata))
}

func (ctx *Context) SetMetaData(metaData *avutil.Dictionary) {
	var cMetaData *C.AVDictionary
	if metaData != nil {
		cMetaData = (*C.AVDictionary)(metaData.Value())
	}
	ctx.CAVFormatContext.metadata = cMetaData
}

func (ctx *Context) DataCodecID() avcodec.CodecID {
	return (avcodec.CodecID)(ctx.CAVFormatContext.data_codec_id)
}

func (ctx *Context) IOOpenCallback() unsafe.Pointer {
	return unsafe.Pointer(ctx.CAVFormatContext.io_open)
}

func (ctx *Context) SetIOOpenCallback(callback unsafe.Pointer) {
	ctx.CAVFormatContext.io_open = (C.AVFormatContextIOOpenCallback)(callback)
}

func (ctx *Context) IOCloseCallback() unsafe.Pointer {
	return unsafe.Pointer(ctx.CAVFormatContext.io_close)
}

func (ctx *Context) SetIOCloseCallback(callback unsafe.Pointer) {
	ctx.CAVFormatContext.io_close = (C.AVFormatContextIOCloseCallback)(callback)
}

func (ctx *Context) OpenInput(fileName string, input *Input, options *avutil.Dictionary) error {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))
	var cInput *C.AVInputFormat
	if input != nil {
		cInput = input.CAVInputFormat
	}
	var cOptions **C.AVDictionary
	if options != nil {
		cOptions = (**C.AVDictionary)(options.Pointer())
	}
	code := C.avformat_open_input(&ctx.CAVFormatContext, cFileName, cInput, cOptions)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) CloseInput() {
	C.avformat_close_input(&ctx.CAVFormatContext)
}

func (ctx *Context) FindStreamInfo(options []*avutil.Dictionary) error {
	var cOptions **C.AVDictionary
	count := ctx.NumberOfStreams()
	if count > 0 && options != nil {
		if uint(len(options)) < count {
			return ErrInvalidArgumentSize
		}
		cOptions = newCAVDictionaryArrayFromDictionarySlice(options[:count])
		defer freeCAVDictionaryArray(cOptions)
	}
	code := C.avformat_find_stream_info(ctx.CAVFormatContext, cOptions)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) Dump(streamIndex int, url string, isOutput bool) {
	var cIsOutput C.int
	if isOutput {
		cIsOutput = C.int(1)
	}
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))
	C.av_dump_format(ctx.CAVFormatContext, C.int(streamIndex), cURL, cIsOutput)
}

func (ctx *Context) GuessFrameRate(stream *Stream, frame *avutil.Frame) *avutil.Rational {
	cStream := (*C.AVStream)(unsafe.Pointer(stream.CAVStream))
	var cFrame *C.AVFrame
	if frame != nil {
		cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	}
	r := C.av_guess_frame_rate(ctx.CAVFormatContext, cStream, cFrame)
	return avutil.NewRationalFromC(unsafe.Pointer(&r))
}

func (ctx *Context) SeekToTimestamp(streamIndex int, min, target, max int64, flags SeekFlags) error {
	code := C.avformat_seek_file(ctx.CAVFormatContext, C.int(streamIndex), C.int64_t(min), C.int64_t(target), C.int64_t(max), C.int(flags))
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) ControlMessage(msg int, data interface{}) (int64, error) {
	if ctx.Output() == nil {
		return 0, errors.New("No output found")
	}
	var cData C.int64_t
	pointer := unsafe.Pointer(nil)
	if data == nil {
		data = int64(0)
	}
	//Convert data to an unsafe pointer
	i64Data, ok := data.(int64)
	if !ok {
		return 0, errors.New("Data is not an int64")
	}
	cData = C.int64_t(i64Data)
	pointer = unsafe.Pointer(&cData)

	code := C.exec_cb(ctx.Output().CAVOutputFormat.control_message, ctx.CAVFormatContext, C.int(msg), pointer, 0)
	if code < 0 {
		return 0, avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return int64(cData), nil
}

func (ctx *Context) InterruptBlockingOperation() {
	data := C.int(1)
	ctx.CAVFormatContext.interrupt_callback.opaque = unsafe.Pointer(&data)
}

func (ctx *Context) UninterruptBlockingOperation() {
	data := C.int(0)
	ctx.CAVFormatContext.interrupt_callback.opaque = unsafe.Pointer(&data)
}

func (ctx *Context) GetOutputTimestamp(streamIdx int) (int, int, error) {
	var dts C.int64_t
	var wall C.int64_t
	cStreamIdx := C.int(streamIdx)
	code := C.av_get_output_timestamp(ctx.CAVFormatContext, cStreamIdx,
		&dts, &wall)
	if code < 0 {
		return 0, 0, avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return int(dts), int(wall), nil
}

type IOContext struct {
	CAVIOContext *C.AVIOContext
}

func OpenIOContext(url string, flags IOFlags, cb *IOInterruptCallback, options *avutil.Dictionary) (*IOContext, error) {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))
	var cCb *C.AVIOInterruptCB
	if cb != nil {
		cCb = cb.CAVIOInterruptCB
	}
	var cOptions **C.AVDictionary
	if options != nil {
		cOptions = (**C.AVDictionary)(options.Pointer())
	}
	var cCtx *C.AVIOContext
	code := C.avio_open2(&cCtx, cURL, (C.int)(flags), cCb, cOptions)
	if code < 0 {
		return nil, avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return NewIOContextFromC(unsafe.Pointer(cCtx)), nil
}

func NewIOContextFromC(cCtx unsafe.Pointer) *IOContext {
	return &IOContext{CAVIOContext: (*C.AVIOContext)(cCtx)}
}

func (ctx *IOContext) Size() int64 {
	return int64(C.avio_size(ctx.CAVIOContext))
}

func (ctx *IOContext) Close() error {
	if ctx.CAVIOContext != nil {
		code := C.avio_closep(&ctx.CAVIOContext)
		if code < 0 {
			return avutil.NewErrorFromCode(avutil.ErrorCode(code))
		}
	}
	return nil
}

func (ctx *IOContext) Write(packet unsafe.Pointer, size int) {
	cSize := C.int(size)
	C.avio_write(ctx.CAVIOContext, (*C.uchar)(packet), cSize)
}

func (ctx *IOContext) WriteBytes(b []byte) {
	buf := C.CBytes(b)
	ctx.Write(buf, len(b))
	C.free(buf)
}

func (ctx *IOContext) Error() error {
	code := (*ctx.CAVIOContext).error
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(int(code)))
	}
	return nil
}

func (ctx *IOContext) Flush() {
	C.avio_flush(ctx.CAVIOContext)
}

func boolToCInt(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func cStringSplit(cStr *C.char, sep string) []string {
	str, ok := cStringToStringOk(cStr)
	if !ok {
		return nil
	}
	return strings.Split(str, sep)
}

func cStringToStringOk(cStr *C.char) (string, bool) {
	if cStr == nil {
		return "", false
	}
	return C.GoString(cStr), true
}

func newCAVDictionaryArrayFromDictionarySlice(dicts []*avutil.Dictionary) **C.AVDictionary {
	arr := C.go_av_alloc_dicts(C.int(len(dicts)))
	for i := range dicts {
		C.go_av_dicts_set(arr, C.uint(i), (*C.AVDictionary)(dicts[i].Value()))
	}
	return nil
}

func freeCAVDictionaryArray(arr **C.AVDictionary) {
	C.av_free(unsafe.Pointer(arr))
}

func NumberedSequenceFormat(format string) bool {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	valid := C.av_filename_number_test(cFormat)
	if valid == 1 {
		return true
	}
	return false
}

func FormatNumberedSequence(format string, num int) (string, error) {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	size := C.size_t(1024)
	buf := (*C.char)(C.av_mallocz(size))
	defer C.av_free(unsafe.Pointer(buf))
	code := C.av_get_frame_filename(buf, C.int(size-1), cFormat, C.int(num))
	if code < 0 {
		return "", avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return C.GoString(buf), nil
}
