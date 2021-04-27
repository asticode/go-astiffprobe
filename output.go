package astiffprobe

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asticode/go-astikit"
)

// Output represents the object FFProbe outputs
// https://ffmpeg.org/doxygen/2.7/structAVFrame.html
type Output struct {
	Format  Format   `json:"format"`
	Frames  []Frame  `json:"frames"`
	Packets []Packet `json:"packets"`
	Streams []Stream `json:"streams"`
}

// Bool represents a boolean in an int format
type Bool bool

// UnmarshalText implements the JSONUnmarshaler interface
// We need to use UnmarshalJSON instead of UnmarshalText otherwise it fails
func (bl *Bool) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "1" {
		*bl = Bool(true)
		return
	}
	*bl = Bool(false)
	return
}

// Rational represents a rational either in "25/1" or "16:9" format
type Rational struct {
	astikit.Rational
}

func newRational(num, den int) Rational {
	return Rational{Rational: *astikit.NewRational(num, den)}
}

// UnmarshalText implements the TextUnmarshaler interface
func (r *Rational) UnmarshalText(b []byte) (err error) {
	sep := "/"
	if strings.Contains(string(b), ":") {
		sep = ":"
	}
	var p = strings.Split(string(b), sep)
	if len(p) == 0 {
		err = fmt.Errorf("astiffprobe: invalid number of args for rational %s", b)
		return
	} else if len(p) == 1 {
		var i int
		if i, err = strconv.Atoi(p[0]); err == nil {
			*r = newRational(i, 1)
		}
	} else {
		var i1, i2 int
		if i1, err = strconv.Atoi(p[0]); err != nil {
			return
		}
		if i2, err = strconv.Atoi(p[1]); err != nil {
			return
		}
		if i1 == 0 || i2 == 0 {
			*r = newRational(0, 0)
		} else {
			*r = newRational(i1, i2)
		}
	}
	return
}

// Duration represents a duration in a string format "1.203" such as the duration is 1.203s
type Duration struct {
	time.Duration
}

// UnmarshalText implements the TextUnmarshaler interface
func (d *Duration) UnmarshalText(b []byte) (err error) {
	var f float64
	if f, err = strconv.ParseFloat(string(b), 64); err != nil {
		return
	}
	*d = Duration{time.Duration(f * 1e9)}
	return
}

// Hexadecimal represents an int in hexadecimal format
type Hexadecimal string

// Hexadecimal implements the TextUnmarshaler interface
func (h *Hexadecimal) UnmarshalText(b []byte) (err error) {
	var n int64
	if n, err = strconv.ParseInt(string(b), 0, 64); err != nil {
		return
	}
	*h = Hexadecimal(strconv.Itoa(int(n)))
	return
}
