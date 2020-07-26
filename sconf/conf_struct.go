package sconf

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
		//meta hold conf struct field info. used by update method
		meta map[string]structMeta
		//conf hold pointer value of struct. to make value settable
		confPtr   atomic.Value
		validator Validator
	}
)

var (
	ErrInvalidConf = errors.New("invalid config struct")
)

//buildConfStruct build confStruct
func buildConfStruct(s interface{}, val Validator) (*confStruct, error) {
	pv := reflect.ValueOf(s)
	kind := pv.Kind()
	if kind != reflect.Ptr {
		return nil, ErrInvalidConf
	}
	orig := pv.Elem()
	if kind := orig.Kind(); kind != reflect.Struct {
		return nil, ErrInvalidConf
	}
	amount := orig.NumField()
	t := orig.Type()
	ret := &confStruct{
		meta: make(map[string]structMeta, amount),
	}

	for i := 0; i < amount; i++ {
		field := t.Field(i)
		ret.meta[field.Name] = structMeta{
			Type: field.Type,
			Idx:  i,
		}
	}

	copy := copyStructToPtr(orig)
	if err := val.Validate(copy.Interface()); err != nil {
		return nil, err
	}
	ret.confPtr.Store(copy)
	ret.validator = val
	return ret, nil
}

//load conf and store into dst. dst should be pointer type of struct
func (cf *confStruct) load(dst interface{}) error {
	val := cf.confPtr.Load().(reflect.Value)
	st := val.Type()
	dt := reflect.TypeOf(dst)
	if st != dt {
		return errors.Errorf("unmatch conf type want '%v' got '%v'", st, dt)
	}
	target := reflect.Indirect(reflect.ValueOf(dst))
	target.Set(val.Elem())
	return nil
}

func (cf *confStruct) update(update map[string]interface{}) error {
	conf := copyStructToPtr(cf.confPtr.Load().(reflect.Value).Elem())
	elem := conf.Elem()
	for key, val := range update {
		meta, ok := cf.meta[key]
		if !ok {
			return errors.Errorf("unkown update filed '%s'", key)
		}
		if t := reflect.TypeOf(val); !t.AssignableTo(meta.Type) {
			return errors.Errorf("unassignble update field '%s' want '%v' got '%v'",
				key, meta.Type, t)
		}
		field := elem.Field(meta.Idx)
		field.Set(reflect.ValueOf(val))
	}
	if err := cf.validator.Validate(conf.Interface()); err != nil {
		return err
	}
	cf.confPtr.Store(conf)
	return nil
}

func copyStructToPtr(src reflect.Value) reflect.Value {
	ret := reflect.New(src.Type())
	t := ret.Elem()
	amount := src.NumField()
	for i := 0; i < amount; i++ {
		sf := src.Field(i)
		df := t.Field(i)
		df.Set(sf)
	}
	return ret
}
