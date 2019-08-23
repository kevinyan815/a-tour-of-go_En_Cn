# Select

> The select statement lets a goroutine wait on multiple communication operations.

`select`语句可以让一个`goroutine`在多通信通道中等待。

> A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

`select`会一直阻塞直到列出的多个通道中已经有一个可以运行，接下来`select`会去执行对应的case. 如果有多个通道可用`select`会随机选择一个去执行。

### Code Example

```
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

```

[Source URL](https://tour.golang.org/concurrency/5)
