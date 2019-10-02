# Methods continued

> You can declare a method on non-struct types, too.

你也可以在一个非结构体类型上定义方法

> In this example we see a numeric type MyFloat with an Abs method.

在下面这个例子中我们可以看到一个数字类型`MyFloat`和它的方法`Abs`。

> You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

你只能为包中定义了类型的接收者声明方法。不能为类型定义在其他包中的接收者定义方法（包括比如说int这样的内建类型）。

（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）

### Code Example

```
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```
