package converter

import (
	"time"

	"github.com/pkg/errors"
)

const (
	PrecisionSec TimePrecision = iota
	PrecisionMsec
)

type (
	TimePrecision int
	//Time hold property which required in convert procedure
	Time struct {
		format    string
		precision TimePrecision
	}
)

func NewTime() *Time {
	return &Time{precision: PrecisionSec}
}

func (t *Time) Format(f string) *Time {
	t.format = f
	return t
}

func (t *Time) Precision(p TimePrecision) *Time {
	t.precision = p
	return t
}

//Convert unix timestamp (secs, millisecs in int64) or format string to time.Time,
func (t *Time) Convert(src, dst interface{}) error {
	val, ok := dst.(*time.Time)
	if !ok {
		return ErrBadDstType
	}

	switch p := src.(type) {
	case int64:
		var sec, nsec int64
		if t.precision == PrecisionSec {
			sec = p
			nsec = 0
		} else {
			sec = p / 1000
			nsec = (p % 1000) * 1e6
		}
		*val = time.Unix(sec, nsec)

	case string:
		r, err := time.Parse(t.format, p)
		if err != nil {
			return errors.WithMessagef(err, "parse time '%s' in format '%s'", p, t.format)
		}
		*val = r

	default:
		return ErrBadInputType
	}
	return nil
}
