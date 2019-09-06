# Numeric Constants

> Numeric constants are high-precision values.

数值常量是高精度的值

> An untyped constant takes the type needed by its context.

一个未指明类型的常量由上下文来决定其类型。

>Try printing needInt(Big) too.

>(An int can store at maximum a 64-bit integer, and sometimes less.)

在下面的代码示例中尝试一下输出 needInt(Big) 。

（int 可以存放最大64位的整数，根据平台不同有时会更少。）

```
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
```

[Source URL](https://tour.golang.org/basics/16)
