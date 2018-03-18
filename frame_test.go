package astiffprobe

import (
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
)

var (
	mockedFrames = []byte(`{
		"frames": [
			{
				"media_type": "video",
				"stream_index": 0,
				"key_frame": 0,
				"pkt_pts": 424080,
				"pkt_pts_time": "4.712000",
				"pkt_dts": 424080,
				"pkt_dts_time": "4.712000",
				"best_effort_timestamp": 424080,
				"best_effort_timestamp_time": "4.712000",
				"pkt_duration": 3600,
				"pkt_duration_time": "0.040000",
				"pkt_pos": "10395880",
				"pkt_size": "52442",
				"width": 1920,
				"height": 1080,
				"pix_fmt": "yuv420p",
				"sample_aspect_ratio": "1:1",
				"pict_type": "B",
				"coded_picture_number": 114,
				"display_picture_number": 0,
				"interlaced_frame": 1,
				"top_field_first": 1,
				"repeat_pict": 0
			},
			{
				"media_type": "video",
				"stream_index": 0,
				"key_frame": 1,
				"pkt_pts": 427680,
				"pkt_pts_time": "4.752000",
				"pkt_dts": 427680,
				"pkt_dts_time": "4.752000",
				"best_effort_timestamp": 427680,
				"best_effort_timestamp_time": "4.752000",
				"pkt_duration": 3600,
				"pkt_duration_time": "0.040000",
				"pkt_pos": "10150122",
				"pkt_size": "195561",
				"width": 1920,
				"height": 1080,
				"pix_fmt": "yuv420p",
				"sample_aspect_ratio": "1:1",
				"pict_type": "I",
				"coded_picture_number": 112,
				"display_picture_number": 0,
				"interlaced_frame": 1,
				"top_field_first": 1,
				"repeat_pict": 0
			}
		]
	}`)
	frame1 = Frame{
		BestEffortTimestamp:     424080,
		BestEffortTimestampTime: Duration{Duration: 4712000000},
		CodedPictureNumber:      114,
		DisplayPictureNumber:    0,
		Height:                  1080,
		InterlacedFrame:         true,
		KeyFrame:                false,
		MediaType:               "video",
		PictType:                "B",
		PixFmt:                  "yuv420p",
		PktDuration:             3600,
		PktDurationTime:         Duration{Duration: 40000000},
		PktDts:                  424080,
		PktDtsTime:              Duration{Duration: 4712000000},
		PktPos:                  10395880,
		PktPts:                  424080,
		PktPtsTime:              Duration{Duration: 4712000000},
		PktSize:                 52442,
		RepeatPict:              false,
		SampleAspectRatio:       Ratio{Height: 1, Width: 1},
		StreamIndex:             0,
		ToFieldFirst:            true,
		Width:                   1920,
	}
	frame2 = Frame{
		BestEffortTimestamp:     427680,
		BestEffortTimestampTime: Duration{Duration: 4752000000},
		CodedPictureNumber:      112,
		DisplayPictureNumber:    0,
		Height:                  1080,
		InterlacedFrame:         true,
		KeyFrame:                true,
		MediaType:               "video",
		PictType:                "I",
		PixFmt:                  "yuv420p",
		PktDuration:             3600,
		PktDurationTime:         Duration{Duration: 40000000},
		PktDts:                  427680,
		PktDtsTime:              Duration{Duration: 4752000000},
		PktPos:                  10150122,
		PktPts:                  427680,
		PktPtsTime:              Duration{Duration: 4752000000},
		PktSize:                 195561,
		RepeatPict:              false,
		SampleAspectRatio:       Ratio{Height: 1, Width: 1},
		StreamIndex:             0,
		ToFieldFirst:            true,
		Width:                   1920,
	}
)

func TestFFProbe_Frames(t *testing.T) {
	// Setup
	f, e := setupFFProbe(mockedFrames)

	// Exercise
	fs, err := f.Frames(context.Background(), "/src", 1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, []Frame{frame1, frame2}, fs)
	assert.Equal(t, []string{"/binary", "-loglevel", "error", "-show_frames", "-select_streams", "1", "-print_format", "json", "/src"}, e.args)
}
