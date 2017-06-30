// +build ffmpeg30

package avfilter

func (l *Link) Status() int {
	return int(l.CAVFilterLink.status)
}

func (l *Link) FrameCount() int64 {
	return int64(l.CAVFilterLink.frame_count)
}
