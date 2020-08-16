# sconf
struct based configuration

## feature
* load `config` into `golang struct`
* concurrent operate `config` support (Get and Update)
* custom validator

### concurrent
`sconf` use `atomic.Value` to store `config` struct and implement concurrent operate


## example
```go

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
	ins, _ := sconf.Add("test", &cs, sconf.NoValidate)
	var d confStruct
	sconf.Get("test", &d)
	fmt.Printf("%v\n", d)

	//update with name
	sconf.Update("test", map[string]interface{}{"Age": 21, "Val": 11.0})
	sconf.Get("test", &d)
	fmt.Printf("%v\n", d)

	u := upConf{
		Name: "new name",
		Val:  2.333,
	}

	//update with instance
	ins.Update(&u)
	ins.Get(&d)
	fmt.Printf("%v\n", d)
}

```