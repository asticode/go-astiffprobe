package astiffprobe

import (
	"context"
	"reflect"
	"testing"
)

var (
	mockedFormat = []byte(`{
		"format": {
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
		}
	}`)
	format = Format{
		Bitrate:          14893544,
		Duration:         Duration{Duration: 298177978000},
		Filename:         "ocscity.ts",
		FormatName:       "mpegts",
		FormatNameLong:   "MPEG-TS (MPEG-2 Transport Stream)",
		NumberOfPrograms: 1,
		NumberOfStreams:  4,
		ProbeScore:       50,
		Size:             555115872,
		StartTime:        Duration{Duration: 1102122000},
	}
)

func TestFFProbe_Format(t *testing.T) {
	// Setup
	f, e := setupFFProbe(mockedFormat)

	// Exercise
	fmt, err := f.Format(context.Background(), "/src")

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if !reflect.DeepEqual(format, fmt) {
		t.Errorf("expected %+v, got %+v", format, fmt)
	}
	if ex := []string{"/binary", "-loglevel", "error", "-show_format", "-print_format", "json", "/src"}; !reflect.DeepEqual(ex, e.args) {
		t.Errorf("expected %+v, got %+v", ex, e.args)
	}
}
