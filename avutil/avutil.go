package avutil

//go:generate go run hackgenerator.go

//#include <libavutil/avutil.h>
//#include <libavutil/channel_layout.h>
//#include <libavutil/dict.h>
//#include <libavutil/pixdesc.h>
//#include <libavutil/opt.h>
//#include <libavutil/frame.h>
//#include <libavutil/parseutils.h>
//#include <libavutil/common.h>
//#include <libavutil/eval.h>
//
//#ifdef AV_LOG_TRACE
//#define GO_AV_LOG_TRACE AV_LOG_TRACE
//#else
//#define GO_AV_LOG_TRACE AV_LOG_DEBUG
//#endif
//
//#ifdef AV_PIX_FMT_XVMC_MPEG2_IDCT
//#define GO_AV_PIX_FMT_XVMC_MPEG2_IDCT AV_PIX_FMT_XVMC_MPEG2_MC
//#else
//#define GO_AV_PIX_FMT_XVMC_MPEG2_IDCT 0
//#endif
//
//#ifdef AV_PIX_FMT_XVMC_MPEG2_MC
//#define GO_AV_PIX_FMT_XVMC_MPEG2_MC AV_PIX_FMT_XVMC_MPEG2_MC
//#else
//#define GO_AV_PIX_FMT_XVMC_MPEG2_MC 0
//#endif
//
//static const AVDictionaryEntry *go_av_dict_next(const AVDictionary *m, const AVDictionaryEntry *prev)
//{
//  return av_dict_get(m, "", prev, AV_DICT_IGNORE_SUFFIX);
//}
//
//static const int go_av_dict_has(const AVDictionary *m, const char *key, int flags)
//{
//  if (av_dict_get(m, key, NULL, flags) != NULL)
//  {
//    return 1;
//  }
//  return 0;
//}
//
//static int go_av_expr_parse2(AVExpr **expr, const char *s, const char * const *const_names, int log_offset, void *log_ctx)
//{
//  return av_expr_parse(expr, s, const_names, NULL, NULL, NULL, NULL, log_offset, log_ctx);
//}
//
//static const int go_av_errno_to_error(int e)
//{
//  return AVERROR(e);
//}
//
// #cgo pkg-config: libavutil
import "C"

import (
	"errors"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var (
	ErrAllocationError     = errors.New("allocation error")
	ErrInvalidArgumentSize = errors.New("invalid argument size")
)

type LogLevel int

const (
	LogLevelQuiet   LogLevel = C.AV_LOG_QUIET
	LogLevelPanic   LogLevel = C.AV_LOG_PANIC
	LogLevelFatal   LogLevel = C.AV_LOG_FATAL
	LogLevelError   LogLevel = C.AV_LOG_ERROR
	LogLevelWarning LogLevel = C.AV_LOG_WARNING
	LogLevelInfo    LogLevel = C.AV_LOG_INFO
	LogLevelVerbose LogLevel = C.AV_LOG_VERBOSE
	LogLevelDebug   LogLevel = C.AV_LOG_DEBUG
	LogLevelTrace   LogLevel = C.GO_AV_LOG_TRACE
)

type MediaType C.enum_AVMediaType

const (
	MediaTypeUnknown    MediaType = C.AVMEDIA_TYPE_UNKNOWN
	MediaTypeVideo      MediaType = C.AVMEDIA_TYPE_VIDEO
	MediaTypeAudio      MediaType = C.AVMEDIA_TYPE_AUDIO
	MediaTypeData       MediaType = C.AVMEDIA_TYPE_DATA
	MediaTypeSubtitle   MediaType = C.AVMEDIA_TYPE_SUBTITLE
	MediaTypeAttachment MediaType = C.AVMEDIA_TYPE_ATTACHMENT
)

type PictureType C.enum_AVPictureType

const (
	PictureTypeNone PictureType = C.AV_PICTURE_TYPE_NONE
	PictureTypeI    PictureType = C.AV_PICTURE_TYPE_I
	PictureTypeP    PictureType = C.AV_PICTURE_TYPE_P
	PictureTypeB    PictureType = C.AV_PICTURE_TYPE_B
	PictureTypeS    PictureType = C.AV_PICTURE_TYPE_S
	PictureTypeSI   PictureType = C.AV_PICTURE_TYPE_SI
	PictureTypeSP   PictureType = C.AV_PICTURE_TYPE_SP
	PictureTypeBI   PictureType = C.AV_PICTURE_TYPE_BI
)

type ChromaLocation C.enum_AVChromaLocation

const (
	ChromaLocationUnspecified ChromaLocation = C.AVCHROMA_LOC_UNSPECIFIED
	ChromaLocationLeft        ChromaLocation = C.AVCHROMA_LOC_LEFT
	ChromaLocationCenter      ChromaLocation = C.AVCHROMA_LOC_CENTER
	ChromaLocationTopLeft     ChromaLocation = C.AVCHROMA_LOC_TOPLEFT
	ChromaLocationTop         ChromaLocation = C.AVCHROMA_LOC_TOP
	ChromaLocationBottomLeft  ChromaLocation = C.AVCHROMA_LOC_BOTTOMLEFT
	ChromaLocationBottom      ChromaLocation = C.AVCHROMA_LOC_BOTTOM
)

type ErrorCode int

type OptionSearchFlags int

const (
	OptionSearchChildren OptionSearchFlags = C.AV_OPT_SEARCH_CHILDREN
	OptionSearchFakeObj  OptionSearchFlags = C.AV_OPT_SEARCH_FAKE_OBJ
)

type LossFlags int

const (
	LossFlagNone       LossFlags = 0
	LossFlagResolution LossFlags = C.FF_LOSS_RESOLUTION
	LossFlagDepth      LossFlags = C.FF_LOSS_DEPTH
	LossFlagColorspace LossFlags = C.FF_LOSS_COLORSPACE
	LossFlagAlpha      LossFlags = C.FF_LOSS_ALPHA
	LossFlagColorquant LossFlags = C.FF_LOSS_COLORQUANT
	LossFlagChroma     LossFlags = C.FF_LOSS_CHROMA
	LossFlagAll        LossFlags = -1
)

func init() {
	SetLogLevel(LogLevelQuiet)
}

func Version() (int, int, int) {
	return int(C.LIBAVUTIL_VERSION_MAJOR), int(C.LIBAVUTIL_VERSION_MINOR), int(C.LIBAVUTIL_VERSION_MICRO)
}

func SetLogLevel(level LogLevel) {
	C.av_log_set_level(C.int(level))
}

type SampleFormat C.enum_AVSampleFormat

const (
	SampleFormatNone SampleFormat = C.AV_SAMPLE_FMT_NONE
	SampleFormatU8   SampleFormat = C.AV_SAMPLE_FMT_U8
	SampleFormatS16  SampleFormat = C.AV_SAMPLE_FMT_S16
	SampleFormatS32  SampleFormat = C.AV_SAMPLE_FMT_S32
	SampleFormatFLT  SampleFormat = C.AV_SAMPLE_FMT_FLT
	SampleFormatDBL  SampleFormat = C.AV_SAMPLE_FMT_DBL
	SampleFormatU8P  SampleFormat = C.AV_SAMPLE_FMT_U8P
	SampleFormatS16P SampleFormat = C.AV_SAMPLE_FMT_S16P
	SampleFormatS32P SampleFormat = C.AV_SAMPLE_FMT_S32P
	SampleFormatFLTP SampleFormat = C.AV_SAMPLE_FMT_FLTP
	SampleFormatDBLP SampleFormat = C.AV_SAMPLE_FMT_DBLP
)

func FindSampleFormatByName(name string) (SampleFormat, bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cSampleFormat := C.av_get_sample_fmt(cName)
	return SampleFormat(cSampleFormat), (cSampleFormat != C.AV_SAMPLE_FMT_NONE)
}

func (sfmt SampleFormat) Name() string {
	str, _ := sfmt.NameOk()
	return str
}

func (sfmt SampleFormat) NameOk() (string, bool) {
	return cStringToStringOk(C.av_get_sample_fmt_name((C.enum_AVSampleFormat)(sfmt)))
}

type PixelFormat C.enum_AVPixelFormat

const (
	PixelFormatNone PixelFormat = C.AV_PIX_FMT_NONE
)

func FindPixelFormatByName(name string) (PixelFormat, bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cPixelFormat := C.av_get_pix_fmt(cName)
	return PixelFormat(cPixelFormat), (cPixelFormat != C.AV_PIX_FMT_NONE)
}

func (pfmt PixelFormat) Name() string {
	str, _ := pfmt.NameOk()
	return str
}

func (pfmt PixelFormat) NameOk() (string, bool) {
	return cStringToStringOk(C.av_get_pix_fmt_name((C.enum_AVPixelFormat)(pfmt)))
}

type PixelFormatDescriptor struct {
	CAVPixFmtDescriptor *C.AVPixFmtDescriptor
}

func NewPixelFormatDescriptorFromC(cCtx unsafe.Pointer) *PixelFormatDescriptor {
	return &PixelFormatDescriptor{CAVPixFmtDescriptor: (*C.AVPixFmtDescriptor)(cCtx)}
}

func FindPixelFormatDescriptorByPixelFormat(pixelFormat PixelFormat) *PixelFormatDescriptor {
	cDescriptor := C.av_pix_fmt_desc_get(C.enum_AVPixelFormat(pixelFormat))
	if cDescriptor == nil {
		return nil
	}
	return NewPixelFormatDescriptorFromC(unsafe.Pointer(cDescriptor))
}

func (d *PixelFormatDescriptor) ComponentCount() int {
	return int(d.CAVPixFmtDescriptor.nb_components)
}

type ChannelLayout uint64

const (
	ChannelLayoutMono            ChannelLayout = C.AV_CH_LAYOUT_MONO
	ChannelLayoutStereo          ChannelLayout = C.AV_CH_LAYOUT_STEREO
	ChannelLayout2Point1         ChannelLayout = C.AV_CH_LAYOUT_2POINT1
	ChannelLayout21              ChannelLayout = C.AV_CH_LAYOUT_2_1
	ChannelLayoutSurround        ChannelLayout = C.AV_CH_LAYOUT_SURROUND
	ChannelLayout3Point1         ChannelLayout = C.AV_CH_LAYOUT_3POINT1
	ChannelLayout4Point0         ChannelLayout = C.AV_CH_LAYOUT_4POINT0
	ChannelLayout4Point1         ChannelLayout = C.AV_CH_LAYOUT_4POINT1
	ChannelLayout22              ChannelLayout = C.AV_CH_LAYOUT_2_2
	ChannelLayoutQuad            ChannelLayout = C.AV_CH_LAYOUT_QUAD
	ChannelLayout5Point0         ChannelLayout = C.AV_CH_LAYOUT_5POINT0
	ChannelLayout5Point1         ChannelLayout = C.AV_CH_LAYOUT_5POINT1
	ChannelLayout5Point0Back     ChannelLayout = C.AV_CH_LAYOUT_5POINT0_BACK
	ChannelLayout5Point1Back     ChannelLayout = C.AV_CH_LAYOUT_5POINT1_BACK
	ChannelLayout6Point0         ChannelLayout = C.AV_CH_LAYOUT_6POINT0
	ChannelLayout6Point0Front    ChannelLayout = C.AV_CH_LAYOUT_6POINT0_FRONT
	ChannelLayoutHexagonal       ChannelLayout = C.AV_CH_LAYOUT_HEXAGONAL
	ChannelLayout6Point1         ChannelLayout = C.AV_CH_LAYOUT_6POINT1
	ChannelLayout6Point1Back     ChannelLayout = C.AV_CH_LAYOUT_6POINT1_BACK
	ChannelLayout6Point1Front    ChannelLayout = C.AV_CH_LAYOUT_6POINT1_FRONT
	ChannelLayout7Point0         ChannelLayout = C.AV_CH_LAYOUT_7POINT0
	ChannelLayout7Point0Front    ChannelLayout = C.AV_CH_LAYOUT_7POINT0_FRONT
	ChannelLayout7Point1         ChannelLayout = C.AV_CH_LAYOUT_7POINT1
	ChannelLayout7Point1Wide     ChannelLayout = C.AV_CH_LAYOUT_7POINT1_WIDE
	ChannelLayout7Point1WideBack ChannelLayout = C.AV_CH_LAYOUT_7POINT1_WIDE_BACK
	ChannelLayoutOctagonal       ChannelLayout = C.AV_CH_LAYOUT_OCTAGONAL
	ChannelLayoutStereoDownmix   ChannelLayout = C.AV_CH_LAYOUT_STEREO_DOWNMIX
)

func FindChannelLayoutByName(name string) (ChannelLayout, bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cChannelLayout := C.av_get_channel_layout(cName)
	return ChannelLayout(cChannelLayout), (cChannelLayout != 0)
}

func FindDefaultChannelLayout(numberOfChannels int) (ChannelLayout, bool) {
	cl := C.av_get_default_channel_layout(C.int(numberOfChannels))
	if cl <= 0 {
		return 0, false
	}
	return ChannelLayout(cl), true
}

func (cl ChannelLayout) NumberOfChannels() int {
	return int(C.av_get_channel_layout_nb_channels((C.uint64_t)(cl)))
}

func (cl ChannelLayout) Name() string {
	str, _ := cl.NameOk()
	return str
}

func (cl ChannelLayout) NameOk() (string, bool) {
	for index := C.unsigned(0); ; index++ {
		var cCL C.uint64_t
		var cName *C.char
		if C.av_get_standard_channel_layout(index, &cCL, &cName) != 0 {
			break
		}
		if ChannelLayout(cCL) == cl {
			return cStringToStringOk(cName)
		}
	}
	return "", false
}

func (cl ChannelLayout) DescriptionOk() (string, bool) {
	return cStringToStringOk(C.av_get_channel_description((C.uint64_t)(cl)))
}

func ChannelLayouts() []ChannelLayout {
	var cls []ChannelLayout
	for index := C.unsigned(0); ; index++ {
		var cCL C.uint64_t
		if C.av_get_standard_channel_layout(index, &cCL, nil) != 0 {
			break
		}
		cls = append(cls, ChannelLayout(cCL))
	}
	return cls
}

type Rational struct {
	CAVRational C.AVRational
}

func NewRational(numerator, denominator int) *Rational {
	r := &Rational{}
	r.CAVRational.num = C.int(numerator)
	r.CAVRational.den = C.int(denominator)
	return r
}

var zeroRational = NewRational(0, 1)

func NewRationalFromC(cRational unsafe.Pointer) *Rational {
	rational := (*C.AVRational)(cRational)
	return NewRational(int(rational.num), int(rational.den))
}

func (r *Rational) String() string {
	return strconv.Itoa(r.Numerator()) + ":" + strconv.Itoa(r.Denominator())
}

func (r *Rational) Numerator() int {
	return int(r.CAVRational.num)
}

func (r *Rational) SetNumerator(numerator int) {
	r.CAVRational.num = (C.int)(numerator)
}

func (r *Rational) Denominator() int {
	return int(r.CAVRational.den)
}

func (r *Rational) SetDenominator(denominator int) {
	r.CAVRational.den = (C.int)(denominator)
}

func (r *Rational) Add(r2 *Rational) {
	r.CAVRational = C.av_add_q(r.CAVRational, r2.CAVRational)
}

func (r *Rational) Sub(r2 *Rational) {
	r.CAVRational = C.av_sub_q(r.CAVRational, r2.CAVRational)
}

func (r *Rational) Mul(r2 *Rational) {
	r.CAVRational = C.av_mul_q(r.CAVRational, r2.CAVRational)
}

func (r *Rational) Div(r2 *Rational) {
	r.CAVRational = C.av_div_q(r.CAVRational, r2.CAVRational)
}

func (r *Rational) Invert() {
	r.CAVRational = C.av_inv_q(r.CAVRational)
}

func (r *Rational) Reduce() {
	r.Add(zeroRational)
}

func (r *Rational) Compare(r2 *Rational) (int, bool) {
	res := C.av_cmp_q(r.CAVRational, r2.CAVRational)
	switch res {
	case 0, 1, -1:
		return int(res), true
	}
	return 0, false
}

func (r *Rational) Nearer(r2 *Rational, r3 *Rational) int {
	return int(C.av_nearer_q(r.CAVRational, r2.CAVRational, r3.CAVRational))
}

func (r *Rational) Nearest(rs []*Rational) *Rational {
	var nearest *Rational
	for _, r2 := range rs {
		if nearest == nil {
			nearest = r2
		} else {
			res := C.av_nearer_q(r.CAVRational, r2.CAVRational, nearest.CAVRational)
			if res > 0 {
				nearest = r2
			}
		}
	}
	return nearest
}

func (r *Rational) Copy() *Rational {
	r2 := &Rational{}
	r2.CAVRational.num = r.CAVRational.num
	r2.CAVRational.den = r.CAVRational.den
	return r2
}

func (r *Rational) Float64() float64 {
	return float64(r.CAVRational.num) / float64(r.CAVRational.den)
}

var StandardTimeBase = NewRational(1, C.AV_TIME_BASE)

type Time struct {
	Point int64
	Base  *Rational
}

func NewTime(point int64, base *Rational) *Time {
	return &Time{
		Point: point,
		Base:  base,
	}
}

func (t *Time) Valid() bool {
	return time.Duration(t.Base.Denominator()) > 0
}

func (t *Time) Duration() (time.Duration, bool) {
	if !t.Valid() {
		return 0, false
	}
	x := t.Point * int64(t.Base.Numerator())
	d := time.Duration(x) * time.Second / time.Duration(t.Base.Denominator())
	return d, true
}

func ErrnoErrorCode(e syscall.Errno) ErrorCode {
	return ErrorCode(C.go_av_errno_to_error(C.int(e)))
}

type Error struct {
	code ErrorCode
	err  error
}

func NewErrorFromCode(code ErrorCode) *Error {
	return &Error{
		code: code,
		err:  errors.New(strError(C.int(code))),
	}
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Error() string {
	return e.err.Error()
}

func strError(code C.int) string {
	size := C.size_t(256)
	buf := (*C.char)(C.av_mallocz(size))
	defer C.av_free(unsafe.Pointer(buf))
	if C.av_strerror(code, buf, size-1) == 0 {
		return C.GoString(buf)
	}
	return "Unknown error"
}

var _ error = (*Error)(nil)

type Dictionary struct {
	CAVDictionary  **C.AVDictionary
	pCAVDictionary *C.AVDictionary
}

func NewDictionary() *Dictionary {
	return NewDictionaryFromC(nil)
}

func NewDictionaryFromC(cDictionary unsafe.Pointer) *Dictionary {
	return &Dictionary{CAVDictionary: (**C.AVDictionary)(cDictionary)}
}

func (dict *Dictionary) Free() {
	C.av_dict_free(dict.pointer())
}

func (dict *Dictionary) Pointer() unsafe.Pointer {
	return unsafe.Pointer(dict.pointer())
}

func (dict *Dictionary) pointer() **C.AVDictionary {
	if dict.CAVDictionary != nil {
		return dict.CAVDictionary
	}
	return &dict.pCAVDictionary
}

func (dict *Dictionary) Value() unsafe.Pointer {
	return unsafe.Pointer(dict.value())
}

func (dict *Dictionary) value() *C.AVDictionary {
	if dict.CAVDictionary != nil {
		return *dict.CAVDictionary
	}
	return dict.pCAVDictionary
}

func (dict *Dictionary) has(key string, flags C.int) bool {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	has := C.go_av_dict_has(dict.value(), cKey, flags)
	if has == 0 {
		return false
	}
	return true
}

func (dict *Dictionary) Has(key string) bool {
	return dict.has(key, C.AV_DICT_MATCH_CASE)
}

func (dict *Dictionary) HasInsensitive(key string) bool {
	return dict.has(key, 0)
}

func (dict *Dictionary) get(key string, flags C.int) (string, bool) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	entry := C.av_dict_get(dict.value(), cKey, nil, flags)
	if entry == nil {
		return "", false
	}
	return C.GoString(entry.value), true
}

func (dict *Dictionary) Get(key string) string {
	str, _ := dict.GetOk(key)
	return str
}

func (dict *Dictionary) GetOk(key string) (string, bool) {
	return dict.get(key, C.AV_DICT_MATCH_CASE)
}

func (dict *Dictionary) GetInsensitive(key string) string {
	str, _ := dict.GetInsensitiveOk(key)
	return str
}

func (dict *Dictionary) GetInsensitiveOk(key string) (string, bool) {
	return dict.get(key, 0)
}

func (dict *Dictionary) set(key, value string, flags C.int) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	code := C.av_dict_set(dict.pointer(), cKey, cValue, flags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (dict *Dictionary) Set(key, value string) error {
	return dict.set(key, value, C.AV_DICT_MATCH_CASE)
}

func (dict *Dictionary) Delete(key string) error {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	code := C.av_dict_set(dict.pointer(), cKey, nil, 0)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (dict *Dictionary) SetInsensitive(key, value string) error {
	return dict.set(key, value, 0)
}

func (dict *Dictionary) Count() int {
	return int(C.av_dict_count(dict.value()))
}

func (dict *Dictionary) Keys() []string {
	count := dict.Count()
	if count <= 0 {
		return nil
	}
	keys := make([]string, 0, count)
	var entry *C.AVDictionaryEntry
	for {
		entry = C.go_av_dict_next(dict.value(), entry)
		if entry == nil {
			break
		}
		keys = append(keys, C.GoString(entry.key))
	}
	return keys
}

func (dict *Dictionary) Values() []string {
	count := dict.Count()
	if count <= 0 {
		return nil
	}
	values := make([]string, 0, count)
	var entry *C.AVDictionaryEntry
	for {
		entry = C.go_av_dict_next(dict.value(), entry)
		if entry == nil {
			break
		}
		values = append(values, C.GoString(entry.value))
	}
	return values
}

func (dict *Dictionary) Map() map[string]string {
	count := dict.Count()
	if count <= 0 {
		return nil
	}
	m := make(map[string]string, count)
	var entry *C.AVDictionaryEntry
	for {
		entry = C.go_av_dict_next(dict.value(), entry)
		if entry == nil {
			break
		}
		m[C.GoString(entry.key)] = C.GoString(entry.value)
	}
	return m
}

func (dict *Dictionary) Copy() *Dictionary {
	newDict := NewDictionary()
	C.av_dict_copy(newDict.pointer(), dict.value(), C.AV_DICT_MATCH_CASE)
	return newDict
}

func (dict *Dictionary) String(keyValSep, pairsSep byte) (string, error) {
	buf := (*C.char)(nil)
	defer C.av_freep(unsafe.Pointer(&buf))
	code := C.av_dict_get_string(dict.value(), &buf, C.char(keyValSep), C.char(pairsSep))
	if code < 0 {
		return "", NewErrorFromCode(ErrorCode(code))
	}
	return C.GoString(buf), nil
}

func cStringToStringOk(cStr *C.char) (string, bool) {
	if cStr == nil {
		return "", false
	}
	return C.GoString(cStr), true
}

type Option struct {
	CAVOption *C.struct_AVOption
}

func NewOptionFromC(cOption unsafe.Pointer) *Option {
	return &Option{CAVOption: (*C.struct_AVOption)(cOption)}
}

func (o *Option) Name() string {
	str, _ := o.NameOk()
	return str
}

func (o *Option) NameOk() (string, bool) {
	return cStringToStringOk(o.CAVOption.name)
}

func (o *Option) Help() string {
	str, _ := o.HelpOk()
	return str
}

func (o *Option) HelpOk() (string, bool) {
	return cStringToStringOk(o.CAVOption.help)
}

type Class struct {
	CAVClass *C.AVClass
}

func NewClassFromC(cClass unsafe.Pointer) *Class {
	return &Class{CAVClass: (*C.AVClass)(cClass)}
}

func (c *Class) Name() string {
	str, _ := c.NameOk()
	return str
}

func (c *Class) NameOk() (string, bool) {
	return cStringToStringOk(c.CAVClass.class_name)
}

func (c *Class) Options() []*Option {
	var cur *C.struct_AVOption
	var options []*Option
	for {
		cur = C.av_opt_next(unsafe.Pointer(&c.CAVClass), cur)
		if cur == nil {
			break
		}
		options = append(options, NewOptionFromC(unsafe.Pointer(cur)))
	}
	return options
}

func (c *Class) ChildrenClasses() []*Class {
	var child *C.AVClass
	var children []*Class
	for {
		child = C.av_opt_child_class_next(c.CAVClass, child)
		if child == nil {
			break
		}
		children = append(children, NewClassFromC(unsafe.Pointer(child)))
	}
	return children
}

type Frame struct {
	CAVFrame *C.AVFrame
}

func NewFrame() (*Frame, error) {
	cFrame := C.av_frame_alloc()
	if cFrame == nil {
		return nil, ErrAllocationError
	}
	return NewFrameFromC(unsafe.Pointer(cFrame)), nil
}

func NewFrameFromC(cFrame unsafe.Pointer) *Frame {
	return &Frame{CAVFrame: (*C.AVFrame)(cFrame)}
}

func (f *Frame) Free() {
	if f.CAVFrame != nil {
		C.av_frame_free(&f.CAVFrame)
	}
}

func (f *Frame) Ref(dst *Frame) error {
	code := C.av_frame_ref(dst.CAVFrame, f.CAVFrame)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (f *Frame) Unref() {
	C.av_frame_unref(f.CAVFrame)
}

func (f *Frame) GetBuffer() error {
	return f.GetBufferWithAlignment(32)
}

func (f *Frame) GetBufferWithAlignment(alignment int) error {
	code := C.av_frame_get_buffer(f.CAVFrame, C.int(alignment))
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (f *Frame) Data(index int) unsafe.Pointer {
	return unsafe.Pointer(f.CAVFrame.data[index])
}

func (f *Frame) SetData(index int, data unsafe.Pointer) {
	f.CAVFrame.data[index] = (*C.uint8_t)(data)
}

func (f *Frame) LineSize(index int) int {
	return int(f.CAVFrame.linesize[index])
}

func (f *Frame) SetLineSize(index int, lineSize int) {
	f.CAVFrame.linesize[index] = (C.int)(lineSize)
}

func (f *Frame) ExtendedData() unsafe.Pointer {
	return unsafe.Pointer(f.CAVFrame.extended_data)
}

func (f *Frame) SetExtendedData(data unsafe.Pointer) {
	f.CAVFrame.extended_data = (**C.uint8_t)(data)
}

func (f *Frame) Width() int {
	return int(f.CAVFrame.width)
}

func (f *Frame) SetWidth(width int) {
	f.CAVFrame.width = (C.int)(width)
}

func (f *Frame) Height() int {
	return int(f.CAVFrame.height)
}

func (f *Frame) SetHeight(height int) {
	f.CAVFrame.height = (C.int)(height)
}

func (f *Frame) NumberOfSamples() int {
	return int(f.CAVFrame.nb_samples)
}

func (f *Frame) SetNumberOfSamples(samples int) {
	f.CAVFrame.nb_samples = (C.int)(samples)
}

func (f *Frame) PixelFormat() PixelFormat {
	return PixelFormat(f.CAVFrame.format)
}

func (f *Frame) SetPixelFormat(format PixelFormat) {
	f.CAVFrame.format = (C.int)(format)
}

func (f *Frame) KeyFrame() bool {
	return f.CAVFrame.key_frame != (C.int)(0)
}

func (f *Frame) SetKeyFrame(keyFrame bool) {
	if keyFrame {
		f.CAVFrame.key_frame = 1
	} else {
		f.CAVFrame.key_frame = 0
	}
}

func (f *Frame) PictureType() PictureType {
	return PictureType(f.CAVFrame.pict_type)
}

func (f *Frame) SetPictureType(ptype PictureType) {
	f.CAVFrame.pict_type = (C.enum_AVPictureType)(ptype)
}

func (f *Frame) SampleAspectRatio() *Rational {
	return NewRationalFromC(unsafe.Pointer(&f.CAVFrame.sample_aspect_ratio))
}

func (f *Frame) SetSampleAspectRatio(aspectRatio *Rational) {
	f.CAVFrame.sample_aspect_ratio.num = (C.int)(aspectRatio.Numerator())
	f.CAVFrame.sample_aspect_ratio.den = (C.int)(aspectRatio.Denominator())
}

func (f *Frame) PTS() int64 {
	return int64(f.CAVFrame.pts)
}

func (f *Frame) SetPTS(pts int64) {
	f.CAVFrame.pts = (C.int64_t)(pts)
}

func (f *Frame) PacketPTS() int64 {
	return int64(f.CAVFrame.pkt_pts)
}

func (f *Frame) SetPacketPTS(pts int64) {
	f.CAVFrame.pkt_pts = (C.int64_t)(pts)
}

func (f *Frame) PacketDTS() int64 {
	return int64(f.CAVFrame.pkt_dts)
}

func (f *Frame) SetPacketDTS(dts int64) {
	f.CAVFrame.pkt_dts = (C.int64_t)(dts)
}

func (f *Frame) CodedPictureNumber() int {
	return int(f.CAVFrame.coded_picture_number)
}

func (f *Frame) SetCodedPictureNumber(number int) {
	f.CAVFrame.coded_picture_number = (C.int)(number)
}

func (f *Frame) DisplayPictureNumber() int {
	return int(f.CAVFrame.display_picture_number)
}

func (f *Frame) SetDisplayPictureNumber(number int) {
	f.CAVFrame.display_picture_number = (C.int)(number)
}

func (f *Frame) Quality() int {
	return int(f.CAVFrame.quality)
}

func (f *Frame) SetQuality(quality int) {
	f.CAVFrame.quality = (C.int)(quality)
}

func (f *Frame) Opaque() unsafe.Pointer {
	return unsafe.Pointer(f.CAVFrame.opaque)
}

func (f *Frame) SetOpaque(opaque unsafe.Pointer) {
	f.CAVFrame.opaque = opaque
}

func (f *Frame) Metadata() *Dictionary {
	dict := C.av_frame_get_metadata(f.CAVFrame)
	if dict == nil {
		return nil
	}
	return NewDictionaryFromC(unsafe.Pointer(&dict))
}

func (f *Frame) SetMetadata(dict *Dictionary) {
	if dict == nil {
		C.av_frame_set_metadata(f.CAVFrame, nil)
		return
	}
	C.av_frame_set_metadata(f.CAVFrame, dict.value())
}

func (f *Frame) BestEffortTimestamp() int64 {
	return int64(C.av_frame_get_best_effort_timestamp(f.CAVFrame))
}

func (f *Frame) PacketDuration() int64 {
	return int64(C.av_frame_get_pkt_duration(f.CAVFrame))
}

type OptionAccessor struct {
	obj  unsafe.Pointer
	fake bool
}

func NewOptionAccessor(obj unsafe.Pointer, fake bool) *OptionAccessor {
	return &OptionAccessor{obj: obj, fake: fake}
}

func (oa *OptionAccessor) GetOption(name string) (string, bool, error) {
	return oa.GetOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetOptionWithFlags(name string, flags OptionSearchFlags) (string, bool, error) {
	var cOut *C.uint8_t
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get(oa.obj, cName, searchFlags, &cOut)
	if code < 0 {
		return "", false, getOptionError(code)
	}
	defer C.av_free(unsafe.Pointer(cOut))
	cStr := (*C.char)(unsafe.Pointer(cOut))
	return C.GoString(cStr), true, nil
}

func (oa *OptionAccessor) GetInt64Option(name string) (int64, bool, error) {
	return oa.GetInt64OptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetInt64OptionWithFlags(name string, flags OptionSearchFlags) (int64, bool, error) {
	var cOut C.int64_t
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_int(oa.obj, cName, searchFlags, &cOut)
	if code < 0 {
		return 0, false, getOptionError(code)
	}
	return int64(cOut), true, nil
}

func (oa *OptionAccessor) GetFloat64Option(name string) (float64, bool, error) {
	return oa.GetFloat64OptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetFloat64OptionWithFlags(name string, flags OptionSearchFlags) (float64, bool, error) {
	var cOut C.double
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_double(oa.obj, cName, searchFlags, &cOut)
	if code < 0 {
		return 0, false, getOptionError(code)
	}
	return float64(cOut), true, nil
}

func (oa *OptionAccessor) GetRationalOption(name string) (*Rational, error) {
	return oa.GetRationalOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetRationalOptionWithFlags(name string, flags OptionSearchFlags) (*Rational, error) {
	cOut := &Rational{}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_q(oa.obj, cName, searchFlags, &cOut.CAVRational)
	if code < 0 {
		return nil, getOptionError(code)
	}
	return cOut, nil
}

func (oa *OptionAccessor) GetImageSizeOption(name string) (int, int, bool, error) {
	return oa.GetImageSizeOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetImageSizeOptionWithFlags(name string, flags OptionSearchFlags) (int, int, bool, error) {
	var cOut1, cOut2 C.int
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_image_size(oa.obj, cName, searchFlags, &cOut1, &cOut2)
	if code < 0 {
		return 0, 0, false, getOptionError(code)
	}
	return int(cOut1), int(cOut2), true, nil
}

func (oa *OptionAccessor) GetPixelFormatOption(name string) (PixelFormat, bool, error) {
	return oa.GetPixelFormatOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetPixelFormatOptionWithFlags(name string, flags OptionSearchFlags) (PixelFormat, bool, error) {
	var cOut C.enum_AVPixelFormat
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_pixel_fmt(oa.obj, cName, searchFlags, &cOut)
	if code < 0 {
		return 0, false, getOptionError(code)
	}
	return PixelFormat(cOut), true, nil
}

func (oa *OptionAccessor) GetSampleFormatOption(name string) (SampleFormat, bool, error) {
	return oa.GetSampleFormatOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetSampleFormatOptionWithFlags(name string, flags OptionSearchFlags) (SampleFormat, bool, error) {
	var cOut C.enum_AVSampleFormat
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_sample_fmt(oa.obj, cName, searchFlags, &cOut)
	if code < 0 {
		return 0, false, getOptionError(code)
	}
	return SampleFormat(cOut), true, nil
}

func (oa *OptionAccessor) GetVideoRateOption(name string) (*Rational, error) {
	return oa.GetVideoRateOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetVideoRateOptionWithFlags(name string, flags OptionSearchFlags) (*Rational, error) {
	cOut := &Rational{}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_video_rate(oa.obj, cName, searchFlags, &cOut.CAVRational)
	if code < 0 {
		return nil, getOptionError(code)
	}
	return cOut, nil
}

func (oa *OptionAccessor) GetChannelLayoutOption(name string) (int64, bool, error) {
	return oa.GetChannelLayoutOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetChannelLayoutOptionWithFlags(name string, flags OptionSearchFlags) (int64, bool, error) {
	var cOut C.int64_t
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_channel_layout(oa.obj, cName, searchFlags, &cOut)
	if code < 0 {
		return 0, false, getOptionError(code)
	}
	return int64(cOut), true, nil
}

func (oa *OptionAccessor) GetDictionaryOption(name string) (*Dictionary, error) {
	return oa.GetDictionaryOptionWithFlags(name, OptionSearchChildren)
}

func (oa *OptionAccessor) GetDictionaryOptionWithFlags(name string, flags OptionSearchFlags) (*Dictionary, error) {
	cOut := &Dictionary{}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_get_dict_val(oa.obj, cName, searchFlags, cOut.pointer())
	if code < 0 {
		return nil, getOptionError(code)
	}
	return cOut, nil
}

func (oa *OptionAccessor) SetOption(name, value string) error {
	return oa.SetOptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetOptionWithFlags(name, value string, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set(oa.obj, cName, cValue, searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetInt64Option(name string, value int64) error {
	return oa.SetInt64OptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetInt64OptionWithFlags(name string, value int64, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_int(oa.obj, cName, (C.int64_t)(value), searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetFloat64Option(name string, value float64) error {
	return oa.SetFloat64OptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetFloat64OptionWithFlags(name string, value float64, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_double(oa.obj, cName, (C.double)(value), searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetRationalOption(name string, value *Rational) error {
	return oa.SetRationalOptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetRationalOptionWithFlags(name string, value *Rational, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_q(oa.obj, cName, value.CAVRational, searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetBinaryOption(name string, value unsafe.Pointer, size int) error {
	return oa.SetBinaryOptionWithFlags(name, value, size, OptionSearchChildren)
}

func (oa *OptionAccessor) SetBinaryOptionWithFlags(name string, value unsafe.Pointer, size int, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_bin(oa.obj, cName, (*C.uint8_t)(value), (C.int)(size), searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetIntArrayOption(name string, value []int) error {
	return oa.SetIntArrayOptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetIntArrayOptionWithFlags(name string, value []int, flags OptionSearchFlags) error {
	var ptr unsafe.Pointer
	var value2 []C.int
	num := len(value)
	if num > 0 {
		if C.sizeof_int == unsafe.Sizeof(value[0]) {
			ptr = unsafe.Pointer(&value[0])
		} else {
			value2 = make([]C.int, num, num)
			for i := 0; i < num; i++ {
				value2[i] = C.int(value[i])
			}
			ptr = unsafe.Pointer(&value2[0])
		}
	}
	return oa.SetBinaryOptionWithFlags(name, ptr, num*C.sizeof_int, flags)
}

func (oa *OptionAccessor) SetImageSizeOption(name string, width, height int) error {
	return oa.SetImageSizeOptionWithFlags(name, width, height, OptionSearchChildren)
}

func (oa *OptionAccessor) SetImageSizeOptionWithFlags(name string, width, height int, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_image_size(oa.obj, cName, (C.int)(width), (C.int)(height), searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetPixelFormatOption(name string, value PixelFormat) error {
	return oa.SetPixelFormatOptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetPixelFormatOptionWithFlags(name string, value PixelFormat, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_pixel_fmt(oa.obj, cName, (C.enum_AVPixelFormat)(value), searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetVideoRateOption(name string, value *Rational) error {
	return oa.SetVideoRateOptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetVideoRateOptionWithFlags(name string, value *Rational, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_video_rate(oa.obj, cName, value.CAVRational, searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetChannelLayoutOption(name string, value int64) error {
	return oa.SetChannelLayoutOptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetChannelLayoutOptionWithFlags(name string, value int64, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_channel_layout(oa.obj, cName, (C.int64_t)(value), searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) SetDictionaryOption(name string, value *Dictionary) error {
	return oa.SetDictionaryOptionWithFlags(name, value, OptionSearchChildren)
}

func (oa *OptionAccessor) SetDictionaryOptionWithFlags(name string, value *Dictionary, flags OptionSearchFlags) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	searchFlags := oa.searchFlags(flags)
	code := C.av_opt_set_dict_val(oa.obj, cName, value.value(), searchFlags)
	if code < 0 {
		return NewErrorFromCode(ErrorCode(code))
	}
	return nil
}

func (oa *OptionAccessor) searchFlags(flags OptionSearchFlags) C.int {
	flags &^= OptionSearchFakeObj
	if oa.fake {
		flags |= OptionSearchFakeObj
	}
	return C.int(flags)
}

func getOptionError(code C.int) error {
	if ErrorCode(code) == ErrorCodeOptionNotFound {
		return nil
	}
	return NewErrorFromCode(ErrorCode(code))
}

type Expr struct {
	CAVExpr *C.struct_AVExpr
}

func NewExpr(value string, constNames []string) (*Expr, error) {
	e := NewExprFromC(nil)
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	cConstNames := make([]*C.char, C.int(len(constNames)+1))
	for i, constName := range constNames {
		cConstNames[i] = C.CString(constName)
		defer C.free(unsafe.Pointer(cConstNames[i]))
	}
	code := C.go_av_expr_parse2(&e.CAVExpr, cValue, (**C.char)(&cConstNames[0]), 0, nil)
	if code < 0 {
		return nil, NewErrorFromCode(ErrorCode(code))
	}
	return e, nil
}

func NewExprFromC(cExpr unsafe.Pointer) *Expr {
	return &Expr{CAVExpr: (*C.struct_AVExpr)(cExpr)}
}

func (e *Expr) Evaluate(constValues []float64) (float64, error) {
	if len(constValues) == 0 {
		return 0, ErrInvalidArgumentSize
	}
	var cRet C.double
	cRet = C.av_expr_eval(e.CAVExpr, (*C.double)(&constValues[0]), nil)
	return float64(cRet), nil
}

func (e *Expr) Free() {
	if e.CAVExpr != nil {
		defer C.av_expr_free(e.CAVExpr)
		e.CAVExpr = nil
	}
}

func String(str string) *string {
	return &str
}

func Rescale(a, b, c int64) int64 {
	return int64(C.av_rescale(C.int64_t(a), C.int64_t(b), C.int64_t(c)))
}

func RescaleByRationals(a int64, bq, cq *Rational) int64 {
	return int64(C.av_rescale_q(C.int64_t(a), bq.CAVRational, cq.CAVRational))
}

func ParseRational(ratio string, max int) (*Rational, error) {
	cRatio := C.CString(ratio)
	defer C.free(unsafe.Pointer(cRatio))
	var cRational C.AVRational
	code := C.av_parse_ratio(&cRational, cRatio, C.int(max), C.int(0), nil)
	if code < 0 {
		return nil, NewErrorFromCode(ErrorCode(code))
	}
	return NewRationalFromC(unsafe.Pointer(&cRational)), nil
}

func ParseTime(timestr string, duration bool) (int64, error) {
	cTimestr := C.CString(timestr)
	defer C.free(unsafe.Pointer(cTimestr))
	x := C.int64_t(0)
	code := C.av_parse_time(&x, cTimestr, boolToCInt(duration))
	if code < 0 {
		return 0, NewErrorFromCode(ErrorCode(code))
	}
	return int64(x), nil
}

func Clip(x, min, max int) int {
	return int(C.av_clip(C.int(x), C.int(min), C.int(max)))
}

func boolToCInt(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
