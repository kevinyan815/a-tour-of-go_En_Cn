# Buffered Channels

> Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

通道可以被缓冲，提供缓冲长度作为`make`函数的第二个实参来初始化带缓冲的通道。

```
ch := make(chan int, 100)
```
> Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

只有缓冲通道被填满后向通道发送值才会被阻塞，读取在通道为空时被阻塞。

> Modify the example to overfill the buffer and see what happens.

```
package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```
