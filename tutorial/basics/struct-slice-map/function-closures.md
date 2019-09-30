# Function closures

> Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

Go函数可以是一个闭包。闭包是一个函数值，它引用了自身函数体外部的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

> For example, the adder function returns a closure. Each closure is bound to its own sum variable.

例如，函数`adder`返回一个闭包。每个闭包都被绑定在其各自的`sum`变量上。


### Code Example

```
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

[SourceURL](https://tour.golang.org/moretypes/25)
