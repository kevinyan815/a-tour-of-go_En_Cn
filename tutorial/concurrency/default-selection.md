# Default Selection

> The default case in a `select` is run if no other case is ready.

如果select中列出的所有通道都未就绪，那么select中的default case会被执行。

> Use a default case to try a send or receive without blocking:

```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

### Code Example

```
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

[Source Url](https://tour.golang.org/concurrency/6)

