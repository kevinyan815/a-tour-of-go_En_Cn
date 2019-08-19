# Goroutines

>A goroutine is a lightweight thread managed by the Go runtime.

一个`goroutine`是由Go运行时(Go runtime)管理的一个轻量级线程
```
go f(x, y, z)
```
>Starts a new goroutine running

开启一个新的goroutine运行
```
f(x, y, z)
```
>The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

`f`，`x`，`y`，`z`在当前的`goroutine`中进行赋值计算，在新开启的`goroutine`中执行

> Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)

所有的`goroutine`运行在相同的地址空间，所以访问共享内存的方式必须是同步的。`sync`包提供了有用的原函数，尽管在Go中你不太需要他们，因为有其他更好的方式来同步地访问共享内存（看下一节）


### Code Example
```
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```

### Source
[Source URL](https://tour.golang.org/concurrency/1)
