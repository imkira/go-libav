# CHANGELOG

# [ Unreleased ]

- Change avformat.Context.ControlMessage to return (int64, error) so we can return data from the control message
- Expose the AV_NO_PTS_VALUE in avutil
- Add ability to unset the 'interrupt' flag for a fmt context so you can stop it returning AVERROR_EXIT