# Type inference (类型推断)

> When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.

当定义变量未明确指定变量的类型时（使用`:=`语法或者`var =`语法时），变量的类型将通过右侧赋给变量的值的类型来推断出变量的类型。

> When the right hand side of the declaration is typed, the new variable is of that same type:

声明变量并初始值时右值是声明了类型的变量，新变量的类型与其相同：

```
var i int
j := i // j is an int 变量j类型也是int
```

> But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128 depending on the precision of the constant:

但是当右边包含未指明类型的数值常量时，新变量的类型可能是int，float，或者complex128，视数值常量的精度而定。

```
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

> Try changing the initial value of v in the example code and observe how its type is affected.

试试修改示例代码中v的初始值然后观察它是如何影响类型的。

```
package main

import "fmt"

func main() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
}
```

[Source URL](https://tour.golang.org/basics/14)
