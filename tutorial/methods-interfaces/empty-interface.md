# The empty interface

> The interface type that specifies zero methods is known as the empty interface:

指定了另个方法的接口类型被称为空接口。

```
interface{}
```

> An empty interface may hold values of any type. (Every type implements at least zero methods.)

一个空接口可能会保存任何类型的值。(因为每个类型都至少实现了零个方法。)

> Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

空接口被处理未知类型的值的代码使用。比如说，`fmt.Print`方法接收类型为`interface{}`的任意数量的参数。

### Code Example

```
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

[SourceURL](https://tour.golang.org/methods/14)
