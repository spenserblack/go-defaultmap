# defaultmap

A map that supports default values.

## Basic Example

```go
import (
	"fmt"

	"github.com/spenserblack/go-defaultmap"
)

m := defaultmap.NewMap(func() string { return "I'm the default!" })
m.Insert("exists", "Hello, World!")

fmt.Println(m.Get("exists")) // Hello, World!
fmt.Println(m.Get("doesn't exist")) // I'm the default!
fmt.Println(m.GetOr("I'm a one-time default!")) // I'm a one-time default!
```
