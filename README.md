# About

TODO

# Usage

WARNING: the code below doesn't handle errors for readibility purposes. However you SHOULD!

```go
// Build astiffprobe
var f = astiffprobe.New(astiffprobe.Configuration{BinaryPath: <your binary path>})

// Retrieve format
var fmt astiffprobe.Format
fmt, _ = f.Format(context.Background(), <your input path>)

// Retrieve streams
var streams []astiffprobe.Stream
streams, _ = f.Streams(context.Background(), <your input path>)

// Retrieve frames
var frames []astiffprobe.Frame
frames, _ = f.Frames(context.Background(), <your input path>, <your stream index>)
```