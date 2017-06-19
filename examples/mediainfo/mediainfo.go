// Lists the streams and some codec details of a media file
//
// Tested with
//
// $ go run examples/mediainfo/medianfo.go --input=https://bintray.com/imkira/go-libav/download_file?file_path=sample_iPod.m4v
//
// stream 0: eng aac audio, 2 channels, 44100 Hz
// stream 1: h264 video, 320x240

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/imkira/go-libav/avcodec"
	"github.com/imkira/go-libav/avformat"
	"github.com/imkira/go-libav/avutil"
)

var inputFileName string

func init() {
	flag.StringVar(&inputFileName, "input", "", "source file to probe")
	flag.Parse()
}

func main() {
	if len(inputFileName) == 0 {
		log.Fatalf("Missing --input=file\n")
	}

	// open format (container) context
	decFmt, err := avformat.NewContextForInput()
	if err != nil {
		log.Fatalf("Failed to open input context: %v", err)
	}

	// set some options for opening file
	options := avutil.NewDictionary()
	defer options.Free()

	// open file for decoding
	if err := decFmt.OpenInput(inputFileName, nil, options); err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer decFmt.CloseInput()

	// initialize context with stream information
	if err := decFmt.FindStreamInfo(nil); err != nil {
		log.Fatalf("Failed to find stream info: %v", err)
	}

	// show stream info
	for _, stream := range decFmt.Streams() {
		language := stream.MetaData().Get("language")
		streamCtx := stream.CodecContext()
		codecID := streamCtx.CodecID()
		descriptor := avcodec.CodecDescriptorByID(codecID)
		switch streamCtx.CodecType() {
		case avutil.MediaTypeVideo:
			width := streamCtx.Width()
			height := streamCtx.Height()
			fmt.Printf("stream %d: %s video, %dx%d\n",
				stream.Index(),
				descriptor.Name(),
				width,
				height)
		case avutil.MediaTypeAudio:
			channels := streamCtx.Channels()
			sampleRate := streamCtx.SampleRate()
			fmt.Printf("stream %d: %s %s audio, %d channels, %d Hz\n",
				stream.Index(),
				language,
				descriptor.Name(),
				channels,
				sampleRate)
		case avutil.MediaTypeSubtitle:
			fmt.Printf("stream %d: %s %s subtitle\n",
				stream.Index(),
				language,
				descriptor.Name())
		}
	}
}
