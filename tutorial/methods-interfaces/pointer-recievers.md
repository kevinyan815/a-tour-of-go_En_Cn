# Pointer receivers

> You can declare methods with pointer receivers.

你可以为指针接收者声明方法。

> This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

这意味着对于某类型 T，接收者的类型可以用 *T 的语法法。（此外，T 不能是像 *int 这样的指针。）

> For example, the Scale method here is defined on *Vertex.

举例来说，这里的`Scale`方法定义在接收者`*Vertex`上。

> Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

拥有指针接收者的方法可以修改接收者指向的值(就像这里`Scale`做的那样)。因为方法经常是需要修改其接收者的，所以指针接收者要比值接收者更常用一些。

> Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

试着去掉`Scale`方法声明中接收者的`*`, 观察一下程序的行为会怎么改变。

> With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.


使用值接收者，`Scale`方法操作的是源`Vertex`值的一个值拷贝。（接收者参数拥有和函数其他参数一样的行为）`Scale`方法必须拥有一个指针接收者才能改变在`main`函数中定义的`Vertex`类型的值


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

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

[SourceURL](https://tour.golang.org/methods/4)
