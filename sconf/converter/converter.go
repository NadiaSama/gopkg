package converter

import (
	"github.com/pkg/errors"
)

type (
	//Instance do type convert. convert raw input to config struct field
	Instance interface {
		//dst must be pointer of config struct field
		Convert(src, dst interface{}) error
	}
)

var (
	//ErrBadInput means input type is incorrect
	ErrBadInputType = errors.New("bad input type")
	//ErrBadDstType means dst type is incorrect
	ErrBadDstType = errors.New("bad dest type")
)
