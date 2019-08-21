# Range and Close

> A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

发送者可以关闭通道来表示不会再有值被发送给通道。接受者可以通过提供第二个参数给通道接受表达式来测试通道是否已经被关闭了。

```
v, ok := <-ch
```

> ok is false if there are no more values to receive and the channel is closed.

当通道为空并且被关闭是`ok`的值为`false`

> The loop for i := range c receives values from the channel repeatedly until it is closed.

循环`for i := range c`反复地从通道中读取数据，直到用到被关闭循环才退出。

> Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

只有发送者需要关闭通道，接受者永远不需要关闭通道。向一个已关闭通道发送数据将会导致运行时panic。

> Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

通道并不像文件，通常来说你不需要关闭他们。只有在接受者必须被告知通道中不会再有数据传输过来是关闭通道才是必需的，比如说终结`range`循环。

### Code Example

```
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```
译者注：斐波那契数列由0和1开始，之后的斐波那契系数就是由之前的两数相加而得出

[Source URL](https://tour.golang.org/concurrency/4)
