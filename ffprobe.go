package astiffprobe

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
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
		err = errors.Wrap(err, "astiffprobe: executing failed")
		return
	}

	// Unmarshal
	if err = json.NewDecoder(b).Decode(&o); err != nil {
		err = errors.Wrapf(err, "astiffprobe: unmarshaling %s failed", b.Bytes())
		return
	}
	return
}
