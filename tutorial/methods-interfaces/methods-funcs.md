# Methods are functions

> Remember: a method is just a function with a receiver argument.

记住：方法只是一个有接收者参数的函数。

> Here's Abs written as a regular function with no change in functionality.

这里`Abs`作为一个普通的函数与上节中看到的方法在功能是并没有任何改变。

```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

[SourceURL](https://tour.golang.org/methods/2)
