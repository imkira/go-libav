// +build ffmpeg33

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
	"strings"
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

func (ctx *Context) Decode(pkt *Packet) (bool, []*avutil.Frame, error) {
	frames := []*avutil.Frame{}

	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
	C.avcodec_send_packet(ctx.CAVCodecContext, cPkt)
	var err error
	var code C.int
	for {
		frame, err := avutil.NewFrame()
		if err != nil {
			return false, nil, err
		}
		cFrame := (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
		code = C.avcodec_receive_frame(ctx.CAVCodecContext, cFrame)
		if code < 0 {
			err = avutil.NewErrorFromCode(avutil.ErrorCode(code))
			if code == C.GO_AVERROR(C.EAGAIN) {
				// We have all the output we are going to get
				break
			}
			code = 0
			return false, nil, err
		}
		frames = append(frames, frame)
	}
	return code == C.GO_AVERROR(C.EAGAIN), frames, err
}

func (ctx *Context) Encode(pkt *Packet, frames []*avutil.Frame) (int, error) {
	var cGotPkt C.int
	var cFrame *C.AVFrame
	var code C.int
	var frame *avutil.Frame
	var count int
	if frames != nil {
		for {
			if len(frames) == 0 {
				break
			}
			frame, frames = frames[0], frames[1:]
			cFrame = (*C.AVFrame)(unsafe.Pointer(frame.CAVFrame))
			code = C.avcodec_send_frame(ctx.CAVCodecContext, cFrame)
			var err error
			if code < 0 {
				err = avutil.NewErrorFromCode(avutil.ErrorCode(code))
				if !strings.HasPrefix(err.Error(), "Resource temporarily unavailable") {
					return count, err
				} else {
					// Stop calling this as we need to drain the buffer
					break
				}
			}
			count += 1
		}
	}
	cPkt := (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))

	cGotPkt = C.avcodec_receive_packet(ctx.CAVCodecContext, cPkt)
	if cGotPkt == (C.int)(0) {
		return count, ErrGotNoPacket
	}
	return count, nil
}
