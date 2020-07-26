package converter

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

type (
	//StrInt convert string to (int, int8, int16, int32, int64)
	StrInt struct {
	}
)

func NewStrInt() *StrInt {
	return &StrInt{}
}

func (si *StrInt) Convert(src, dst interface{}) error {
	source, ok := src.(string)
	if !ok {
		return ErrBadInputType
	}

	var bitSize int
	switch dst.(type) {
	case *int8:
		bitSize = 8
	case *int16:
		bitSize = 16
	case *int32:
		bitSize = 32
	case *int64:
		bitSize = 64
	case *int:
		bitSize = 0
	default:
		return ErrBadDstType
	}

	ic, err := strconv.ParseInt(source, 10, bitSize)
	if err != nil {
		return errors.WithMessagef(err, "parse '%s' to int fail", source)
	}
	val := reflect.ValueOf(dst).Elem()
	val.SetInt(ic)
	return nil
}
