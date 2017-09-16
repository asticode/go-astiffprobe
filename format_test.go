package astiffprobe

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockedFormat = `{
	"filename": "ocscity.ts",
	"nb_streams": 4,
	"nb_programs": 1,
	"format_name": "mpegts",
	"format_long_name": "MPEG-TS (MPEG-2 Transport Stream)",
	"start_time": "1.102122",
	"duration": "298.177978",
	"size": "555115872",
	"bit_rate": "14893544",
	"probe_score": 50
}`

func TestFFProbe_Format(t *testing.T) {
	var f Format
	err := json.Unmarshal([]byte(mockedFormat), &f)
	assert.NoError(t, err)
	assert.Equal(t, Format{Bitrate: 14893544, Duration: Duration{Duration: 298177978000}, Filename: "ocscity.ts", FormatName: "mpegts", FormatNameLong: "MPEG-TS (MPEG-2 Transport Stream)", NumberOfPrograms: 1, NumberOfStreams: 4, ProbeScore: 50, Size: 555115872, StartTime: Duration{Duration: 1102122000}}, f)
}
