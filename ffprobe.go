package astiffprobe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// FFProbe represents an entity capable of running an FFProbe binary
type FFProbe struct {
	binaryPath string
	executer   Executer
}

// New creates a new FFProbe
func New(c Configuration) *FFProbe {
	return &FFProbe{
		binaryPath: c.BinaryPath,
		executer:   executer{},
	}
}

func (f *FFProbe) exec(ctx context.Context, args ...string) (o Output, err error) {
	// Get output
	var b *bytes.Buffer
	if b, err = f.executer.exec(ctx, args...); err != nil {
		err = fmt.Errorf("astiffprobe: executing failed: %w", err)
		return
	}

	// Unmarshal
	if err = json.NewDecoder(b).Decode(&o); err != nil {
		err = fmt.Errorf("astiffprobe: unmarshaling %s failed: %w", b.Bytes(), err)
		return
	}
	return
}
