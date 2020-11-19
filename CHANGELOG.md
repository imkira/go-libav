# CHANGELOG

# [ Unreleased ]

- Change avformat.Context.ControlMessage to return (int64, error) so we can return data from the control message
- Expose the AV_NO_PTS_VALUE in avutil
- Add ability to unset the 'interrupt' flag for a fmt context so you can stop it returning AVERROR_EXIT
- Change avfilter.Context.SendCommand to return (string, error) so we can read the return data. User now sets the expected length and go allocates a buffer
- Adds a packet.Copy method
- Increment version to compile against 4.3
- Add `avcodec.Packet.Write`, `avcodec.Packet.WriteBytes`, `avcodec.Packet.GetDataAt` methods
- Add `avcodec.IOContext.Error` method