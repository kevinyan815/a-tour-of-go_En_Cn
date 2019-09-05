#Type conversions

> The expression T(v) converts the value v to the type T.

表达式`T(v)`将值v的类型转换为类型T

> Some numeric conversions:

一些数字类型转换：

```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

> Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

不像C语言，在Go中不同类型元素之间的赋值需要显式的类型转换。去掉下面示例代码中`float64`或者`unit`类型转换，运行看看会发生什么。

```
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
```

[Source URL](https://tour.golang.org/basics/13)
