# Interface values

> Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

注：字面上under the hood指的是引擎盖下面的那些平常看不到的构造，主要是引擎、再加上那些乱七八糟的管线等等。所以你应该明白，其实引申的含义就是「表面之下的、内在的、深入的东西」。

在Go内部，接口只可以被看作是一个包含值和具体类型的元组：

> (value, type)

(值, 类型)

> An interface value holds a value of a specific underlying concrete type.

接口值保存了一个具体底层类型的具体值。

> Calling a method on an interface value executes the method of the same name on its underlying type.

接口值调用方法时会执行其底层类型的同名方法。

```
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```
[SourceURL](https://tour.golang.org/methods/11)
