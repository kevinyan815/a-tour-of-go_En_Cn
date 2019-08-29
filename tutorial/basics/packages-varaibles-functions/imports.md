
# Imports

> This code groups the imports into a parenthesized, "factored" import statement.

下面的代码用圆括号组合了多个包导入，这是“分组”形式的包导入语句。

```
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

> You can also write multiple import statements, like:

你也可以像下面这样使用多条包导入语句来导入包

```
import "fmt"
import "math"
```

> But it is good style to use the factored import statement.

不过使用分组导入语句是更好的形式。
