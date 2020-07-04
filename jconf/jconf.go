package jconf

import (
	"reflect"
	"sync/atomic"

	"github.com/pkg/errors"
)

type (
	structMeta struct {
		Type reflect.Type
		Idx  int
	}

	confStruct struct {
		meta map[string]structMeta
		conf atomic.Value
	}
)

var (
	ErrInvalidConf = errors.New("invalid config struct")
)

//buildConfStruct build confStruct
func buildConfStruct(s interface{}) (*confStruct, error) {
	value := reflect.Indirect(reflect.ValueOf(s))
	kind := value.Kind()
	if kind != reflect.Struct {
		return nil, ErrInvalidConf
	}

	if !value.CanSet() {
		return nil, errors.WithMessage(ErrInvalidConf, "conf can not set")
	}

	t := value.Type()
	amount := value.NumField()
	ret := &confStruct{
		meta: make(map[string]structMeta, amount),
	}

	for i := 0; i < amount; i++ {
		field := t.Field(i)
		if name, ok := field.Tag.Lookup("json"); !ok {
			return nil, errors.WithMessagef(ErrInvalidConf, "field '%s' missing json tag", field.Name)
		} else {
			ret.meta[name] = structMeta{
				Type: field.Type,
				Idx:  i,
			}
		}
	}

	ret.conf.Store(value.Interface())
	return ret, nil
}

//load conf and store into dst. dst should be pointer type of conf
func (cf *confStruct) load(dst interface{}) error {
	val := cf.conf.Load()
	st := reflect.PtrTo(reflect.TypeOf(val))
	dt := reflect.TypeOf(dst)
	if st != dt {
		return errors.Errorf("unmatch conf type want '%v' got '%v'", st, dt)
	}
	orig := reflect.Indirect(reflect.ValueOf(dst))
	orig.Set(reflect.ValueOf(val))
	return nil
}
