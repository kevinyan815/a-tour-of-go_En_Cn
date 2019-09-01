
# Variables

> The var statement declares a list of variables; as in function argument lists, the type is last.

`var` 语句用来声明变量列表；跟函数中的参数列表一样，类型放在最后面。

> A var statement can be at package or function level. We see both in this example.

`var`语句可以存在于包或者函数级的作用域中。在下面的例子中可以看到。

```
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
```

### Variables with initializers

> A var declaration can include initializers, one per variable.

`var`声明变量是可以带初始值，每个变量一个。

> If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

在变量声明时如果有初始值，那么类型是可以省略的，变量类型即初始值的类型。

```
package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
```

### Short variable declarations

> Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

在函数中，短赋值语句`:=`可以用来替代具有隐式类型的`var`语句（带初始值的var语句）

> Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

在函数外，每个语句都要以关键字开头(var, func等)，所以`:=`在函数外是不可用的。

```
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
```
