package astiffprobe

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// FFProbe represents an entity capable of running an FFProbe binary
type FFProbe struct {
	binaryPath string
}

// New creates a new FFProbe
func New(c Configuration) *FFProbe {
	return &FFProbe{binaryPath: c.BinaryPath}
}

var execOutput = func(ctx context.Context, args ...string) (b *bytes.Buffer, err error) {
	// Init
	var cmd = exec.CommandContext(ctx, args[0], args[1:]...)
	var bufOut, bufErr = &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr

	// Run cmd
	if err = cmd.Run(); err != nil {
		err = errors.Wrapf(err, "astiffprobe: running %s failed with stderr %s", strings.Join(args, " "), bufErr.Bytes())
		return
	}
	b = bufOut
	return
}

func (f *FFProbe) exec(ctx context.Context, args ...string) (o Output, err error) {
	// Get output
	var b *bytes.Buffer
	if b, err = execOutput(ctx, args...); err != nil {
		err = errors.Wrap(err, "astiffprobe: getting output failed")
		return
	}

	// Unmarshal
	if err = json.NewDecoder(b).Decode(&o); err != nil {
		err = errors.Wrapf(err, "astiffprobe: unmarshaling %s failed", b.Bytes())
		return
	}
	return
}
