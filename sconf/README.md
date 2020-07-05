# sconf
struct based configuration

## feature
* load `config` into `golang struct`
* concurrent operate `config` support (Get and Update)
* custom validator

### concurrent
`sconf` use `atomic.Value` to store `config` struct. and to achieve concurrent operate


## example
```go
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

```