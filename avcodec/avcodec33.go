package avcodec

//#include <libavutil/avutil.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/error.h>
//static const int GO_AVERROR(int e)
//{
//  return AVERROR(e);
//}
// #cgo LDFLAGS: -lavcodec -lavutil
import "C"

import (
	"errors"
	"unsafe"

	"github.com/SpalkLtd/go-libav/avutil"
)

var (
	ErrGotNoFrame  = errors.New("avcodec: GotFrame == 0, this means we got no frame from the encoder")
	ErrGotNoPacket = errors.New("avcodec: GotPacket == 0, this means we got no packet from the encoder")
)

type CodecParameters struct {
	CAVCodecParameters *C.AVCodecParameters
}

func NewCodecParameters() (*CodecParameters, error) {
	cPkt := (*C.AVCodecParameters)(C.avcodec_parameters_alloc())
	if cPkt == nil {
		return nil, ErrAllocationError
	}
	return NewCodecParametersFromC(unsafe.Pointer(cPkt)), nil
}

func NewCodecParametersFromC(cPSD unsafe.Pointer) *CodecParameters {
	return &CodecParameters{CAVCodecParameters: (*C.AVCodecParameters)(cPSD)}
}

func (cParams *CodecParameters) Free() {
	C.avcodec_parameters_free(&cParams.CAVCodecParameters)
}

func (ctx *Context) CopyTo(dst *Context) error {
	// added in lavc 57.33.100
	parameters, err := NewCodecParameters()
	if err != nil {
		return err
	}
	defer parameters.Free()
	cParams := (*C.AVCodecParameters)(unsafe.Pointer(parameters.CAVCodecParameters))
	code := C.avcodec_parameters_from_context(cParams, ctx.CAVCodecContext)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	code = C.avcodec_parameters_to_context(dst.CAVCodecContext, cParams)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) SendPacket(pkt *Packet) error {

	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	code := C.avcodec_send_packet(ctx.CAVCodecContext, cPkt)
	if code < 0 {
		err := avutil.NewErrorFromCode(avutil.ErrorCode(code))
		return err
	}
	return nil
}

func (ctx *Context) ReceiveFrame(frame *avutil.Frame) error {
	cFrame := (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	code := C.avcodec_receive_frame(ctx.CAVCodecContext, cFrame)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) SendFrame(frame *avutil.Frame) error {
	cFrame := (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
	code := C.avcodec_send_frame(ctx.CAVCodecContext, cFrame)
	if code < 0 {
		err := avutil.NewErrorFromCode(avutil.ErrorCode(code))
		return err
	}
	return nil
}

func (ctx *Context) ReceivePacket(pkt *Packet) error {
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	code := C.avcodec_receive_packet(ctx.CAVCodecContext, cPkt)
	if code < 0 {
		err := avutil.NewErrorFromCode(avutil.ErrorCode(code))
		return err
	}
	return nil
}
