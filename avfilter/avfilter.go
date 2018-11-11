package avfilter

//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavutil/error.h>
//#include <libavfilter/avfilter.h>
//#include <libavfilter/buffersrc.h>
//#include <libavfilter/buffersink.h>
//
//#ifdef AV_BUFFERSRC_FLAG_NO_COPY
//#define GO_AV_BUFFERSRC_FLAG_NO_COPY AV_BUFFERSRC_FLAG_NO_COPY
//#else
//#define GO_AV_BUFFERSRC_FLAG_NO_COPY 0
//#endif
//
//static const AVFilterLink *go_av_links_get(const AVFilterLink **links, unsigned int n)
//{
//  return links[n];
//}
//
//static const int GO_AVERROR(int e)
//{
//  return AVERROR(e);
//}
//
// int GO_AVFILTER_VERSION_MAJOR = LIBAVFILTER_VERSION_MAJOR;
// int GO_AVFILTER_VERSION_MINOR = LIBAVFILTER_VERSION_MINOR;
// int GO_AVFILTER_VERSION_MICRO = LIBAVFILTER_VERSION_MICRO;
//
//#define GO_AVFILTER_AUTO_CONVERT_ALL ((unsigned)AVFILTER_AUTO_CONVERT_ALL)
//#define GO_AVFILTER_AUTO_CONVERT_NONE ((unsigned)AVFILTER_AUTO_CONVERT_NONE)
//
//#cgo pkg-config: libavfilter libavutil
import "C"

import (
	"errors"
	"unsafe"

	"github.com/baohavan/go-libav/avutil"
)

var (
	ErrAllocationError = errors.New("allocation error")
)

type Flags int

const (
	FlagDynamicInputs           Flags = C.AVFILTER_FLAG_DYNAMIC_INPUTS
	FlagDynamicOutputs          Flags = C.AVFILTER_FLAG_DYNAMIC_OUTPUTS
	FlagSliceThreads            Flags = C.AVFILTER_FLAG_SLICE_THREADS
	FlagSupportTimelineGeneric  Flags = C.AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC
	FlagSupportTimelineInternal Flags = C.AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL
	FlagSupportTimeline         Flags = C.AVFILTER_FLAG_SUPPORT_TIMELINE
)

type BufferSrcFlags C.int

const (
	BufferSrcFlagNoCheckFormat BufferSrcFlags = C.AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT
	BufferSrcFlagNoCopy        BufferSrcFlags = C.GO_AV_BUFFERSRC_FLAG_NO_COPY
	BufferSrcFlagPush          BufferSrcFlags = C.AV_BUFFERSRC_FLAG_PUSH
	BufferSrcFlagKeepRef       BufferSrcFlags = C.AV_BUFFERSRC_FLAG_KEEP_REF
)

type GraphAutoConvertFlags uint

const (
	GraphAutoConvertFlagAll  GraphAutoConvertFlags = C.GO_AVFILTER_AUTO_CONVERT_ALL
	GraphAutoConvertFlagNone GraphAutoConvertFlags = C.GO_AVFILTER_AUTO_CONVERT_NONE
)

func init() {
	RegisterAll()
}

func Version() (int, int, int) {
	return int(C.GO_AVFILTER_VERSION_MAJOR), int(C.GO_AVFILTER_VERSION_MINOR), int(C.GO_AVFILTER_VERSION_MICRO)
}

func RegisterAll() {
	C.avfilter_register_all()
}

type Filter struct {
	CAVFilter *C.AVFilter
}

func NewFilterFromC(cFilter unsafe.Pointer) *Filter {
	return &Filter{CAVFilter: (*C.AVFilter)(cFilter)}
}

func (f *Filter) Name() string {
	str, _ := f.NameOk()
	return str
}

func (f *Filter) NameOk() (string, bool) {
	return cStringToStringOk(f.CAVFilter.name)
}

func (f *Filter) Description() string {
	str, _ := f.DescriptionOk()
	return str
}

func (f *Filter) DescriptionOk() (string, bool) {
	return cStringToStringOk(f.CAVFilter.description)
}

func (f *Filter) PrivateClass() *avutil.Class {
	if f.CAVFilter.priv_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(f.CAVFilter.priv_class))
}

func (f *Filter) Flags() Flags {
	return Flags(f.CAVFilter.flags)
}

func Filters() []*Filter {
	var filters []*Filter
	var cPrev *C.AVFilter
	for {
		if cPrev = C.avfilter_next(cPrev); cPrev == nil {
			break
		}
		filters = append(filters, NewFilterFromC(unsafe.Pointer(cPrev)))
	}
	return filters
}

func FindFilterByName(name string) *Filter {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cFilter := C.avfilter_get_by_name(cName)
	if cFilter == nil {
		return nil
	}
	return NewFilterFromC(unsafe.Pointer(cFilter))
}

type Link struct {
	CAVFilterLink *C.AVFilterLink
}

func NewLinkFromC(cLink unsafe.Pointer) *Link {
	return &Link{CAVFilterLink: (*C.AVFilterLink)(cLink)}
}

func (l *Link) Src() *Context {
	cContext := l.CAVFilterLink.src
	if cContext == nil {
		return nil
	}
	return NewContextFromC(unsafe.Pointer(cContext))
}

func (l *Link) Dst() *Context {
	cContext := l.CAVFilterLink.dst
	if cContext == nil {
		return nil
	}
	return NewContextFromC(unsafe.Pointer(cContext))
}

func (l *Link) Type() avutil.MediaType {
	return (avutil.MediaType)(l.CAVFilterLink._type)
}

func (l *Link) Width() int {
	return int(l.CAVFilterLink.w)
}

func (l *Link) Height() int {
	return int(l.CAVFilterLink.h)
}

func (l *Link) SampleAspectRatio() *avutil.Rational {
	return avutil.NewRationalFromC(unsafe.Pointer(&l.CAVFilterLink.sample_aspect_ratio))
}

func (l *Link) ChannelLayout() avutil.ChannelLayout {
	return (avutil.ChannelLayout)(l.CAVFilterLink.channel_layout)
}

func (l *Link) SampleRate() int {
	return int(l.CAVFilterLink.sample_rate)
}

func (l *Link) Format() int {
	return int(l.CAVFilterLink.format)
}

func (l *Link) PixelFormat() avutil.PixelFormat {
	return avutil.PixelFormat(l.CAVFilterLink.format)
}

func (l *Link) SampleFormat() avutil.SampleFormat {
	return avutil.SampleFormat(l.CAVFilterLink.format)
}

func (l *Link) TimeBase() *avutil.Rational {
	return avutil.NewRationalFromC(unsafe.Pointer(&l.CAVFilterLink.time_base))
}

func (l *Link) RequestSamples() int {
	return int(l.CAVFilterLink.request_samples)
}

func (l *Link) CurrentPTS() int64 {
	return int64(l.CAVFilterLink.current_pts)
}

func (l *Link) AgeIndex() int {
	return int(l.CAVFilterLink.age_index)
}

func (l *Link) FrameRate() *avutil.Rational {
	return avutil.NewRationalFromC(unsafe.Pointer(&l.CAVFilterLink.frame_rate))
}

func (l *Link) MinSamples() int {
	return int(l.CAVFilterLink.min_samples)
}

func (l *Link) MaxSamples() int {
	return int(l.CAVFilterLink.max_samples)
}

func (l *Link) Channels() int {
	return int(C.avfilter_link_get_channels(l.CAVFilterLink))
}

type Context struct {
	CAVFilterContext *C.AVFilterContext
	*avutil.OptionAccessor
}

func NewContextFromC(cCtx unsafe.Pointer) *Context {
	return &Context{
		CAVFilterContext: (*C.AVFilterContext)(cCtx),
		OptionAccessor:   avutil.NewOptionAccessor(cCtx, false),
	}
}

func (ctx *Context) Init() error {
	options := avutil.NewDictionary()
	defer options.Free()
	return ctx.InitWithDictionary(options)
}

func (ctx *Context) InitWithString(args string) error {
	cArgs := C.CString(args)
	defer C.free(unsafe.Pointer(cArgs))
	code := C.avfilter_init_str(ctx.CAVFilterContext, cArgs)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) InitWithDictionary(options *avutil.Dictionary) error {
	var cOptions **C.AVDictionary
	if options != nil {
		cOptions = (**C.AVDictionary)(options.Pointer())
	}
	code := C.avfilter_init_dict(ctx.CAVFilterContext, cOptions)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) Link(srcPad uint, dst *Context, dstPad uint) error {
	cSrc := ctx.CAVFilterContext
	cDst := dst.CAVFilterContext
	code := C.avfilter_link(cSrc, C.uint(srcPad), cDst, C.uint(dstPad))
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) AddFrame(frame *avutil.Frame) error {
	var cFrame *C.AVFrame
	if frame != nil {
		cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	}
	code := C.av_buffersrc_add_frame(ctx.CAVFilterContext, cFrame)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) AddFrameWithFlags(frame *avutil.Frame, flags BufferSrcFlags) error {
	var cFrame *C.AVFrame
	if frame != nil {
		cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	}
	code := C.av_buffersrc_add_frame_flags(ctx.CAVFilterContext, cFrame, (C.int)(flags))
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) WriteFrame(frame *avutil.Frame) error {
	var cFrame *C.AVFrame
	if frame != nil {
		cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	}
	code := C.av_buffersrc_write_frame(ctx.CAVFilterContext, cFrame)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) GetFrame(frame *avutil.Frame) (bool, error) {
	var cFrame *C.AVFrame
	if frame != nil {
		cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	}
	code := C.av_buffersink_get_frame(ctx.CAVFilterContext, cFrame)
	if code < 0 {
		switch avutil.ErrorCode(code) {
		case avutil.ErrorCode(C.GO_AVERROR(C.EAGAIN)):
			return false, nil
		case avutil.ErrorCodeEOF:
			return false, nil
		}
		return false, avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return true, nil
}

func (ctx *Context) SetFrameSize(size uint) {
	C.av_buffersink_set_frame_size(ctx.CAVFilterContext, C.uint(size))
}

func (ctx *Context) Class() *avutil.Class {
	if ctx.CAVFilterContext.av_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(ctx.CAVFilterContext.av_class))
}

func (ctx *Context) Name() string {
	str, _ := ctx.NameOk()
	return str
}

func (ctx *Context) NameOk() (string, bool) {
	return cStringToStringOk(ctx.CAVFilterContext.name)
}

func (ctx *Context) Filter() *Filter {
	cFilter := ctx.CAVFilterContext.filter
	if cFilter == nil {
		return nil
	}
	return NewFilterFromC(unsafe.Pointer(cFilter))
}

func (ctx *Context) Inputs() []*Link {
	count := ctx.NumberOfInputs()
	if count <= 0 {
		return nil
	}
	links := make([]*Link, 0, count)
	for i := uint(0); i < count; i++ {
		cLink := C.go_av_links_get(ctx.CAVFilterContext.inputs, C.uint(i))
		link := NewLinkFromC(unsafe.Pointer(cLink))
		links = append(links, link)
	}
	return links
}

func (ctx *Context) NumberOfInputs() uint {
	return uint(ctx.CAVFilterContext.nb_inputs)
}

func (ctx *Context) Outputs() []*Link {
	count := ctx.NumberOfOutputs()
	if count <= 0 {
		return nil
	}
	links := make([]*Link, 0, count)
	for i := uint(0); i < count; i++ {
		cLink := C.go_av_links_get(ctx.CAVFilterContext.outputs, C.uint(i))
		link := NewLinkFromC(unsafe.Pointer(cLink))
		links = append(links, link)
	}
	return links
}

func (ctx *Context) NumberOfOutputs() uint {
	return uint(ctx.CAVFilterContext.nb_outputs)
}

func (ctx *Context) FrameRate() *avutil.Rational {
	r := C.av_buffersink_get_frame_rate(ctx.CAVFilterContext)
	return avutil.NewRationalFromC(unsafe.Pointer(&r))
}

type Graph struct {
	CAVFilterGraph *C.AVFilterGraph
}

func NewGraph() (*Graph, error) {
	cGraph := C.avfilter_graph_alloc()
	if cGraph == nil {
		return nil, ErrAllocationError
	}
	return NewGraphFromC(unsafe.Pointer(cGraph)), nil
}

func NewGraphFromC(cGraph unsafe.Pointer) *Graph {
	return &Graph{CAVFilterGraph: (*C.AVFilterGraph)(cGraph)}
}

func (g *Graph) Free() {
	if g.CAVFilterGraph != nil {
		C.avfilter_graph_free(&g.CAVFilterGraph)
	}
}

func (g *Graph) Class() *avutil.Class {
	if g.CAVFilterGraph.av_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(g.CAVFilterGraph.av_class))
}

func (g *Graph) NumberOfFilters() uint {
	return uint(g.CAVFilterGraph.nb_filters)
}

func (g *Graph) AddFilter(filter *Filter, name string) (*Context, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cCtx := C.avfilter_graph_alloc_filter(g.CAVFilterGraph, filter.CAVFilter, cName)
	if cCtx == nil {
		return nil, ErrAllocationError
	}
	return NewContextFromC(unsafe.Pointer(cCtx)), nil
}

func (g *Graph) Parse(filters string, input, output *InOut) error {
	cFilters := C.CString(filters)
	defer C.free(unsafe.Pointer(cFilters))
	cInput := &input.CAVFilterInOut
	cOutput := &output.CAVFilterInOut
	code := C.avfilter_graph_parse_ptr(g.CAVFilterGraph, cFilters, cInput, cOutput, nil)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (g *Graph) Config() error {
	code := C.avfilter_graph_config(g.CAVFilterGraph, nil)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (g *Graph) RequestOldest() error {
	code := C.avfilter_graph_request_oldest(g.CAVFilterGraph)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (g *Graph) Dump() (string, error) {
	cStr := C.avfilter_graph_dump(g.CAVFilterGraph, nil)
	if cStr == nil {
		return "", ErrAllocationError
	}
	defer C.av_free(unsafe.Pointer(cStr))
	return C.GoString(cStr), nil
}

func (g *Graph) DumpWithOptions(options string) (string, error) {
	cOptions := C.CString(options)
	defer C.free(unsafe.Pointer(cOptions))
	cStr := C.avfilter_graph_dump(g.CAVFilterGraph, cOptions)
	if cStr == nil {
		return "", ErrAllocationError
	}
	defer C.av_free(unsafe.Pointer(cStr))
	return C.GoString(cStr), nil
}

func (g *Graph) AutoConvertFlags() GraphAutoConvertFlags {
	return GraphAutoConvertFlags(g.CAVFilterGraph.disable_auto_convert)
}

func (g *Graph) SetAutoConvertFlags(flags GraphAutoConvertFlags) {
	C.avfilter_graph_set_auto_convert(g.CAVFilterGraph, (C.uint)(flags))
}

type InOut struct {
	CAVFilterInOut *C.AVFilterInOut
}

func NewInOut() (*InOut, error) {
	cInOut := C.avfilter_inout_alloc()
	if cInOut == nil {
		return nil, ErrAllocationError
	}
	return NewInOutFromC(unsafe.Pointer(cInOut)), nil
}

func NewInOutFromC(cInOut unsafe.Pointer) *InOut {
	return &InOut{CAVFilterInOut: (*C.AVFilterInOut)(cInOut)}
}

func (io *InOut) Free() {
	if io.CAVFilterInOut != nil {
		C.avfilter_inout_free(&io.CAVFilterInOut)
	}
}

func (io *InOut) Name() string {
	str, _ := io.NameOk()
	return str
}

func (io *InOut) NameOk() (string, bool) {
	return cStringToStringOk(io.CAVFilterInOut.name)
}

func (io *InOut) SetName(name string) error {
	C.free(unsafe.Pointer(io.CAVFilterInOut.name))
	io.CAVFilterInOut.name = C.CString(name)
	if io.CAVFilterInOut.name == nil {
		return ErrAllocationError
	}
	return nil
}

func (io *InOut) PadIndex() int {
	return int(io.CAVFilterInOut.pad_idx)
}

func (io *InOut) SetPadIndex(index int) {
	io.CAVFilterInOut.pad_idx = (C.int)(index)
}

func (io *InOut) Context() *Context {
	if io.CAVFilterInOut.filter_ctx == nil {
		return nil
	}
	return NewContextFromC(unsafe.Pointer(io.CAVFilterInOut.filter_ctx))
}

func (io *InOut) SetContext(ctx *Context) {
	var cCtx *C.AVFilterContext
	if ctx != nil {
		cCtx = ctx.CAVFilterContext
	}
	io.CAVFilterInOut.filter_ctx = cCtx
}

func (io *InOut) Next() *InOut {
	if io.CAVFilterInOut.next == nil {
		return nil
	}
	return NewInOutFromC(unsafe.Pointer(io.CAVFilterInOut.next))
}

func (io *InOut) SetNext(next *InOut) {
	var cInOut *C.AVFilterInOut
	if next != nil {
		cInOut = next.CAVFilterInOut
	}
	io.CAVFilterInOut.next = cInOut
}

func cStringToStringOk(cStr *C.char) (string, bool) {
	if cStr == nil {
		return "", false
	}
	return C.GoString(cStr), true
}
