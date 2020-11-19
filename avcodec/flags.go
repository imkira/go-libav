// +build !ffmpeg43

package avcodec

//#include <libavutil/avutil.h>
//#include <libavcodec/avcodec.h>
//
//
//#ifdef CODEC_CAP_HWACCEL
//#define GO_CODEC_CAP_HWACCEL CODEC_CAP_HWACCEL
//#else
//#define GO_CODEC_CAP_HWACCEL 0
//#endif
//
//#ifdef CODEC_CAP_HWACCEL_VDPAU
//#define GO_CODEC_CAP_HWACCEL_VDPAU CODEC_CAP_HWACCEL_VDPAU
//#else
//#define GO_CODEC_CAP_HWACCEL_VDPAU 0
//#endif
//
// #cgo LDFLAGS: -lavcodec -lavutil
import "C"

type Flags int

const (
	FlagUnaligned     Flags = C.CODEC_FLAG_UNALIGNED
	FlagQScale        Flags = C.CODEC_FLAG_QSCALE
	Flag4MV           Flags = C.CODEC_FLAG_4MV
	FlagOutputCorrupt Flags = C.CODEC_FLAG_OUTPUT_CORRUPT
	FlagQPEL          Flags = C.CODEC_FLAG_QPEL
	FlagPass1         Flags = C.CODEC_FLAG_PASS1
	FlagPass2         Flags = C.CODEC_FLAG_PASS2
	FlagGray          Flags = C.CODEC_FLAG_GRAY
	FlagPSNR          Flags = C.CODEC_FLAG_PSNR
	FlagTruncated     Flags = C.CODEC_FLAG_TRUNCATED
	FlagInterlacedDCT Flags = C.CODEC_FLAG_INTERLACED_DCT
	FlagLowDelay      Flags = C.CODEC_FLAG_LOW_DELAY
	FlagGlobalHeader  Flags = C.CODEC_FLAG_GLOBAL_HEADER
	FlagBitExact      Flags = C.CODEC_FLAG_BITEXACT
	FlagACPred        Flags = C.CODEC_FLAG_AC_PRED
	FlagLoopFilter    Flags = C.CODEC_FLAG_LOOP_FILTER
	FlagInterlacedME  Flags = C.CODEC_FLAG_INTERLACED_ME
	FlagClosedGOP     Flags = C.CODEC_FLAG_CLOSED_GOP
)

type Flags2 int

const (
	Flag2Fast              Flags2 = C.CODEC_FLAG2_FAST
	Flag2NoOutput          Flags2 = C.CODEC_FLAG2_NO_OUTPUT
	Flag2LocalHeader       Flags2 = C.CODEC_FLAG2_LOCAL_HEADER
	Flag2DropFrameTimecode Flags2 = C.CODEC_FLAG2_DROP_FRAME_TIMECODE
	Flag2IgnoreCrop        Flags2 = C.CODEC_FLAG2_IGNORE_CROP
	Flag2Chunks            Flags2 = C.CODEC_FLAG2_CHUNKS
	Flag2ShowAll           Flags2 = C.CODEC_FLAG2_SHOW_ALL
	Flag2ExportMvs         Flags2 = C.CODEC_FLAG2_EXPORT_MVS
	Flag2SkipManual        Flags2 = C.CODEC_FLAG2_SKIP_MANUAL
)

type Capabilities int

const (
	CapabilityDrawHorizBand     Capabilities = C.CODEC_CAP_DRAW_HORIZ_BAND
	CapabilityDR1               Capabilities = C.CODEC_CAP_DR1
	CapabilityTruncated         Capabilities = C.CODEC_CAP_TRUNCATED
	CapabilityHWAccel           Capabilities = C.GO_CODEC_CAP_HWACCEL
	CapabilityDelay             Capabilities = C.CODEC_CAP_DELAY
	CapabilitySmallLastFrame    Capabilities = C.CODEC_CAP_SMALL_LAST_FRAME
	CapabilityHWAccelVDPAU      Capabilities = C.GO_CODEC_CAP_HWACCEL_VDPAU
	CapabilitySubframes         Capabilities = C.CODEC_CAP_SUBFRAMES
	CapabilityExperimental      Capabilities = C.CODEC_CAP_EXPERIMENTAL
	CapabilityChannelConf       Capabilities = C.CODEC_CAP_CHANNEL_CONF
	CapabilityFrameThreads      Capabilities = C.CODEC_CAP_FRAME_THREADS
	CapabilitySliceThreads      Capabilities = C.CODEC_CAP_SLICE_THREADS
	CapabilityParamChange       Capabilities = C.CODEC_CAP_PARAM_CHANGE
	CapabilityAutoThreads       Capabilities = C.CODEC_CAP_AUTO_THREADS
	CapabilityVariableFrameSize Capabilities = C.CODEC_CAP_VARIABLE_FRAME_SIZE
	CapabilityIntraOnly         Capabilities = C.CODEC_CAP_INTRA_ONLY
	CapabilityLossless          Capabilities = C.CODEC_CAP_LOSSLESS
)
