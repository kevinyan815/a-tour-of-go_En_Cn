# Interfaces

> An interface type is defined as a set of method signatures.

一个接口类型被定义为一系列方法签名的集合。

> A value of interface type can hold any value that implements those methods.

一个接口类型的变量可以保存任何实现了这些方法的值。

> Note: There is an error in the example code on line 22. Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).

注意：示例代码的22行有一个错误。`Vertex`(值类型)没有实现`Abser`接口的方法因为`Abs`方法值定义在了`*Vertex`(指针类型)上。

### Code Example

```
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}
```

[SourceURL](https://tour.golang.org/methods/9)

