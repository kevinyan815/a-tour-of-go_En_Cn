# Switch evaluation order

> Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

switch的case语句从上到下计算case的值，知道遇到匹配的case为止。

```
(For example,

switch i {
case 0:
case f():
}
does not call f if i==0.)// 当i == 0是不会调用函数f (Go的switch在执行完case块后即退出)
```


> Note: Time in the Go playground always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader.

*注意：* Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始，该值的意义作为作业练习留给读者去发现。

### Code Example

```
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
```
[SourceURL](https://tour.golang.org/flowcontrol/10)


### Supplement

如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。上面的例子修改为

```
switch i {
case 0: fullthrough
case f():
}

// 当i == 0时 会调用函数f
```
