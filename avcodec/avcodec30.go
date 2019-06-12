// +build ffmpeg30

package avcodec

//#include <libavutil/avutil.h>
//#include <libavcodec/avcodec.h>
//
// #cgo pkg-config: libavcodec libavutil
import "C"

import (
	"unsafe"

	"github.com/baohavan/go-libav/avutil"
)

func (ctx *Context) CopyTo(dst *Context) error {
	code := C.avcodec_copy_context(dst.CAVCodecContext, ctx.CAVCodecContext)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) DecodeVideo(pkt *Packet, frame *avutil.Frame) (bool, int, error) {
	var cGotFrame C.int
	cFrame := (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	code := C.avcodec_decode_video2(ctx.CAVCodecContext, cFrame, &cGotFrame, cPkt)
	var err error
	if code < 0 {
		err = avutil.NewErrorFromCode(avutil.ErrorCode(code))
		code = 0
	}
	return (cGotFrame != (C.int)(0)), int(code), err
}

func (ctx *Context) DecodeAudio(pkt *Packet, frame *avutil.Frame) (bool, int, error) {
	var cGotFrame C.int
	cFrame := (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	code := C.avcodec_decode_audio4(ctx.CAVCodecContext, cFrame, &cGotFrame, cPkt)
	var err error
	if code < 0 {
		err = avutil.NewErrorFromCode(avutil.ErrorCode(code))
		code = 0
	}
	return (cGotFrame != (C.int)(0)), int(code), err
}

func (ctx *Context) EncodeVideo(pkt *Packet, frame *avutil.Frame) (bool, error) {
	var cGotFrame C.int
	var cFrame *C.AVFrame
	if frame != nil {
		cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	}
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	code := C.avcodec_encode_video2(ctx.CAVCodecContext, cPkt, cFrame, &cGotFrame)
	var err error
	if code < 0 {
		err = avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return (cGotFrame != (C.int)(0)), err
}

func (ctx *Context) EncodeAudio(pkt *Packet, frame *avutil.Frame) (bool, error) {
	var cGotFrame C.int
	var cFrame *C.AVFrame
	if frame != nil {
		cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	}
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	code := C.avcodec_encode_audio2(ctx.CAVCodecContext, cPkt, cFrame, &cGotFrame)
	var err error
	if code < 0 {
		err = avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return (cGotFrame != (C.int)(0)), err
}

func (pkt *Packet) SplitSideData() error {
	code := C.av_packet_split_side_data(pkt.CAVPacket)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

type BitStreamFilterContext struct {
	CAVBitStreamFilterContext *C.AVBitStreamFilterContext
}

func NewBitStreamFilterContextFromC(cCtx unsafe.Pointer) *BitStreamFilterContext {
	return &BitStreamFilterContext{CAVBitStreamFilterContext: (*C.AVBitStreamFilterContext)(cCtx)}
}

func NewBitStreamFilterContextFromName(name string) (*BitStreamFilterContext, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cCtx := C.av_bitstream_filter_init(cName)
	if cCtx == nil {
		return nil, ErrBitStreamFilterNotFound
	}
	return NewBitStreamFilterContextFromC(unsafe.Pointer(cCtx)), nil
}

func (ctx *BitStreamFilterContext) Close() {
	if ctx.CAVBitStreamFilterContext != nil {
		C.av_bitstream_filter_close(ctx.CAVBitStreamFilterContext)
		ctx.CAVBitStreamFilterContext = nil
	}
}

func (ctx *BitStreamFilterContext) Next() *BitStreamFilterContext {
	next := ctx.CAVBitStreamFilterContext.next
	if next == nil {
		return nil
	}
	return NewBitStreamFilterContextFromC(unsafe.Pointer(next))
}

func (ctx *BitStreamFilterContext) SetNext(next *BitStreamFilterContext) {
	ctx.CAVBitStreamFilterContext.next = next.CAVBitStreamFilterContext
}

func (ctx *BitStreamFilterContext) Args() string {
	args, _ := ctx.ArgsOK()
	return args
}

func (ctx *BitStreamFilterContext) ArgsOK() (string, bool) {
	return cStringToStringOk(ctx.CAVBitStreamFilterContext.args)
}

func (ctx *BitStreamFilterContext) SetArgs(args *string) error {
	C.av_freep(unsafe.Pointer(&ctx.CAVBitStreamFilterContext.args))
	if args == nil {
		return nil
	}
	bArgs := []byte(*args)
	length := len(bArgs)
	cArgs := (*C.char)(C.av_malloc(C.size_t(length + 1)))
	if cArgs == nil {
		return ErrAllocationError
	}
	if length > 0 {
		C.memcpy(unsafe.Pointer(cArgs), unsafe.Pointer(&bArgs[0]), C.size_t(length))
	}
	C.memset(unsafe.Pointer(uintptr(unsafe.Pointer(cArgs))+uintptr(length)), 0, 1)
	ctx.CAVBitStreamFilterContext.args = cArgs
	return nil
}
