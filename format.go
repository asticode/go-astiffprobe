package astiffprobe

import (
	"context"
	"fmt"
)

// Format represents a format
type Format struct {
	Bitrate          int      `json:"bit_rate,string"`
	Duration         Duration `json:"duration"`
	Filename         string   `json:"filename"`
	FormatName       string   `json:"format_name"`
	FormatNameLong   string   `json:"format_long_name"`
	NumberOfPrograms int      `json:"nb_programs"`
	NumberOfStreams  int      `json:"nb_streams"`
	ProbeScore       int      `json:"probe_score"`
	Size             int      `json:"size,string"`
	StartTime        Duration `json:"start_time"`
}

// Format returns the format of a video
func (f *FFProbe) Format(ctx context.Context, src string) (ft Format, err error) {
	// Execute
	var o Output
	if o, err = f.exec(ctx, f.binaryPath, "-loglevel", "error", "-show_format", "-print_format", "json", src); err != nil {
		err = fmt.Errorf("astiffprobe: executing failed: %w", err)
		return
	}
	return o.Format, nil
}
