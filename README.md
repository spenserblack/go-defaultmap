# defaultmap

[![CI](https://github.com/spenserblack/go-defaultmap/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/go-defaultmap/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/go-defaultmap/branch/main/graph/badge.svg?token=WFufDgESL3)](https://codecov.io/gh/spenserblack/go-defaultmap)

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
