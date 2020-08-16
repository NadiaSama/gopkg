package main

import (
	"fmt"

	"github.com/NadiaSama/gopkg/sconf"
)

type (
	confStruct struct {
		Name string
		Age  int
		Val  float64
	}

	upConf struct {
		Name string
		Val  float64
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

	sconf.Update("test", map[string]interface{}{"Age": 21, "Val": 11.0})
	sconf.Get("test", &d)
	fmt.Printf("%v\n", d)

	u := upConf{
		Name: "new name",
		Val:  2.333,
	}
	sconf.Update("test", &u)
	sconf.Get("test", &d)
	fmt.Printf("%v\n", d)

}
