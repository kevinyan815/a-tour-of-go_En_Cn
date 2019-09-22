# Arrays

> The type [n]T is an array of n values of type T.

类型`[n]T`表示一个具有n个类型为T的值的数组。


> The expression

>var a [10]int
>declares a variable a as an array of ten integers.

表达式

```
var a [10]int
```
声明了一个拥有十个整数的数组变量`a`

> An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

数组的长度是其类型的一部分，所以数组不能改变大小。这看起来是个限制，但是不用担心；Go提供了一种更便利的方式来使用数组。

### Code Example

```
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

[SourceURL](https://tour.golang.org/moretypes/6)
