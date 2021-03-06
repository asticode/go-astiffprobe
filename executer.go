package astiffprobe

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// Executer represents an executer
type Executer interface {
	exec(ctx context.Context, args ...string) (*bytes.Buffer, error)
}

type executer struct{}

func (e executer) exec(ctx context.Context, args ...string) (b *bytes.Buffer, err error) {
	// Init
	var cmd = exec.CommandContext(ctx, args[0], args[1:]...)
	var bufOut, bufErr = &bytes.Buffer{}, &bytes.Buffer{}
	cmd.Stdout = bufOut
	cmd.Stderr = bufErr

	// Run cmd
	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("astiffprobe: running %s failed with stderr %s: %w", strings.Join(args, " "), bufErr.Bytes(), err)
		return
	}
	b = bufOut
	return
}
