package converter

import "testing"

func TestStrInt(t *testing.T) {
	conv := NewStrInt()
	if err := conv.Convert(1, nil); err != ErrBadInputType {
		t.Errorf("test bad input type fail")
	}
	if err := conv.Convert("1", "2"); err != ErrBadDstType {
		t.Errorf("test bad dst type fail")
	}

	var dst int
	if err := conv.Convert("1a", &dst); err == nil {
		t.Errorf("test bad input data fail")
	}

	if err := conv.Convert("1234", &dst); err != nil || dst != 1234 {
		t.Errorf("bad convert value %d", dst)
	}

}
