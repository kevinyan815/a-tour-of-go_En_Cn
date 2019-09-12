
# Defer

> A defer statement defers the execution of a function until the surrounding function returns.

defer语句可以延迟一个函数的执行，直到外层函数返回时才去执行它。

> The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

推迟的函数的参数会被立即计算，但是直到外层函数返回后被推迟的函数才会执行。

```
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

## Stacking defers

> Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

推迟的函数会被放入一个堆栈中。当一个函数返回后，它的推迟函数会按照后进先出的顺序依次执行。

> To learn more about defer statements read this blog post.

更多关于 defer 语句的信息，请阅读[此博文](http://blog.go-zh.org/defer-panic-and-recover)。

```
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
```
