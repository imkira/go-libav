package avcodec

//#include <libavutil/avutil.h>
//#include <libavcodec/avcodec.h>
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
//#ifdef FF_DCT_INT
//#define GO_FF_DCT_INT FF_DCT_INT
//#else
//#define GO_FF_DCT_INT 0
//#endif
//
//#ifdef FF_IDCT_SH4
//#define GO_FF_IDCT_SH4 FF_IDCT_SH4
//#else
//#define GO_FF_IDCT_SH4 0
//#endif
//
//#ifdef FF_IDCT_IPP
//#define GO_FF_IDCT_IPP FF_IDCT_IPP
//#else
//#define GO_FF_IDCT_IPP 0
//#endif
//
//#ifdef FF_IDCT_XVIDMMX
//#define GO_FF_IDCT_XVIDMMX FF_IDCT_XVIDMMX
//#else
//#define GO_FF_IDCT_XVIDMMX 0
//#endif
//
//#ifdef FF_IDCT_SIMPLEVIS
//#define GO_FF_IDCT_SIMPLEVIS FF_IDCT_SIMPLEVIS
//#else
//#define GO_FF_IDCT_SIMPLEVIS 0
//#endif
//
//#ifdef FF_IDCT_SIMPLEALPHA
//#define GO_FF_IDCT_SIMPLEALPHA FF_IDCT_SIMPLEALPHA
//#else
//#define GO_FF_IDCT_SIMPLEALPHA 0
//#endif
//
//static const AVPacketSideData *go_av_packetsidedata_get(const AVPacketSideData *side_data, int n)
//{
//  return &side_data[n];
//}
//
//static const AVRational *go_av_rational_get(const AVRational *r, int n)
//{
//  if (r == NULL)
//  {
//    return NULL;
//  }
//  return &r[n];
//}
//
//static enum AVPixelFormat *go_av_pixfmt_get(enum AVPixelFormat *pixfmt, int n)
//{
//  if (pixfmt == NULL)
//  {
//    return NULL;
//  }
//  return &pixfmt[n];
//}
//
//static enum AVSampleFormat *go_av_samplefmt_get(enum AVSampleFormat *samplefmt, int n)
//{
//  if (samplefmt == NULL)
//  {
//    return NULL;
//  }
//  return &samplefmt[n];
//}
//
//static const AVProfile *go_av_profile_get(const AVProfile *profile, int n)
//{
//  if (profile == NULL)
//  {
//    return NULL;
//  }
//  return &profile[n];
//}
//
//static int *go_av_int_get(int *arr, int n)
//{
//  if (arr == NULL)
//  {
//    return NULL;
//  }
//  return &arr[n];
//}
//
//static uint64_t *go_av_uint64_get(uint64_t *arr, int n)
//{
//  if (arr == NULL)
//  {
//    return NULL;
//  }
//  return &arr[n];
//}
//
//static const char* get_list_at(const char **list, const int idx)
//{
//  return list[idx];
//}
//
//static void go_av_packet_free(void *pPkt) {
//	av_packet_free((AVPacket**)&pPkt);
//}
//
// int GO_AVCODEC_VERSION_MAJOR = LIBAVCODEC_VERSION_MAJOR;
// int GO_AVCODEC_VERSION_MINOR = LIBAVCODEC_VERSION_MINOR;
// int GO_AVCODEC_VERSION_MICRO = LIBAVCODEC_VERSION_MICRO;
//
// #cgo pkg-config: libavcodec libavutil
import "C"

import (
	"errors"
	"strings"
	"unsafe"

	"github.com/baohavan/go-libav/avutil"
)

var (
	ErrAllocationError         = errors.New("allocation error")
	ErrEncoderNotFound         = errors.New("encoder not found")
	ErrDecoderNotFound         = errors.New("decoder not found")
	ErrBitStreamFilterNotFound = errors.New("bitstreamfilter not found")
)

type CodecID C.enum_AVCodecID

const (
	CodecIDNone  CodecID = C.AV_CODEC_ID_NONE
	CodecIDMJpeg CodecID = C.AV_CODEC_ID_MJPEG
	CodecIDLJpeg CodecID = C.AV_CODEC_ID_LJPEG
)

type Flags int64

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

type Capabilities int64

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

type Compliance int

const (
	ComplianceVeryStrict   Compliance = C.FF_COMPLIANCE_VERY_STRICT
	ComplianceStrict       Compliance = C.FF_COMPLIANCE_STRICT
	ComplianceNormal       Compliance = C.FF_COMPLIANCE_NORMAL
	ComplianceUnofficial   Compliance = C.FF_COMPLIANCE_UNOFFICIAL
	ComplianceExperimental Compliance = C.FF_COMPLIANCE_EXPERIMENTAL
)

type PacketSideDataType C.enum_AVPacketSideDataType

const (
	PacketSideDataPalette                 PacketSideDataType = C.AV_PKT_DATA_PALETTE
	PacketSideDataNewExtraData            PacketSideDataType = C.AV_PKT_DATA_NEW_EXTRADATA
	PacketSideDataParamChange             PacketSideDataType = C.AV_PKT_DATA_PARAM_CHANGE
	PacketSideDataH263MBInfo              PacketSideDataType = C.AV_PKT_DATA_H263_MB_INFO
	PacketSideDataReplayGain              PacketSideDataType = C.AV_PKT_DATA_REPLAYGAIN
	PacketSideDataDisplayMatrix           PacketSideDataType = C.AV_PKT_DATA_DISPLAYMATRIX
	PacketSideDataStereo3D                PacketSideDataType = C.AV_PKT_DATA_STEREO3D
	PacketSideDataAudioServiceType        PacketSideDataType = C.AV_PKT_DATA_AUDIO_SERVICE_TYPE
	PacketSideDataSkipSamples             PacketSideDataType = C.AV_PKT_DATA_SKIP_SAMPLES
	PacketSideDataJPDualMono              PacketSideDataType = C.AV_PKT_DATA_JP_DUALMONO
	PacketSideDataStringsMetaData         PacketSideDataType = C.AV_PKT_DATA_STRINGS_METADATA
	PacketSideDataSubtitlePosition        PacketSideDataType = C.AV_PKT_DATA_SUBTITLE_POSITION
	PacketSideDataMatroskaBlockAdditional PacketSideDataType = C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL
	PacketSideDataWebVTTIdentifier        PacketSideDataType = C.AV_PKT_DATA_WEBVTT_IDENTIFIER
	PacketSideDataWebVTTSettings          PacketSideDataType = C.AV_PKT_DATA_WEBVTT_SETTINGS
	PacketSideDataMetaDataUpdate          PacketSideDataType = C.AV_PKT_DATA_METADATA_UPDATE
)

type PacketFlags int

const (
	PacketFlagKey     PacketFlags = C.AV_PKT_FLAG_KEY
	PacketFlagCorrupt PacketFlags = C.AV_PKT_FLAG_CORRUPT
)

type DCTAlgorithm int

const (
	DCTAlgorithmAuto    DCTAlgorithm = C.FF_DCT_AUTO
	DCTAlgorithmFastInt DCTAlgorithm = C.FF_DCT_FASTINT
	DCTAlgorithmInt     DCTAlgorithm = C.GO_FF_DCT_INT
	DCTAlgorithmMMX     DCTAlgorithm = C.FF_DCT_MMX
	DCTAlgorithmAltiVec DCTAlgorithm = C.FF_DCT_ALTIVEC
	DCTAlgorithmFAAN    DCTAlgorithm = C.FF_DCT_FAAN
)

type IDCTAlgorithm int

const (
	IDCTAlgorithmAuto          IDCTAlgorithm = C.FF_IDCT_AUTO
	IDCTAlgorithmInt           IDCTAlgorithm = C.FF_IDCT_INT
	IDCTAlgorithmSimple        IDCTAlgorithm = C.FF_IDCT_SIMPLE
	IDCTAlgorithmSimpleMMX     IDCTAlgorithm = C.FF_IDCT_SIMPLEMMX
	IDCTAlgorithmARM           IDCTAlgorithm = C.FF_IDCT_ARM
	IDCTAlgorithmAltiVec       IDCTAlgorithm = C.FF_IDCT_ALTIVEC
	IDCTAlgorithmSH4           IDCTAlgorithm = C.GO_FF_IDCT_SH4
	IDCTAlgorithmSimpleARM     IDCTAlgorithm = C.FF_IDCT_SIMPLEARM
	IDCTAlgorithmIPP           IDCTAlgorithm = C.GO_FF_IDCT_IPP
	IDCTAlgorithmXvid          IDCTAlgorithm = C.FF_IDCT_XVID
	IDCTAlgorithmXvidMMX       IDCTAlgorithm = C.GO_FF_IDCT_XVIDMMX
	IDCTAlgorithmSimpleARMv5TE IDCTAlgorithm = C.FF_IDCT_SIMPLEARMV5TE
	IDCTAlgorithmSimpleARMv6   IDCTAlgorithm = C.FF_IDCT_SIMPLEARMV6
	IDCTAlgorithmSimpleVis     IDCTAlgorithm = C.GO_FF_IDCT_SIMPLEVIS
	IDCTAlgorithmFAAN          IDCTAlgorithm = C.FF_IDCT_FAAN
	IDCTAlgorithmSimpleNEON    IDCTAlgorithm = C.FF_IDCT_SIMPLENEON
	IDCTAlgorithmSimpleAlpha   IDCTAlgorithm = C.GO_FF_IDCT_SIMPLEALPHA
	IDCTAlgorithmSimpleAuto    IDCTAlgorithm = C.FF_IDCT_SIMPLEAUTO
)

type ThreadType int

const (
	ThreadTypeFrame ThreadType = C.FF_THREAD_FRAME
	ThreadTypeSlice ThreadType = C.FF_THREAD_SLICE
)

const (
	ProfileUnknown int = C.FF_PROFILE_UNKNOWN
)

type SubtitlesEncodingMode int

const (
	SubtitlesEncodingModeDoNothing  SubtitlesEncodingMode = C.FF_SUB_CHARENC_MODE_DO_NOTHING
	SubtitlesEncodingModeAutomatic  SubtitlesEncodingMode = C.FF_SUB_CHARENC_MODE_AUTOMATIC
	SubtitlesEncodingModePreDecoder SubtitlesEncodingMode = C.FF_SUB_CHARENC_MODE_PRE_DECODER
)

type CodecProps int

const (
	CodecPropIntraOnly CodecProps = C.AV_CODEC_PROP_INTRA_ONLY
	CodecPropLossy     CodecProps = C.AV_CODEC_PROP_LOSSY
	CodecPropLossless  CodecProps = C.AV_CODEC_PROP_LOSSLESS
	CodecPropReorder   CodecProps = C.AV_CODEC_PROP_REORDER
	CodecPropBitmapSub CodecProps = C.AV_CODEC_PROP_BITMAP_SUB
	CodecPropTextSub   CodecProps = C.AV_CODEC_PROP_TEXT_SUB
)

func init() {
	RegisterAll()
}

func Version() (int, int, int) {
	return int(C.GO_AVCODEC_VERSION_MAJOR), int(C.GO_AVCODEC_VERSION_MINOR), int(C.GO_AVCODEC_VERSION_MICRO)
}


func RegisterAll() {
	C.avcodec_register_all()
}

type PacketSideData struct {
	CAVPacketSideData *C.AVPacketSideData
}

func NewPacketSideDataFromC(cPSD unsafe.Pointer) *PacketSideData {
	return &PacketSideData{CAVPacketSideData: (*C.AVPacketSideData)(cPSD)}
}

func (psd *PacketSideData) Data() unsafe.Pointer {
	return unsafe.Pointer(psd.CAVPacketSideData.data)
}

func (psd *PacketSideData) SetData(data unsafe.Pointer) {
	psd.CAVPacketSideData.data = (*C.uint8_t)(data)
}

func (psd *PacketSideData) Size() int {
	return int(psd.CAVPacketSideData.size)
}

func (psd *PacketSideData) SetSize(size int) {
	psd.CAVPacketSideData.size = (C.int)(size)
}

func (psd *PacketSideData) Type() PacketSideDataType {
	return PacketSideDataType(psd.CAVPacketSideData._type)
}

func (psd *PacketSideData) SetType(t PacketSideDataType) {
	psd.CAVPacketSideData._type = (C.enum_AVPacketSideDataType)(t)
}

type Packet struct {
	//CAVPacket *C.AVPacket
	CAVPacket uintptr
}

func NewPacket() (Packet, error) {
	cPkt := uintptr(unsafe.Pointer(C.av_packet_alloc()))
	if cPkt == 0 {
		return Packet{}, ErrAllocationError
	}
	return NewPacketFromC(uintptr(unsafe.Pointer(cPkt))), nil
}

func NewPacketFromC(cPkt uintptr) Packet {
	return Packet{CAVPacket: cPkt}
}

func (pkt *Packet) Packet() *C.AVPacket {
	return (*C.AVPacket)(unsafe.Pointer(pkt.CAVPacket))
}

func (pkt *Packet) Free() {
	if pkt.CAVPacket != 0 {
		C.go_av_packet_free(unsafe.Pointer(pkt.CAVPacket))
		pkt.CAVPacket = 0
	}
}

func (pkt *Packet) Ref(dst *Packet) error {
	code := C.av_packet_ref(dst.Packet(), pkt.Packet())
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (pkt *Packet) Unref() {
	C.av_packet_unref(pkt.Packet())
}

func (pkt *Packet) ConsumeData(size int) {
	data := unsafe.Pointer(pkt.Packet().data)
	if data != nil {
		pkt.Packet().size -= C.int(size)
		pkt.Packet().data = (*C.uint8_t)(unsafe.Pointer(uintptr(data) + uintptr(size)))
	}
}

func (pkt *Packet) RescaleTime(srcTimeBase, dstTimeBase *avutil.Rational) {
	src := (*C.AVRational)(unsafe.Pointer(&srcTimeBase.CAVRational))
	dst := (*C.AVRational)(unsafe.Pointer(&dstTimeBase.CAVRational))
	C.av_packet_rescale_ts(pkt.Packet(), *src, *dst)
}

func (pkt *Packet) RescaleTime2(srcTimeBase, dstTimeBase *avutil.Rational) {
	src := (*C.AVRational)(unsafe.Pointer(&srcTimeBase.CAVRational))
	dst := (*C.AVRational)(unsafe.Pointer(&dstTimeBase.CAVRational))

	pkt.SetPTS(int64(C.av_rescale_q_rnd(pkt.Packet().pts, *src, *dst, C.AV_ROUND_NEAR_INF|C.AV_ROUND_PASS_MINMAX)))
	pkt.SetDTS(int64(C.av_rescale_q_rnd(pkt.Packet().dts, *src, *dst, C.AV_ROUND_NEAR_INF|C.AV_ROUND_PASS_MINMAX)))
	pkt.SetDuration(int64(C.av_rescale_q(pkt.Packet().duration, *src, *dst)))
	pkt.SetPosition(-1)
}

func (pkt *Packet) RescalePTS(pts int64, base *avutil.Rational) int64 {
	src := (*C.AVRational)(unsafe.Pointer(&base.CAVRational))
	return int64(C.av_rescale_q((C.int64_t)(pts), *src, C.AV_TIME_BASE_Q));
}

func (pkt *Packet) RescalePTS2(pts int64, base *avutil.Rational) int64 {
	src := (*C.AVRational)(unsafe.Pointer(&base.CAVRational))
	return int64(C.av_rescale_q((C.int64_t)(pts), C.AV_TIME_BASE_Q, *src));
}

func (pkt *Packet) PTS() int64 {
	return int64(pkt.Packet().pts)
}

func (pkt *Packet) SetPTS(pts int64) {
	pkt.Packet().pts = (C.int64_t)(pts)
}

func (pkt *Packet) DTS() int64 {
	return int64(pkt.Packet().dts)
}

func (pkt *Packet) SetDTS(dts int64) {
	pkt.Packet().dts = (C.int64_t)(dts)
}

func (pkt *Packet) Duration() int64 {
	return int64(pkt.Packet().duration)
}

func (pkt *Packet) SetDuration(duration int64) {
	pkt.Packet().duration = (C.int64_t)(duration)
}

func (pkt *Packet) Data() unsafe.Pointer {
	return unsafe.Pointer(pkt.Packet().data)
}

func (pkt *Packet) SetData(data unsafe.Pointer) {
	pkt.Packet().data = (*C.uint8_t)(data)
}

func (pkt *Packet) Size() int {
	return int(pkt.Packet().size)
}

func (pkt *Packet) SetSize(size int) {
	pkt.Packet().size = (C.int)(size)
}

func (pkt *Packet) SideData() []*PacketSideData {
	count := int(pkt.Packet().side_data_elems)
	if count <= 0 {
		return nil
	}
	psds := make([]*PacketSideData, 0, count)
	for i := 0; i < count; i++ {
		cPSD := C.go_av_packetsidedata_get(pkt.Packet().side_data, C.int(i))
		psd := NewPacketSideDataFromC(unsafe.Pointer(cPSD))
		psds = append(psds, psd)
	}
	return psds
}

func (pkt *Packet) StreamIndex() int {
	return int(pkt.Packet().stream_index)
}

func (pkt *Packet) SetStreamIndex(streamIndex int) {
	pkt.Packet().stream_index = (C.int)(streamIndex)
}

func (pkt *Packet) Flags() PacketFlags {
	return PacketFlags(pkt.Packet().flags)
}

func (pkt *Packet) SetFlags(flags PacketFlags) {
	pkt.Packet().flags = (C.int)(flags)
}

func (pkt *Packet) Position() int64 {
	return int64(pkt.Packet().pos)
}

func (pkt *Packet) SetPosition(position int64) {
	pkt.Packet().pos = (C.int64_t)(position)
}

func (pkt *Packet) ConvergenceDuration() int64 {
	return int64(pkt.Packet().convergence_duration)
}

func (pkt *Packet) SetConvergenceDuration(convergenceDuration int64) {
	pkt.Packet().convergence_duration = (C.int64_t)(convergenceDuration)
}

type Profile struct {
	CAVProfile *C.AVProfile
}

func NewProfileFromC(cProfile unsafe.Pointer) *Profile {
	return &Profile{CAVProfile: (*C.AVProfile)(cProfile)}
}

func (p *Profile) Name() string {
	name, _ := p.NameOK()
	return name
}

func (p *Profile) NameOK() (string, bool) {
	return cStringToStringOk(p.CAVProfile.name)
}

func (p *Profile) ID() int {
	return int(p.CAVProfile.profile)
}

type Codec struct {
	//CAVCodec *C.AVCodec
	CAVCodec uintptr
}

func NewCodecFromC(cCodec uintptr) *Codec {
	return &Codec{CAVCodec: cCodec}
}

func FindEncoderByID(codecID CodecID) *Codec {
	cCodec := uintptr(unsafe.Pointer(C.avcodec_find_encoder((C.enum_AVCodecID)(codecID))))
	if cCodec == 0 {
		return nil
	}
	return NewCodecFromC(cCodec)
}

func FindEncoderByName(name string) *Codec {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cCodec := uintptr(unsafe.Pointer(C.avcodec_find_encoder_by_name(cName)))
	if cCodec == 0 {
		return nil
	}
	return NewCodecFromC(cCodec)
}

func FindDecoderByID(codecID CodecID) *Codec {
	cCodec := uintptr(unsafe.Pointer(C.avcodec_find_decoder((C.enum_AVCodecID)(codecID))))
	if cCodec == 0 {
		return nil
	}
	return NewCodecFromC(cCodec)
}

func FindDecoderByName(name string) *Codec {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cCodec := uintptr(unsafe.Pointer(C.avcodec_find_decoder_by_name(cName)))
	if cCodec == 0 {
		return nil
	}
	return NewCodecFromC(cCodec)
}

func (c *Codec) PrivateClass() *avutil.Class {
	if c.Codec().priv_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(c.Codec().priv_class))
}

func (c *Codec) Codec() *C.AVCodec {
	return (*C.AVCodec)(unsafe.Pointer(c.CAVCodec))
}

func (c *Codec) Name() string {
	str, _ := c.NameOk()
	return str
}

func (c *Codec) NameOk() (string, bool) {
	return cStringToStringOk(c.Codec().name)
}

func (c *Codec) Type() avutil.MediaType {
	return (avutil.MediaType)(c.Codec()._type)
}

func (c *Codec) ID() CodecID {
	return CodecID(c.Codec().id)
}

func (c *Codec) Capabilities() Capabilities {
	return Capabilities(c.Codec().capabilities)
}

func (c *Codec) SupportedFrameRates() []*avutil.Rational {
	var frameRates []*avutil.Rational
	for i := 0; ; i++ {
		cFrameRate := C.go_av_rational_get(c.Codec().supported_framerates, C.int(i))
		if cFrameRate == nil || (cFrameRate.num == 0 && cFrameRate.den == 0) {
			break
		}
		frameRate := avutil.NewRationalFromC(unsafe.Pointer(cFrameRate))
		frameRates = append(frameRates, frameRate)
	}
	return frameRates
}

func (c *Codec) SupportedPixelFormats() []avutil.PixelFormat {
	var pixelFormats []avutil.PixelFormat
	for i := 0; ; i++ {
		cPixelFormat := C.go_av_pixfmt_get(c.Codec().pix_fmts, C.int(i))
		if cPixelFormat == nil {
			break
		}
		pixelFormat := (avutil.PixelFormat)(*cPixelFormat)
		if pixelFormat == -1 {
			break
		}
		pixelFormats = append(pixelFormats, pixelFormat)
	}
	return pixelFormats
}

func (c *Codec) SupportedSampleRates() []int {
	var sampleRates []int
	for i := 0; ; i++ {
		cSampleRate := C.go_av_int_get(c.Codec().supported_samplerates, C.int(i))
		if cSampleRate == nil {
			break
		}
		sampleRate := (int)(*cSampleRate)
		if sampleRate == 0 {
			break
		}
		sampleRates = append(sampleRates, sampleRate)
	}
	return sampleRates
}

func (c *Codec) SupportedSampleFormats() []avutil.SampleFormat {
	var sampleFormats []avutil.SampleFormat
	for i := 0; ; i++ {
		cSampleFormat := C.go_av_samplefmt_get(c.Codec().sample_fmts, C.int(i))
		if cSampleFormat == nil {
			break
		}
		sampleFormat := (avutil.SampleFormat)(*cSampleFormat)
		if sampleFormat == -1 {
			break
		}
		sampleFormats = append(sampleFormats, sampleFormat)
	}
	return sampleFormats
}

func (c *Codec) SupportedChannelLayouts() []avutil.ChannelLayout {
	var channelLayouts []avutil.ChannelLayout
	for i := 0; ; i++ {
		cChannelLayout := C.go_av_uint64_get(c.Codec().channel_layouts, C.int(i))
		if cChannelLayout == nil {
			break
		}
		channelLayout := (avutil.ChannelLayout)(*cChannelLayout)
		if channelLayout == 0 {
			break
		}
		channelLayouts = append(channelLayouts, channelLayout)
	}
	return channelLayouts
}

func (c *Codec) Profiles() []*Profile {
	var profiles []*Profile
	for i := 0; ; i++ {
		cProfile := C.go_av_profile_get(c.Codec().profiles, C.int(i))
		if cProfile == nil || int(cProfile.profile) == ProfileUnknown {
			break
		}
		profile := NewProfileFromC(unsafe.Pointer(cProfile))
		profiles = append(profiles, profile)
	}
	return profiles
}

func (c *Codec) ProfileName(id int) string {
	name, _ := c.ProfileNameOK(id)
	return name
}

func (c *Codec) ProfileNameOK(id int) (string, bool) {
	return cStringToStringOk(C.av_get_profile_name(c.Codec(), (C.int)(id)))
}

type Context struct {
	//CAVCodecContext *C.AVCodecContext
	//*avutil.OptionAccessor
	CAVCodecContext uintptr
}

func NewContextWithCodec(codec *Codec) (*Context, error) {
	var cCodec *C.AVCodec
	if codec != nil {
		cCodec = codec.Codec()
	}
	cCtx := uintptr(unsafe.Pointer(C.avcodec_alloc_context3(cCodec)))
	if cCtx == 0 {
		return nil, ErrAllocationError
	}
	return NewContextFromC(cCtx), nil
}

func NewContextFromC(cCtx uintptr) *Context {
	return &Context{
		CAVCodecContext: cCtx,
		//OptionAccessor:  avutil.NewOptionAccessor(cCtx, false),
	}
}

func (ctx *Context) CodeContext() (*C.AVCodecContext) {
	if ctx.CAVCodecContext != 0 {
		return (*C.AVCodecContext)(unsafe.Pointer(ctx.CAVCodecContext))
	}
	return nil
}

func (ctx *Context) Free() {
	if ctx.CAVCodecContext != 0 {
		C.avcodec_free_context((**C.AVCodecContext)(unsafe.Pointer(&ctx.CAVCodecContext)))
	}
}

func (ctx *Context) Open(options *avutil.Dictionary) error {
	return ctx.OpenWithCodec(nil, options)
}

func (ctx *Context) OpenWithCodec(codec *Codec, options *avutil.Dictionary) error {
	var cCodec *C.AVCodec
	if codec != nil {
		cCodec = codec.Codec()
	}
	var cOptions **C.AVDictionary
	if options != nil {
		cOptions = (**C.AVDictionary)(options.Pointer())
	}
	code := C.avcodec_open2(ctx.CodeContext(), cCodec, cOptions)
	if code < 0 {
		return avutil.NewErrorFromCode(avutil.ErrorCode(code))
	}
	return nil
}

func (ctx *Context) OpenForEncoding(options *avutil.Dictionary) error {
	encoder := FindEncoderByID(ctx.CodecID())
	if encoder == nil {
		return ErrEncoderNotFound
	}
	return ctx.OpenWithCodec(encoder, options)
}

func (ctx *Context) OpenForDecoding(options *avutil.Dictionary) error {
	decoder := FindDecoderByID(ctx.CodecID())
	if decoder == nil {
		return ErrDecoderNotFound
	}
	return ctx.OpenWithCodec(decoder, options)
}

func (ctx *Context) Close() {
	C.avcodec_close(ctx.CodeContext())
}

func (ctx *Context) Class() *avutil.Class {
	if ctx.CodeContext().av_class == nil {
		return nil
	}
	return avutil.NewClassFromC(unsafe.Pointer(ctx.CodeContext().av_class))
}

func (ctx *Context) CodecType() avutil.MediaType {
	return (avutil.MediaType)(ctx.CodeContext().codec_type)
}

func (ctx *Context) SetCodecType(codecType avutil.MediaType) {
	ctx.CodeContext().codec_type = (C.enum_AVMediaType)(codecType)
}

func (ctx *Context) Codec() *Codec {
	if ctx.CodeContext().codec == nil {
		return nil
	}
	return NewCodecFromC(uintptr(unsafe.Pointer(ctx.CodeContext().codec)))
}

func (ctx *Context) SetCodec(codec *Codec) {
	var cCodec *C.AVCodec
	if codec != nil {
		cCodec = codec.Codec()
	}
	ctx.CodeContext().codec = cCodec
}

func (ctx *Context) CodecID() CodecID {
	return (CodecID)(ctx.CodeContext().codec_id)
}

func (ctx *Context) SetCodecID(id CodecID) {
	ctx.CodeContext().codec_id = (C.enum_AVCodecID)(id)
}

func (ctx *Context) CodecTag() uint {
	return uint(ctx.CodeContext().codec_tag)
}

func (ctx *Context) SetCodecTag(codecTag uint) {
	ctx.CodeContext().codec_tag = (C.uint)(codecTag)
}

func (ctx *Context) PrivData() unsafe.Pointer {
	return unsafe.Pointer(ctx.CodeContext().priv_data)
}

func (ctx *Context) SetPrivData(privData unsafe.Pointer) {
	ctx.CodeContext().priv_data = (unsafe.Pointer)(privData)
}

func (ctx *Context) Opaque() unsafe.Pointer {
	return unsafe.Pointer(ctx.CodeContext().opaque)
}

func (ctx *Context) SetOpaque(opaque unsafe.Pointer) {
	ctx.CodeContext().opaque = opaque
}

func (ctx *Context) BitRate() int64 {
	return int64(ctx.CodeContext().bit_rate)
}

func (ctx *Context) SetBitRate(bitRate int64) {
	ctx.CodeContext().bit_rate = (C.int64_t)(bitRate)
}

func (ctx *Context) BitRateTolerance() int {
	return int(ctx.CodeContext().bit_rate_tolerance)
}

func (ctx *Context) SetBitRateTolerance(bitRateTolerance int) {
	ctx.CodeContext().bit_rate_tolerance = (C.int)(bitRateTolerance)
}

func (ctx *Context) GlobalQuality() int {
	return int(ctx.CodeContext().global_quality)
}

func (ctx *Context) SetGlobalQuality(globalQuality int) {
	ctx.CodeContext().global_quality = (C.int)(globalQuality)
}

func (ctx *Context) CompressionLevel() int {
	return int(ctx.CodeContext().compression_level)
}

func (ctx *Context) SetCompressionLevel(compressionLevel int) {
	ctx.CodeContext().compression_level = (C.int)(compressionLevel)
}

func (ctx *Context) Flags() Flags {
	return Flags(ctx.CodeContext().flags)
}

func (ctx *Context) SetFlags(flags Flags) {
	ctx.CodeContext().flags = (C.int)(flags)
}

func (ctx *Context) Flags2() Flags2 {
	return Flags2(ctx.CodeContext().flags2)
}

func (ctx *Context) SetFlags2(flags2 Flags2) {
	ctx.CodeContext().flags2 = (C.int)(flags2)
}

func (ctx *Context) ExtraData() unsafe.Pointer {
	return unsafe.Pointer(ctx.CodeContext().extradata)
}

func (ctx *Context) SetExtraData(data unsafe.Pointer) {
	ctx.CodeContext().extradata = (*C.uint8_t)(data)
}

func (ctx *Context) ExtraDataSize() int {
	return int(ctx.CodeContext().extradata_size)
}

func (ctx *Context) SetExtraDataSize(extraDataSize int) {
	ctx.CodeContext().extradata_size = (C.int)(extraDataSize)
}

func (ctx *Context) TimeBase() *avutil.Rational {
	return avutil.NewRationalFromC(unsafe.Pointer(&ctx.CodeContext().time_base))
}

func (ctx *Context) SetTimeBase(timeBase *avutil.Rational) {
	ctx.CodeContext().time_base.num = (C.int)(timeBase.Numerator())
	ctx.CodeContext().time_base.den = (C.int)(timeBase.Denominator())
}

func (ctx *Context) TicksPerFrame() int {
	return int(ctx.CodeContext().ticks_per_frame)
}

func (ctx *Context) Delay() int {
	return int(ctx.CodeContext().delay)
}

func (ctx *Context) Width() int {
	return int(ctx.CodeContext().width)
}

func (ctx *Context) SetWidth(width int) {
	ctx.CodeContext().width = (C.int)(width)
}

func (ctx *Context) Height() int {
	return int(ctx.CodeContext().height)
}

func (ctx *Context) SetHeight(height int) {
	ctx.CodeContext().height = (C.int)(height)
}

func (ctx *Context) CodedWidth() int {
	return int(ctx.CodeContext().coded_width)
}

func (ctx *Context) SetCodedWidth(codedWidth int) {
	ctx.CodeContext().coded_width = (C.int)(codedWidth)
}

func (ctx *Context) CodedHeight() int {
	return int(ctx.CodeContext().coded_height)
}

func (ctx *Context) SetCodedHeight(codedHeight int) {
	ctx.CodeContext().coded_height = (C.int)(codedHeight)
}

func (ctx *Context) GOPSize() int {
	return int(ctx.CodeContext().gop_size)
}

func (ctx *Context) SetGOPSize(gOPSize int) {
	ctx.CodeContext().gop_size = (C.int)(gOPSize)
}

func (ctx *Context) PixelFormat() avutil.PixelFormat {
	return (avutil.PixelFormat)(ctx.CodeContext().pix_fmt)
}

func (ctx *Context) SetPixelFormat(pixelFormat avutil.PixelFormat) {
	ctx.CodeContext().pix_fmt = (C.enum_AVPixelFormat)(pixelFormat)
}

func (ctx *Context) MaxBFrames() int {
	return int(ctx.CodeContext().max_b_frames)
}

func (ctx *Context) SetMaxBFrames(maxBFrames int) {
	ctx.CodeContext().max_b_frames = (C.int)(maxBFrames)
}

func (ctx *Context) BQuantFactor() float32 {
	return float32(ctx.CodeContext().b_quant_factor)
}

func (ctx *Context) SetBQuantFactor(bQuantFactor float32) {
	ctx.CodeContext().b_quant_factor = (C.float)(bQuantFactor)
}

func (ctx *Context) BFrameStrategy() int {
	return int(ctx.CodeContext().b_frame_strategy)
}

func (ctx *Context) SetBFrameStrategy(bFrameStrategy int) {
	ctx.CodeContext().b_frame_strategy = (C.int)(bFrameStrategy)
}

func (ctx *Context) BQuantOffset() float32 {
	return float32(ctx.CodeContext().b_quant_offset)
}

func (ctx *Context) SetBQuantOffset(bQuantOffset float32) {
	ctx.CodeContext().b_quant_offset = (C.float)(bQuantOffset)
}

func (ctx *Context) HasBFrames() int {
	return int(ctx.CodeContext().has_b_frames)
}

func (ctx *Context) SetHasBFrames(hasBFrames int) {
	ctx.CodeContext().has_b_frames = (C.int)(hasBFrames)
}

func (ctx *Context) MPEGQuant() int {
	return int(ctx.CodeContext().mpeg_quant)
}

func (ctx *Context) SetMPEGQuant(mPEGQuant int) {
	ctx.CodeContext().mpeg_quant = (C.int)(mPEGQuant)
}

func (ctx *Context) IQuantFactor() float32 {
	return float32(ctx.CodeContext().i_quant_factor)
}

func (ctx *Context) SetIQuantFactor(iQuantFactor float32) {
	ctx.CodeContext().i_quant_factor = (C.float)(iQuantFactor)
}

func (ctx *Context) IQuantOffset() float32 {
	return float32(ctx.CodeContext().i_quant_offset)
}

func (ctx *Context) SetIQuantOffset(iQuantOffset float32) {
	ctx.CodeContext().i_quant_offset = (C.float)(iQuantOffset)
}

func (ctx *Context) LumiMasking() float32 {
	return float32(ctx.CodeContext().lumi_masking)
}

func (ctx *Context) SetLumiMasking(lumiMasking float32) {
	ctx.CodeContext().lumi_masking = (C.float)(lumiMasking)
}

func (ctx *Context) TemporalCplxMasking() float32 {
	return float32(ctx.CodeContext().temporal_cplx_masking)
}

func (ctx *Context) SetTemporalCplxMasking(temporalCplxMasking float32) {
	ctx.CodeContext().temporal_cplx_masking = (C.float)(temporalCplxMasking)
}

func (ctx *Context) SpatialCplxMasking() float32 {
	return float32(ctx.CodeContext().spatial_cplx_masking)
}

func (ctx *Context) SetSpatialCplxMasking(spatialCplxMasking float32) {
	ctx.CodeContext().spatial_cplx_masking = (C.float)(spatialCplxMasking)
}

func (ctx *Context) PMasking() float32 {
	return float32(ctx.CodeContext().p_masking)
}

func (ctx *Context) SetPMasking(pMasking float32) {
	ctx.CodeContext().p_masking = (C.float)(pMasking)
}

func (ctx *Context) DarkMasking() float32 {
	return float32(ctx.CodeContext().dark_masking)
}

func (ctx *Context) SetDarkMasking(darkMasking float32) {
	ctx.CodeContext().dark_masking = (C.float)(darkMasking)
}

func (ctx *Context) SliceCount() int {
	return int(ctx.CodeContext().slice_count)
}

func (ctx *Context) SetSliceCount(sliceCount int) {
	ctx.CodeContext().slice_count = (C.int)(sliceCount)
}

func (ctx *Context) PredictionMethod() int {
	return int(ctx.CodeContext().prediction_method)
}

func (ctx *Context) SetPredictionMethod(predictionMethod int) {
	ctx.CodeContext().prediction_method = (C.int)(predictionMethod)
}

func (ctx *Context) SliceOffset() unsafe.Pointer {
	return unsafe.Pointer(ctx.CodeContext().slice_offset)
}

func (ctx *Context) SetSliceOffset(sliceOffset unsafe.Pointer) {
	ctx.CodeContext().slice_offset = (*C.int)(sliceOffset)
}

func (ctx *Context) SampleAspectRatio() *avutil.Rational {
	return avutil.NewRationalFromC(unsafe.Pointer(&ctx.CodeContext().sample_aspect_ratio))
}

func (ctx *Context) SetSampleAspectRatio(aspectRatio *avutil.Rational) {
	ctx.CodeContext().sample_aspect_ratio.num = (C.int)(aspectRatio.Numerator())
	ctx.CodeContext().sample_aspect_ratio.den = (C.int)(aspectRatio.Denominator())
}

func (ctx *Context) MECmp() int {
	return int(ctx.CodeContext().me_cmp)
}

func (ctx *Context) SetMECmp(meCmp int) {
	ctx.CodeContext().me_cmp = (C.int)(meCmp)
}

func (ctx *Context) MESubCmp() int {
	return int(ctx.CodeContext().me_sub_cmp)
}

func (ctx *Context) SetMESubCmp(meSubCmp int) {
	ctx.CodeContext().me_sub_cmp = (C.int)(meSubCmp)
}

func (ctx *Context) MBCmp() int {
	return int(ctx.CodeContext().mb_cmp)
}

func (ctx *Context) SetMBCmp(mbCmp int) {
	ctx.CodeContext().mb_cmp = (C.int)(mbCmp)
}

func (ctx *Context) ILDCTCmp() int {
	return int(ctx.CodeContext().ildct_cmp)
}

func (ctx *Context) SetILDCTCmp(ildctCmp int) {
	ctx.CodeContext().ildct_cmp = (C.int)(ildctCmp)
}

func (ctx *Context) DiaSize() int {
	return int(ctx.CodeContext().dia_size)
}

func (ctx *Context) SetDiaSize(diaSize int) {
	ctx.CodeContext().dia_size = (C.int)(diaSize)
}

func (ctx *Context) LastPredictorCount() int {
	return int(ctx.CodeContext().last_predictor_count)
}

func (ctx *Context) SetLastPredictorCount(count int) {
	ctx.CodeContext().last_predictor_count = (C.int)(count)
}

func (ctx *Context) PreME() int {
	return int(ctx.CodeContext().pre_me)
}

func (ctx *Context) SetPreME(preME int) {
	ctx.CodeContext().pre_me = (C.int)(preME)
}

func (ctx *Context) MEPreCmp() int {
	return int(ctx.CodeContext().me_pre_cmp)
}

func (ctx *Context) SetMEPreCmp(mePreCmp int) {
	ctx.CodeContext().me_pre_cmp = (C.int)(mePreCmp)
}

func (ctx *Context) PreDiaSize() int {
	return int(ctx.CodeContext().pre_dia_size)
}

func (ctx *Context) SetPreDiaSize(preDiaSize int) {
	ctx.CodeContext().pre_dia_size = (C.int)(preDiaSize)
}

func (ctx *Context) MESubpelQuality() int {
	return int(ctx.CodeContext().me_subpel_quality)
}

func (ctx *Context) SetMESubpelQuality(meSubpelQuality int) {
	ctx.CodeContext().me_subpel_quality = (C.int)(meSubpelQuality)
}

func (ctx *Context) MERange() int {
	return int(ctx.CodeContext().me_range)
}

func (ctx *Context) SetMERange(meRange int) {
	ctx.CodeContext().me_range = (C.int)(meRange)
}

func (ctx *Context) MBDecision() int {
	return int(ctx.CodeContext().mb_decision)
}

func (ctx *Context) SetMBDecision(mbDecision int) {
	ctx.CodeContext().mb_decision = (C.int)(mbDecision)
}
func (ctx *Context) ScenechangeThreshold() int {
	return int(ctx.CodeContext().scenechange_threshold)
}

func (ctx *Context) SetScenechangeThreshold(threshold int) {
	ctx.CodeContext().scenechange_threshold = (C.int)(threshold)
}

func (ctx *Context) NoiseReduction() int {
	return int(ctx.CodeContext().noise_reduction)
}

func (ctx *Context) SetNoiseReduction(reduction int) {
	ctx.CodeContext().noise_reduction = (C.int)(reduction)
}

func (ctx *Context) IntraDCPrecision() int {
	return int(ctx.CodeContext().intra_dc_precision)
}

func (ctx *Context) SetIntraDCPrecision(precision int) {
	ctx.CodeContext().intra_dc_precision = (C.int)(precision)
}

func (ctx *Context) SkipTop() int {
	return int(ctx.CodeContext().skip_top)
}

func (ctx *Context) SetSkipTop(skip int) {
	ctx.CodeContext().skip_top = (C.int)(skip)
}

func (ctx *Context) SkipBottom() int {
	return int(ctx.CodeContext().skip_bottom)
}

func (ctx *Context) SetSkipBottom(skip int) {
	ctx.CodeContext().skip_bottom = (C.int)(skip)
}

func (ctx *Context) MBLmin() int {
	return int(ctx.CodeContext().mb_lmin)
}

func (ctx *Context) SetMBLmin(min int) {
	ctx.CodeContext().mb_lmin = (C.int)(min)
}

func (ctx *Context) MBLmax() int {
	return int(ctx.CodeContext().mb_lmax)
}

func (ctx *Context) SetMBLmax(max int) {
	ctx.CodeContext().mb_lmax = (C.int)(max)
}

func (ctx *Context) MEPenaltyCompensation() int {
	return int(ctx.CodeContext().me_penalty_compensation)
}

func (ctx *Context) SetMEPenaltyCompensation(compensation int) {
	ctx.CodeContext().me_penalty_compensation = (C.int)(compensation)
}

func (ctx *Context) BidirRefine() int {
	return int(ctx.CodeContext().bidir_refine)
}

func (ctx *Context) SetBidirRefine(refine int) {
	ctx.CodeContext().bidir_refine = (C.int)(refine)
}

func (ctx *Context) BrdScale() int {
	return int(ctx.CodeContext().brd_scale)
}

func (ctx *Context) SetBrdScale(brdScale int) {
	ctx.CodeContext().brd_scale = (C.int)(brdScale)
}

func (ctx *Context) KeyIntMin() int {
	return int(ctx.CodeContext().keyint_min)
}

func (ctx *Context) SetKeyIntMin(min int) {
	ctx.CodeContext().keyint_min = (C.int)(min)
}

func (ctx *Context) Refs() int {
	return int(ctx.CodeContext().refs)
}

func (ctx *Context) SetRefs(refs int) {
	ctx.CodeContext().refs = (C.int)(refs)
}

func (ctx *Context) ChromaOffset() int {
	return int(ctx.CodeContext().chromaoffset)
}

func (ctx *Context) SetChromaOffset(offset int) {
	ctx.CodeContext().chromaoffset = (C.int)(offset)
}

func (ctx *Context) MV0Threshold() int {
	return int(ctx.CodeContext().mv0_threshold)
}

func (ctx *Context) SetMV0Threshold(threshold int) {
	ctx.CodeContext().mv0_threshold = (C.int)(threshold)
}

func (ctx *Context) BSensitivity() int {
	return int(ctx.CodeContext().b_sensitivity)
}

func (ctx *Context) SetBSensitivity(sensivity int) {
	ctx.CodeContext().b_sensitivity = (C.int)(sensivity)
}

func (ctx *Context) ChromaSampleLocation() avutil.ChromaLocation {
	return (avutil.ChromaLocation)(ctx.CodeContext().chroma_sample_location)
}

func (ctx *Context) SetChromaSampleLocation(location avutil.ChromaLocation) {
	ctx.CodeContext().chroma_sample_location = (C.enum_AVChromaLocation)(location)
}

func (ctx *Context) Slices() int {
	return int(ctx.CodeContext().slices)
}

func (ctx *Context) SetSlices(slices int) {
	ctx.CodeContext().slices = (C.int)(slices)
}

func (ctx *Context) SampleRate() int {
	return int(ctx.CodeContext().sample_rate)
}

func (ctx *Context) SetSampleRate(rate int) {
	ctx.CodeContext().sample_rate = (C.int)(rate)
}

func (ctx *Context) Channels() int {
	return int(ctx.CodeContext().channels)
}

func (ctx *Context) SetChannels(channels int) {
	ctx.CodeContext().channels = (C.int)(channels)
}

func (ctx *Context) SampleFormat() avutil.SampleFormat {
	return (avutil.SampleFormat)(ctx.CodeContext().sample_fmt)
}

func (ctx *Context) SetSampleFormat(format avutil.SampleFormat) {
	ctx.CodeContext().sample_fmt = (C.enum_AVSampleFormat)(format)
}

func (ctx *Context) FrameSize() int {
	return int(ctx.CodeContext().frame_size)
}

func (ctx *Context) SetFrameSize(size int) {
	ctx.CodeContext().frame_size = (C.int)(size)
}

func (ctx *Context) FrameNumber() int {
	return int(ctx.CodeContext().frame_number)
}

func (ctx *Context) SetFrameNumber(number int) {
	ctx.CodeContext().frame_number = (C.int)(number)
}

func (ctx *Context) BlockAlign() int {
	return int(ctx.CodeContext().block_align)
}

func (ctx *Context) SetBlockAlign(blockAlign int) {
	ctx.CodeContext().block_align = (C.int)(blockAlign)
}

func (ctx *Context) Cutoff() int {
	return int(ctx.CodeContext().cutoff)
}

func (ctx *Context) SetCutoff(cutoff int) {
	ctx.CodeContext().cutoff = (C.int)(cutoff)
}

func (ctx *Context) ChannelLayout() avutil.ChannelLayout {
	return (avutil.ChannelLayout)(ctx.CodeContext().channel_layout)
}

func (ctx *Context) SetChannelLayout(layout avutil.ChannelLayout) {
	ctx.CodeContext().channel_layout = (C.uint64_t)(layout)
}

func (ctx *Context) RequestChannelLayout() avutil.ChannelLayout {
	return (avutil.ChannelLayout)(ctx.CodeContext().request_channel_layout)
}

func (ctx *Context) SetRequestChannelLayout(layout avutil.ChannelLayout) {
	ctx.CodeContext().request_channel_layout = (C.uint64_t)(layout)
}
func (ctx *Context) RequestSampleFormat() avutil.SampleFormat {
	return (avutil.SampleFormat)(ctx.CodeContext().request_sample_fmt)
}

func (ctx *Context) SetRequestSampleFormat(format avutil.SampleFormat) {
	ctx.CodeContext().request_sample_fmt = (C.enum_AVSampleFormat)(format)
}

func (ctx *Context) RefCountedFrames() bool {
	return ctx.CodeContext().refcounted_frames != C.int(0)
}

func (ctx *Context) SetRefCountedFrames(refCounted bool) {
	var cRefCounted C.int
	if refCounted {
		cRefCounted = C.int(1)
	}
	ctx.CodeContext().refcounted_frames = cRefCounted
}

func (ctx *Context) QCompress() float32 {
	return float32(ctx.CodeContext().qcompress)
}

func (ctx *Context) SetQCompress(qcompress float32) {
	ctx.CodeContext().qcompress = (C.float)(qcompress)
}

func (ctx *Context) QBlur() float32 {
	return float32(ctx.CodeContext().qblur)
}

func (ctx *Context) SetQBlur(qblur float32) {
	ctx.CodeContext().qblur = (C.float)(qblur)
}

func (ctx *Context) QMin() int {
	return int(ctx.CodeContext().qmin)
}

func (ctx *Context) SetQMin(qmin int) {
	ctx.CodeContext().qmin = (C.int)(qmin)
}

func (ctx *Context) QMax() int {
	return int(ctx.CodeContext().qmax)
}

func (ctx *Context) SetQMax(qmax int) {
	ctx.CodeContext().qmax = (C.int)(qmax)
}

func (ctx *Context) MaxQDiff() int {
	return int(ctx.CodeContext().max_qdiff)
}

func (ctx *Context) SetMaxQDiff(max int) {
	ctx.CodeContext().max_qdiff = (C.int)(max)
}

func (ctx *Context) RCBufferSize() int {
	return int(ctx.CodeContext().rc_buffer_size)
}

func (ctx *Context) SetRCBufferSize(size int) {
	ctx.CodeContext().rc_buffer_size = (C.int)(size)
}

func (ctx *Context) RCOverrideCount() int {
	return int(ctx.CodeContext().rc_override_count)
}

func (ctx *Context) SetRCOverrideCount(count int) {
	ctx.CodeContext().rc_override_count = (C.int)(count)
}
func (ctx *Context) RCMaxRate() int64 {
	return int64(ctx.CodeContext().rc_max_rate)
}

func (ctx *Context) SetRCMaxRate(max int64) {
	ctx.CodeContext().rc_max_rate = (C.int64_t)(max)
}

func (ctx *Context) RCMinRate() int64 {
	return int64(ctx.CodeContext().rc_min_rate)
}

func (ctx *Context) SetRCMinRate(min int64) {
	ctx.CodeContext().rc_min_rate = (C.int64_t)(min)
}

func (ctx *Context) RCMaxAvailableVBVUse() float32 {
	return float32(ctx.CodeContext().rc_max_available_vbv_use)
}

func (ctx *Context) SetRCMaxAvailableVBVUse(max float32) {
	ctx.CodeContext().rc_max_available_vbv_use = (C.float)(max)
}

func (ctx *Context) RCMinVBVOverflowUse() float32 {
	return float32(ctx.CodeContext().rc_min_vbv_overflow_use)
}

func (ctx *Context) SetRCMinVBVOverflowUse(min float32) {
	ctx.CodeContext().rc_min_vbv_overflow_use = (C.float)(min)
}

func (ctx *Context) RCInitialBufferOccupancy() int {
	return int(ctx.CodeContext().rc_initial_buffer_occupancy)
}

func (ctx *Context) SetRCInitialBufferOccupancy(initial int) {
	ctx.CodeContext().rc_initial_buffer_occupancy = (C.int)(initial)
}

func (ctx *Context) ContextModel() int {
	return int(ctx.CodeContext().context_model)
}

func (ctx *Context) SetContextModel(contextModel int) {
	ctx.CodeContext().context_model = (C.int)(contextModel)
}

func (ctx *Context) FrameSkipThreshold() int {
	return int(ctx.CodeContext().frame_skip_threshold)
}

func (ctx *Context) SetFrameSkipThreshold(threshold int) {
	ctx.CodeContext().frame_skip_threshold = (C.int)(threshold)
}

func (ctx *Context) FrameSkipFactor() int {
	return int(ctx.CodeContext().frame_skip_factor)
}

func (ctx *Context) SetFrameSkipFactor(factor int) {
	ctx.CodeContext().frame_skip_factor = (C.int)(factor)
}

func (ctx *Context) FrameSkipExp() int {
	return int(ctx.CodeContext().frame_skip_exp)
}

func (ctx *Context) SetFrameSkipExp(skip int) {
	ctx.CodeContext().frame_skip_exp = (C.int)(skip)
}

func (ctx *Context) FrameSkipCmp() int {
	return int(ctx.CodeContext().frame_skip_cmp)
}

func (ctx *Context) SetFrameSkipCmp(skip int) {
	ctx.CodeContext().frame_skip_cmp = (C.int)(skip)
}

func (ctx *Context) Trellis() int {
	return int(ctx.CodeContext().trellis)
}

func (ctx *Context) SetTrellis(trellis int) {
	ctx.CodeContext().trellis = (C.int)(trellis)
}

func (ctx *Context) MinPredictionOrder() int {
	return int(ctx.CodeContext().min_prediction_order)
}

func (ctx *Context) SetMinPredictionOrder(min int) {
	ctx.CodeContext().min_prediction_order = (C.int)(min)
}

func (ctx *Context) MaxPredictionOrder() int {
	return int(ctx.CodeContext().max_prediction_order)
}

func (ctx *Context) SetMaxPredictionOrder(max int) {
	ctx.CodeContext().max_prediction_order = (C.int)(max)
}

func (ctx *Context) TimecodeFrameStart() int64 {
	return int64(ctx.CodeContext().timecode_frame_start)
}

func (ctx *Context) SetTimecodeFrameStart(start int64) {
	ctx.CodeContext().timecode_frame_start = (C.int64_t)(start)
}

func (ctx *Context) RTPPayloadSize() int {
	return int(ctx.CodeContext().rtp_payload_size)
}

func (ctx *Context) SetRTPPayloadSize(size int) {
	ctx.CodeContext().rtp_payload_size = (C.int)(size)
}

func (ctx *Context) MVBits() int {
	return int(ctx.CodeContext().mv_bits)
}

func (ctx *Context) SetMVBits(bits int) {
	ctx.CodeContext().mv_bits = (C.int)(bits)
}

func (ctx *Context) HeaderBits() int {
	return int(ctx.CodeContext().header_bits)
}

func (ctx *Context) SetHeaderBits(bits int) {
	ctx.CodeContext().header_bits = (C.int)(bits)
}

func (ctx *Context) ITexBits() int {
	return int(ctx.CodeContext().i_tex_bits)
}

func (ctx *Context) SetITexBits(bits int) {
	ctx.CodeContext().i_tex_bits = (C.int)(bits)
}

func (ctx *Context) PTexBits() int {
	return int(ctx.CodeContext().p_tex_bits)
}

func (ctx *Context) SetPTexBits(bits int) {
	ctx.CodeContext().p_tex_bits = (C.int)(bits)
}

func (ctx *Context) ICount() int {
	return int(ctx.CodeContext().i_count)
}

func (ctx *Context) SetICount(count int) {
	ctx.CodeContext().i_count = (C.int)(count)
}

func (ctx *Context) PCount() int {
	return int(ctx.CodeContext().p_count)
}

func (ctx *Context) SetPCount(count int) {
	ctx.CodeContext().p_count = (C.int)(count)
}

func (ctx *Context) SkipCount() int {
	return int(ctx.CodeContext().skip_count)
}

func (ctx *Context) SetSkipCount(skip int) {
	ctx.CodeContext().skip_count = (C.int)(skip)
}

func (ctx *Context) MiscBits() int {
	return int(ctx.CodeContext().misc_bits)
}

func (ctx *Context) SetMiscBits(bits int) {
	ctx.CodeContext().misc_bits = (C.int)(bits)
}

func (ctx *Context) FrameBits() int {
	return int(ctx.CodeContext().frame_bits)
}

func (ctx *Context) SetFrameBits(bits int) {
	ctx.CodeContext().frame_bits = (C.int)(bits)
}

func (ctx *Context) StrictStdCompliance() Compliance {
	return Compliance(ctx.CodeContext().strict_std_compliance)
}

func (ctx *Context) SetStrictStdCompliance(compliance Compliance) {
	ctx.CodeContext().strict_std_compliance = (C.int)(compliance)
}

func (ctx *Context) ReorderedOpaque() int64 {
	return int64(ctx.CodeContext().reordered_opaque)
}

func (ctx *Context) SetReorderedOpaque(opaque int64) {
	ctx.CodeContext().reordered_opaque = (C.int64_t)(opaque)
}

func (ctx *Context) DCTAlgorithm() DCTAlgorithm {
	return DCTAlgorithm(ctx.CodeContext().dct_algo)
}

func (ctx *Context) SetDCTAlgorithm(algo DCTAlgorithm) {
	ctx.CodeContext().dct_algo = (C.int)(algo)
}

func (ctx *Context) IDCTAlgorithm() IDCTAlgorithm {
	return IDCTAlgorithm(ctx.CodeContext().idct_algo)
}

func (ctx *Context) SetIDCTAlgorithm(algo IDCTAlgorithm) {
	ctx.CodeContext().idct_algo = (C.int)(algo)
}

func (ctx *Context) BitsPerCodedSample() int {
	return int(ctx.CodeContext().bits_per_coded_sample)
}

func (ctx *Context) SetBitsPerCodedSample(bits int) {
	ctx.CodeContext().bits_per_coded_sample = (C.int)(bits)
}

func (ctx *Context) BitsPerRawSample() int {
	return int(ctx.CodeContext().bits_per_raw_sample)
}

func (ctx *Context) SetBitsPerRawSample(bits int) {
	ctx.CodeContext().bits_per_raw_sample = (C.int)(bits)
}

func (ctx *Context) LowRes() int {
	return int(ctx.CodeContext().lowres)
}

func (ctx *Context) SetLowRes(res int) {
	ctx.CodeContext().lowres = (C.int)(res)
}

func (ctx *Context) ThreadCount() int {
	return int(ctx.CodeContext().thread_count)
}

func (ctx *Context) SetThreadCount(count int) {
	ctx.CodeContext().thread_count = (C.int)(count)
}

func (ctx *Context) ThreadType() ThreadType {
	return ThreadType(ctx.CodeContext().thread_type)
}

func (ctx *Context) SetThreadType(threadType ThreadType) {
	ctx.CodeContext().thread_type = (C.int)(threadType)
}

func (ctx *Context) ActiveThreadType() ThreadType {
	return ThreadType(ctx.CodeContext().active_thread_type)
}

func (ctx *Context) SetActiveThreadType(threadType ThreadType) {
	ctx.CodeContext().active_thread_type = (C.int)(threadType)
}

func (ctx *Context) ThreadSafeCallbacks() int {
	return int(ctx.CodeContext().thread_safe_callbacks)
}

func (ctx *Context) SetThreadSafeCallbacks(count int) {
	ctx.CodeContext().thread_safe_callbacks = (C.int)(count)
}

func (ctx *Context) NSSEWeight() int {
	return int(ctx.CodeContext().nsse_weight)
}

func (ctx *Context) SetNSSEWeight(weight int) {
	ctx.CodeContext().nsse_weight = (C.int)(weight)
}

func (ctx *Context) Profile() int {
	return int(ctx.CodeContext().profile)
}

func (ctx *Context) SetProfile(profile int) {
	ctx.CodeContext().profile = (C.int)(profile)
}

func (ctx *Context) Level() int {
	return int(ctx.CodeContext().level)
}

func (ctx *Context) SetLevel(level int) {
	ctx.CodeContext().level = (C.int)(level)
}
func (ctx *Context) SubtitleHeaderSize() int {
	return int(ctx.CodeContext().subtitle_header_size)
}

func (ctx *Context) SetSubtitleHeaderSize(size int) {
	ctx.CodeContext().subtitle_header_size = (C.int)(size)
}

func (ctx *Context) VBVDelay() uint64 {
	return uint64(ctx.CodeContext().vbv_delay)
}

func (ctx *Context) SetVBVDelay(delay uint64) {
	ctx.CodeContext().vbv_delay = (C.uint64_t)(delay)
}

func (ctx *Context) SideDataOnlyPackets() bool {
	return ctx.CodeContext().side_data_only_packets != C.int(0)
}

func (ctx *Context) SetSideDataOnlyPackets(sideDataOnly bool) {
	var cSideDataOnly C.int
	if sideDataOnly {
		cSideDataOnly = C.int(1)
	}
	ctx.CodeContext().side_data_only_packets = cSideDataOnly
}

func (ctx *Context) InitialPadding() int {
	return int(ctx.CodeContext().initial_padding)
}

func (ctx *Context) SetInitialPadding(initialPadding int) {
	ctx.CodeContext().initial_padding = (C.int)(initialPadding)
}

func (ctx *Context) FrameRate() *avutil.Rational {
	return avutil.NewRationalFromC(unsafe.Pointer(&ctx.CodeContext().framerate))
}

func (ctx *Context) SetFrameRate(frameRate *avutil.Rational) {
	ctx.CodeContext().framerate.num = (C.int)(frameRate.Numerator())
	ctx.CodeContext().framerate.den = (C.int)(frameRate.Denominator())
}

func (ctx *Context) PTSCorrectionLastPTS() int64 {
	return int64(ctx.CodeContext().pts_correction_last_pts)
}

func (ctx *Context) PTSCorrectionLastDTS() int64 {
	return int64(ctx.CodeContext().pts_correction_last_dts)
}

func (ctx *Context) SubtitlesEncoding() (string, bool) {
	return cStringToStringOk(ctx.CodeContext().sub_charenc)
}

func (ctx *Context) SubtitlesEncodingMode() SubtitlesEncodingMode {
	return SubtitlesEncodingMode(ctx.CodeContext().sub_charenc_mode)
}

func (ctx *Context) SetSubtitlesEncodingMode(mode SubtitlesEncodingMode) {
	ctx.CodeContext().sub_charenc_mode = (C.int)(mode)
}

func (ctx *Context) SkipAlpha() bool {
	return ctx.CodeContext().skip_alpha != C.int(0)
}

func (ctx *Context) SetSkipAlpha(skip bool) {
	var cSkip C.int
	if skip {
		cSkip = C.int(1)
	}
	ctx.CodeContext().skip_alpha = cSkip
}

func (ctx *Context) SeekPreRoll() int {
	return int(ctx.CodeContext().seek_preroll)
}

func (ctx *Context) SetSeekPreRoll(seek int) {
	ctx.CodeContext().seek_preroll = (C.int)(seek)
}

func (ctx *Context) CodecWhitelist() []string {
	return cStringSplit(ctx.CodeContext().codec_whitelist, ",")
}

func (ctx *Context) StatsIn() []byte {
	if ctx.CodeContext().stats_in == nil {
		return nil
	}
	length := int(C.strlen(ctx.CodeContext().stats_in))
	return (*[1 << 30]byte)(unsafe.Pointer(ctx.CodeContext().stats_in))[:length:length]
}

func (ctx *Context) SetStatsIn(in []byte) error {
	C.av_freep(unsafe.Pointer(&ctx.CodeContext().stats_in))
	if len(in) == 0 {
		return nil
	}
	length := len(in)
	cIn := (*C.char)(C.av_malloc(C.size_t(length + 1)))
	if cIn == nil {
		return ErrAllocationError
	}
	if len(in) > 0 {
		C.memcpy(unsafe.Pointer(cIn), unsafe.Pointer(&in[0]), C.size_t(length))
	}
	C.memset(unsafe.Pointer(uintptr(unsafe.Pointer(cIn))+uintptr(length)), 0, 1)
	ctx.CodeContext().stats_in = cIn
	return nil
}

func (ctx *Context) StatsOut() []byte {
	if ctx.CodeContext().stats_out == nil {
		return nil
	}
	length := int(C.strlen(ctx.CodeContext().stats_out))
	return (*[1 << 30]byte)(unsafe.Pointer(ctx.CodeContext().stats_out))[:length:length]
}

func (ctx *Context) SetStatsOut(out []byte) error {
	C.av_freep(unsafe.Pointer(&ctx.CodeContext().stats_out))
	if len(out) == 0 {
		return nil
	}
	length := len(out)
	cOut := (*C.char)(C.av_malloc(C.size_t(length + 1)))
	if cOut == nil {
		return ErrAllocationError
	}
	if len(out) > 0 {
		C.memcpy(unsafe.Pointer(cOut), unsafe.Pointer(&out[0]), C.size_t(length))
	}
	C.memset(unsafe.Pointer(uintptr(unsafe.Pointer(cOut))+uintptr(length)), 0, 1)
	ctx.CodeContext().stats_out = cOut
	return nil
}

func cStringToStringOk(cStr *C.char) (string, bool) {
	if cStr == nil {
		return "", false
	}
	return C.GoString(cStr), true
}

func cStringSplit(cStr *C.char, sep string) []string {
	str, ok := cStringToStringOk(cStr)
	if !ok {
		return nil
	}
	return strings.Split(str, sep)
}

type CodecDescriptor struct {
	CAVCodecDescriptor *C.AVCodecDescriptor
}

func NewCodecDescriptorFromC(cCodec unsafe.Pointer) *CodecDescriptor {
	return &CodecDescriptor{CAVCodecDescriptor: (*C.AVCodecDescriptor)(cCodec)}
}

func (c *CodecDescriptor) ID() CodecID {
	return CodecID(c.CAVCodecDescriptor.id)
}

func (c *CodecDescriptor) CodecType() avutil.MediaType {
	return (avutil.MediaType)(c.CAVCodecDescriptor._type)
}

func (c *CodecDescriptor) Name() string {
	str, _ := c.NameOk()
	return str
}

func (c *CodecDescriptor) NameOk() (string, bool) {
	return cStringToStringOk(c.CAVCodecDescriptor.name)
}

func (c *CodecDescriptor) LongName() string {
	str, _ := c.LongNameOk()
	return str
}

func (c *CodecDescriptor) LongNameOk() (string, bool) {
	return cStringToStringOk(c.CAVCodecDescriptor.long_name)
}

func (c *CodecDescriptor) Props() CodecProps {
	return CodecProps(c.CAVCodecDescriptor.props)
}

func (c *CodecDescriptor) MimeTypes() []string {
	if c.CAVCodecDescriptor.mime_types == nil {
		return nil
	}
	var mimeTypes []string
	for i := 0; ; i++ {
		mimeType := C.get_list_at(c.CAVCodecDescriptor.mime_types, C.int(i))
		if mimeType == nil {
			break
		}
		mimeTypes = append(mimeTypes, C.GoString(mimeType))
	}
	return mimeTypes
}

func (c *CodecDescriptor) Profiles() []*Profile {
	var profiles []*Profile
	for i := 0; ; i++ {
		cProfile := C.go_av_profile_get(c.CAVCodecDescriptor.profiles, C.int(i))
		if cProfile == nil || int(cProfile.profile) == ProfileUnknown {
			break
		}
		profile := NewProfileFromC(unsafe.Pointer(cProfile))
		profiles = append(profiles, profile)
	}
	return profiles
}

func CodecDescriptorByID(codecID CodecID) *CodecDescriptor {
	cCodecDescriptor := C.avcodec_descriptor_get((C.enum_AVCodecID)(codecID))
	if cCodecDescriptor == nil {
		return nil
	}
	return NewCodecDescriptorFromC(unsafe.Pointer(cCodecDescriptor))
}

func CodecDescriptorByName(name string) *CodecDescriptor {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cCodecDescriptor := C.avcodec_descriptor_get_by_name(cName)
	if cCodecDescriptor == nil {
		return nil
	}
	return NewCodecDescriptorFromC(unsafe.Pointer(cCodecDescriptor))
}

func CodecDescriptors() []*CodecDescriptor {
	var prev *C.AVCodecDescriptor
	var descriptors []*CodecDescriptor
	for {
		prev = C.avcodec_descriptor_next(prev)
		if prev == nil {
			break
		}
		descriptors = append(descriptors, NewCodecDescriptorFromC(unsafe.Pointer(prev)))
	}
	return descriptors
}

func FindBestPixelFormat(list []avutil.PixelFormat, src avutil.PixelFormat, alpha bool) avutil.PixelFormat {
	best := findBestPixelFormatWithLossFlags(list, src, alpha, nil)
	return best
}

func FindBestPixelFormatWithLossFlags(list []avutil.PixelFormat, src avutil.PixelFormat, alpha bool, lossFlags avutil.LossFlags) (avutil.PixelFormat, avutil.LossFlags) {
	best := findBestPixelFormatWithLossFlags(list, src, alpha, &lossFlags)
	return best, lossFlags
}

func findBestPixelFormatWithLossFlags(list []avutil.PixelFormat, src avutil.PixelFormat, alpha bool, lossFlags *avutil.LossFlags) avutil.PixelFormat {
	size := len(list)
	value := make([]C.enum_AVPixelFormat, size+1, size+1)
	for i := 0; i < size; i++ {
		value[i] = C.enum_AVPixelFormat(list[i])
	}
	value[size] = C.enum_AVPixelFormat(avutil.PixelFormatNone)
	cList := (*C.enum_AVPixelFormat)(unsafe.Pointer(&value[0]))
	cSrc := (C.enum_AVPixelFormat)(src)
	cAlpha := boolToCInt(alpha)
	var cLossFlags *C.int
	if lossFlags != nil {
		cLossFlagsVal := (C.int)(*lossFlags)
		cLossFlags = &cLossFlagsVal
	}
	best := C.avcodec_find_best_pix_fmt_of_list(cList, cSrc, cAlpha, cLossFlags)
	if lossFlags != nil {
		*lossFlags = avutil.LossFlags(*cLossFlags)
	}
	return avutil.PixelFormat(best)
}

func boolToCInt(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
