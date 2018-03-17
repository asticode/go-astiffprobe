package astiffprobe

import (
	"testing"

	"bytes"
	"context"

	"github.com/stretchr/testify/assert"
)

var mockedPackets = `{
	"packets": [
		{
			"codec_type": "video",
			"stream_index": 0,
			"pts": 10806480,
			"pts_time": "120.072000",
			"dts": 10795680,
			"dts_time": "119.952000",
			"duration": 3600,
			"duration_time": "0.040000",
			"size": "150988",
			"pos": "273529764",
			"flags": "__"
		},
		{
			"codec_type": "video",
			"stream_index": 0,
			"pts": 10799280,
			"pts_time": "119.992000",
			"dts": 10799280,
			"dts_time": "119.992000",
			"duration": 3600,
			"duration_time": "0.040000",
			"size": "79786",
			"pos": "273681328",
			"flags": "__"
		},
		{
			"codec_type": "video",
			"stream_index": 0,
			"pts": 10802880,
			"pts_time": "120.032000",
			"dts": 10802880,
			"dts_time": "120.032000",
			"duration": 3600,
			"duration_time": "0.040000",
			"size": "24454",
			"pos": "273761279",
			"flags": "__"
		}
	]
}`

func TestFFProbe_Packets(t *testing.T) {
	execOutput = func(ctx context.Context, args ...string) (b *bytes.Buffer, err error) {
		return bytes.NewBuffer([]byte(mockedPackets)), nil
	}
	f := New(Configuration{})
	ps, err := f.Packets(context.Background(), "", 0)
	assert.NoError(t, err)
	p1 := Packet{CodecType: "video", Dts: 10795680, DtsTime: Duration{Duration: 119952000000}, Duration: 3600, DurationTime: Duration{Duration: 40000000}, Flags: "__", Position: 0, Pts: 10806480, PtsTime: Duration{Duration: 120072000000}, Size: 150988, StreamIndex: 0}
	p2 := Packet{CodecType: "video", Dts: 10799280, DtsTime: Duration{Duration: 119992000000}, Duration: 3600, DurationTime: Duration{Duration: 40000000}, Flags: "__", Position: 0, Pts: 10799280, PtsTime: Duration{Duration: 119992000000}, Size: 79786, StreamIndex: 0}
	p3 := Packet{CodecType: "video", Dts: 10802880, DtsTime: Duration{Duration: 120032000000}, Duration: 3600, DurationTime: Duration{Duration: 40000000}, Flags: "__", Position: 0, Pts: 10802880, PtsTime: Duration{Duration: 120032000000}, Size: 24454, StreamIndex: 0}
	assert.Equal(t, []Packet{p1, p2, p3}, ps)
	ps, err = f.PacketsOrdered(context.Background(), "", 0)
	assert.NoError(t, err)
	assert.Equal(t, []Packet{p2, p3, p1}, ps)
}
