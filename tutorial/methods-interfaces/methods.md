# Methods

> Go does not have classes. However, you can define methods on types.

Go中没有类。但是你可以在类型上定义方法。

> A method is a function with a special receiver argument.

方法就是一类带有特殊的接收者参数的函数。

> The receiver appears in its own argument list between the func keyword and the method name.

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

> In this example, the Abs method has a receiver of type Vertex named v.

在这个例子中，方法`Abs`拥有一个名为`v`类型为`Vertex`的接收者。

```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

[SourceURL](https://tour.golang.org/methods/1)
