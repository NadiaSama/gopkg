package converter

import (
	"testing"
	"time"
)

func TestTimeConvert(t *testing.T) {
	conv := NewTime()
	conv = conv.Format("2006-01-02 15:04:05").Precision(PrecisionMsec)
	now, _ := time.Parse("2006-01-02 15:04:05", "2020-03-01 12:12:12")
	msecs := now.Unix() * 1000

	var dst time.Time
	if err := conv.Convert(msecs, &dst); err != nil || !dst.Equal(now) {
		t.Errorf("test secs convert fail %v %v %v", err, dst, now)
	}

	var dst2 time.Time
	str := now.Format(conv.format)
	if err := conv.Convert(str, &dst2); err != nil || !dst2.Equal(now) {
		t.Errorf("test format string convert fail %s %v %v %v", str, err, dst2, now)
	}
}
