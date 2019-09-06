# Constants

> Constants are declared like variables, but with the const keyword.

可以像定义变量一样定义常量，不过是使用`const`关键字。

> Constants can be character, string, boolean, or numeric values.

常量可以是字符、字符串、布尔值或者数值。

> Constants cannot be declared using the := syntax.

不能使用`:=`语法定义常量。

```
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
```
