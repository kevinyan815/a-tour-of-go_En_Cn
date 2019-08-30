# Exported names

> In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package.

在Go中，以大写字母开头的名称是被导出的名称。比如说`Pizza`是一个导出名称，`Pi`也一样，它是`math`包的导出的。

> pizza and pi do not start with a capital letter, so they are not exported.

`pizza`和`pi`，没有以大写字母开头，所以他们是未导出的。

> When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

当导入一个包到程序中，只能引用包中的导出名称。任何“未导出”名称在包外都是不可访问的。

> pizza and pi do not start with a capital letter, so they are not exported.

`pizza`和`pi`不是以大写字母开头，所以他们是未导出的。

> Run the code. Notice the error message.

> To fix the error, rename math.pi to math.Pi and try it again.

运行下面的代码，注意返回的错误信息。要修复这个错误，把`math.pi`改名成`math.Pi`再试一遍。

```
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
}
```

[SourceURL](https://tour.golang.org/basics/3)
