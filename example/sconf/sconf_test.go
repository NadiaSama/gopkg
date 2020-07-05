package main

import (
	"fmt"

	"github.com/NadiaSama/gopkg/sconf"
)

type (
	confStruct struct {
		Name string `sconf:"name"`
		Age  int
		Val  float64 `sconf:"val"`
	}
)

func main() {
	cs := confStruct{
		Name: "test",
		Age:  23,
		Val:  120.0,
	}
	sconf.Add("test", &cs, sconf.NoValidate)
	var d confStruct
	sconf.Get("test", &d)
	fmt.Printf("%v\n", d)

	sconf.Update("test", map[string]interface{}{"Age": 21, "val": 11.0})
	sconf.Get("test", &d)
	fmt.Printf("%v\n", d)
}
