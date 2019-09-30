# Exercise: Fibonacci closure

> Let's have some fun with functions.

让我们用函数做些好玩的事情。

> Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。

```
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

[SourceURL](https://tour.golang.org/moretypes/26)

[Solution](https://github.com/kevinyan815/a-tour-of-go_En_Cn/blob/master/tutorial/basics/struct-slice-map/exercises/exercise-fibonacci-closure.go)
