# Packages

> Every Go program is made up of packages.

所有Go程序都由包组成。

> Programs start running in package main.

程序从`main`包中开始运行

> This program is using the packages with import paths "fmt" and "math/rand".

本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。

> By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.

按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。

>Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

> (To see a different number, seed the number generator; see rand.Seed. Time is constant in the playground, so you will need to use something else as the seed.)

*注意：* 此程序的运行环境是固定的，因此 rand.Intn 总是会返回相同的数字。 （要得到不同的数字，需为生成器提供不同的种子数，参见 rand.Seed。 练习场中的时间为常量，因此你需要用其它的值作为种子数。）

译者注：go playground练习场中返回的当前时间永远是Go在2006年诞生的时间。

### Code Example

```
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

[Source URL](https://tour.golang.org/basics/1)
