package astiffprobe

import (
	"bytes"
	"context"
)

type mockedExecuter struct {
	args []string
	b    *bytes.Buffer
	err  error
}

func (e *mockedExecuter) exec(ctx context.Context, args ...string) (*bytes.Buffer, error) {
	e.args = args
	return e.b, e.err
}

func setupFFProbe(b []byte) (*FFProbe, *mockedExecuter) {
	f := New(Configuration{BinaryPath: "/binary"})
	e := setupMockedExecuter(b)
	f.executer = e
	return f, e
}

func setupMockedExecuter(b []byte) *mockedExecuter {
	return &mockedExecuter{b: bytes.NewBuffer(b)}
}
