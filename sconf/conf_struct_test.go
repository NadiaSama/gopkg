package sconf

import (
	"testing"

	"github.com/pkg/errors"
)

type (
	s1 struct {
		Name string
		Age  int
		Val  float64
	}

	s2 struct {
		Name string
		Age  int
		Val  float64
	}

	s2Validate struct {
	}
)

func (sv *s2Validate) Validate(v interface{}) error {
	val := v.(*s2)
	if val.Age > 100 {
		return errors.New("validate fail")
	}
	return nil
}

func TestParse(t *testing.T) {
	ps2 := &s2{Age: 23, Name: "hehe"}
	parsed, _ := buildConfStruct(ps2, NoValidate)
	if parsed.meta["Name"].Idx != 0 || parsed.meta["Name"].Type.String() != "string" ||
		parsed.meta["Age"].Idx != 1 || parsed.meta["Age"].Type.String() != "int" ||
		parsed.meta["Val"].Idx != 2 || parsed.meta["Val"].Type.String() != "float64" {
		t.Errorf("parse fail %v", parsed)
	}
}

func TestLoad(t *testing.T) {
	ps2 := &s2{Age: 23, Name: "hehe"}
	parsed, _ := buildConfStruct(ps2, NoValidate)

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
	parsed, _ := buildConfStruct(ps2, NoValidate)

	if err := parsed.update(map[string]interface{}{"Val": 11.0, "Name": "haha", "key": 1}); err == nil {
		t.Errorf("update bad data failed")
	}
	if err := parsed.update(map[string]interface{}{"Val": 23, "Name": 1.0}); err == nil {
		t.Errorf("update bad fmt failed")
	}
	if err := parsed.update(map[string]interface{}{"Val": 11.0, "Name": "new", "Age": 12}); err != nil {
		t.Errorf("update failed %v", err.Error())
	}

	var d2 s2
	parsed.load(&d2)
	if d2.Name != "new" || d2.Age != 12 || d2.Val != 11.0 {
		t.Errorf("update failed %v", d2)
	}
	parsed.update(map[string]interface{}{"Val": 23.0, "Name": "u2"})
	parsed.load(&d2)
	if d2.Name != "u2" || d2.Age != 12 || d2.Val != 23.0 {
		t.Errorf("update failed %v", d2)
	}

}

func TestValidate(t *testing.T) {
	ps2 := &s2{Age: 101, Name: "hehe", Val: 23.3}
	v2 := &s2Validate{}
	if _, err := buildConfStruct(ps2, v2); err == nil {
		t.Errorf("test init validate fail")
	}
	ps2.Age = 99
	parsed, _ := buildConfStruct(ps2, v2)

	if err := parsed.update(map[string]interface{}{"Age": 101}); err == nil {
		t.Error("test update validate fail")
	}
	var d2 s2
	parsed.load(&d2)
	if d2.Age != 99 {
		t.Errorf("test load fail %v", d2)
	}
	parsed.update(map[string]interface{}{"Age": 98})
	parsed.load(&d2)
	if d2.Age != 98 {
		t.Errorf("test load fail %v", d2)
	}
}

func TestStructToMap(t *testing.T) {
	s := s1{
		Name: "hehe",
		Age:  23,
		Val:  12.0,
	}

	m := structToMap(s)
	if m["Name"].(string) != "hehe" || m["Age"].(int) != 23 || m["Val"].(float64) != 12.0 {
		t.Errorf("bad map %v", m)
	}
}
