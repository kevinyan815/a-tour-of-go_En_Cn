# Pointers and functions

> Here we see the Abs and Scale methods rewritten as functions.

这里我们看一下将`Abs`和`Scale`重写为函数。

> Again, try removing the * from line 16. Can you see why the behavior changes? What else did you need to change for the example to compile?

再一次，试着去掉函数定义里参数前面的`*`。你知道为什么函数的行为会改变吗？你还需要做什么才能让这个例子编译通过？

> (If you're not sure, continue to the next page.)

如果你不确定，继续看下一节。

### Code Example

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

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))
}
```

[SourceURL](https://tour.golang.org/methods/5)
