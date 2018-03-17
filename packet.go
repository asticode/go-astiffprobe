package astiffprobe

import (
	"context"
	"strconv"

	"sort"

	"github.com/pkg/errors"
)

// Packet represents a packet
type Packet struct {
	CodecType    string   `json:"codec_type"`
	Dts          int      `json:"dts"`
	DtsTime      Duration `json:"dts_time"`
	Duration     int      `json:"duration"`
	DurationTime Duration `json:"duration_time"`
	Flags        string   `json:"flags"`
	Position     int      `json:"position,string"`
	Pts          int      `json:"pts"`
	PtsTime      Duration `json:"pts_time"`
	Size         int      `json:"size,string"`
	StreamIndex  int      `json:"stream_index"`
}

// Packets returns the unordered packets of a stream
func (f *FFProbe) Packets(ctx context.Context, src string, streamIndex int) (ps []Packet, err error) {
	// Execute
	var o Output
	if o, err = f.exec(ctx, f.binaryPath, "-loglevel", "error", "-show_packets", "-select_streams", strconv.Itoa(streamIndex), "-print_format", "json", src); err != nil {
		err = errors.Wrap(err, "astiffprobe: executing failed")
		return
	}
	return o.Packets, nil
}

// PacketsOrdered returns the ordered packets of a stream
func (f *FFProbe) PacketsOrdered(ctx context.Context, src string, streamIndex int) (ps []Packet, err error) {
	// Get packets
	var unorderedPackets []Packet
	if unorderedPackets, err = f.Packets(ctx, src, streamIndex); err != nil {
		err = errors.Wrap(err, "astiffprobe: getting packets failed")
		return
	}

	// Index packets
	var keys []int
	var indexedPackets = make(map[int]Packet)
	for _, p := range unorderedPackets {
		keys = append(keys, p.Pts)
		indexedPackets[p.Pts] = p
	}

	// Order packets
	sort.Ints(keys)
	for _, k := range keys {
		ps = append(ps, indexedPackets[k])
	}
	return
}
