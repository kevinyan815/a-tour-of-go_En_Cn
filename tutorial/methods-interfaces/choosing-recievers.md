# Choosing a value or pointer receiver

> There are two reasons to use a pointer receiver.

选择指针作为接收者有两个原因。

> The first is so that the method can modify the value that its receiver points to.

第一个是方法可以修改接收者指向的值。

> The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

第二个是避免每次方法调用时都要发生值拷贝。这在接收者是一个大结构体时能提供很多效率。比如说

> In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

在这个例子当中，`Scale`和`Abs`的接收者是类型`*Vertex`，即使`Abs`方法不需要修改他的接收者。

> In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

总体来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。（我们将在接下来的几节中解释为什么）

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

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

[SourceURL](https://tour.golang.org/methods/8)
