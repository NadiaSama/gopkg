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
		Name string  `json:"name"`
		Age  int     `json:"age"`
		Val  float64 `json:"val"`
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

func TestUpdate(t *testing.T) {
	ps2 := &s2{Age: 23, Name: "hehe", Val: 23.3}
	parsed, _ := buildConfStruct(ps2)

	if err := parsed.update(map[string]interface{}{"val": 11.0, "name": "haha", "key": 1}); err == nil {
		t.Errorf("update bad data failed")
	}
	if err := parsed.update(map[string]interface{}{"val": 23, "name": 1.0}); err == nil {
		t.Errorf("update bad fmt failed")
	}
	if err := parsed.update(map[string]interface{}{"val": 11.0, "name": "new", "age": 12}); err != nil {
		t.Errorf("update failed %v", err.Error())
	}

	var d2 s2
	parsed.load(&d2)
	if d2.Name != "new" || d2.Age != 12 || d2.Val != 11.0 {
		t.Errorf("update failed %v", d2)
	}
	parsed.update(map[string]interface{}{"val": 23.0, "name": "u2"})
	parsed.load(&d2)
	if d2.Name != "u2" || d2.Age != 12 || d2.Val != 23.0 {
		t.Errorf("update failed %v", d2)
	}

}
