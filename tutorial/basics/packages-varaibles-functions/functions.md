# Functions

> A function can take zero or more arguments.

一个函数可以接收0个或多个实参。

> In this example, add takes two parameters of type int.

在这个例子中，函数`add`接收两个整型的参数、

> Notice that the type comes after the variable name.

注意类型要跟在变量名的后面。

> (For more about why types look the way they do, see the article on Go's declaration syntax.)

(更多关于为什么类型要这样，可以看文章[Go声明语法](https://blog.golang.org/gos-declaration-syntax))

```
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

> When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

当连续两个或多个函数的已命名形参类型相同时，除最后一个参数的类型声明以外，其它都可以省略。

> In this example, we shortened `x int, y int` to `x, y int`

在本例中，`x int, y int` 被缩写为 `x, y int`

### Code Example

```
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

