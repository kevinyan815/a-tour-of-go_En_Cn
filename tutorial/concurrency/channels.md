# Channels
> Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

通道是一个类型管道，可以通过通道操作符`<-`向通道发送和从通道接收值。
```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
> (The data flows in the direction of the arrow.)

（数据按照管道操作符箭头的方向流动）
> Like maps and slices, channels must be created before use:

与`map`和`slice`一样，必须先用`make`方法创建通道才能使用通道：

```
ch := make(chan int)
```

> By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

默认情况下，通道的读写会被阻塞直到另一端就绪。这允许`goroutine`无需显示加锁或者声明条件变量就能在相互之前同步状态。

> The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.

示例代码将一个切片中的数字分发给两个`goroutine`协程，当两个`goroutine`协程完成计算后，用计算结果来算出最终的结果。

```
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

### Source

[Source URL](https://tour.golang.org/concurrency/2)
