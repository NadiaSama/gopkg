package jconf

import (
	"testing"

	"github.com/pkg/errors"
)

type (
	s1 struct {
		Name string
		Age  int `json:"age"`
	}

	s2 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

func TestParse(t *testing.T) {
	if _, err := buildConfStruct(nil); err != ErrInvalidConf {
		t.Errorf("test nil parse fail %v", err)
	}

	if _, err := buildConfStruct(&s1{}); errors.Unwrap(err) != ErrInvalidConf {
		t.Errorf("test empty parse fail %v", err)
	}

	ps2 := &s2{Age: 23, Name: "hehe"}
	parsed, _ := buildConfStruct(ps2)
	if parsed.meta["name"].Idx != 0 || parsed.meta["name"].Type.String() != "string" ||
		parsed.meta["age"].Idx != 1 || parsed.meta["age"].Type.String() != "int" {
		t.Errorf("parse fail %v", parsed)
	}
	ps2.Age = 24
	val := parsed.conf.Load().(s2)
	if val.Age != 23 || val.Name != "hehe" {
		t.Errorf("invalid conf %v", val)
	}
}

func TestLoad(t *testing.T) {
	ps2 := &s2{Age: 23, Name: "hehe"}
	parsed, _ := buildConfStruct(ps2)

	var d1 s1
	var d2 s2

	if err := parsed.load(&d1); err == nil {
		t.Errorf("test d1 fail %v", err)
	}
	if err := parsed.load(d2); err == nil {
		t.Errorf("test d2 fail %v", err)
	}
	if err := parsed.load(&d2); err != nil || d2.Age != 23 || d2.Name != "hehe" {
		t.Errorf("test &d2 fail %v", err)
	}
}
