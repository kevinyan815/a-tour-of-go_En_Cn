# Switch with no condition

> Switch without a condition is the same as switch true.

一个没有条件的switch等同于`switch true`

> This construct can be a clean way to write long if-then-else chains.

这个结构可以作为一种编写长`if-then-else`链的简洁方式。

### Code Example

```
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
```

[SourceURL](https://tour.golang.org/flowcontrol/11)
