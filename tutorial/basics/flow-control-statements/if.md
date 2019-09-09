# If
> Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

Go的if语句跟它的for循环一样，表达式不需要被括号`()`环绕，但是花括号`{}`是必须的。

```
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```
