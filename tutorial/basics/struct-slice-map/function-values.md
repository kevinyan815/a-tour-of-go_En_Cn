# Function values

> Functions are values too. They can be passed around just like other values.

函数也是值，他们可以像其他值一样被传递。

> Function values may be used as function arguments and return values.

函数值可以作为另外一个函数的参数和返回值。

### Code Example

```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```

[SourceURL](https://tour.golang.org/moretypes/24)
