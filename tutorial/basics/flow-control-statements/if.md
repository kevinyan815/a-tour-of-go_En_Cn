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


## If with a short statement

if 的简短语句

> Like for, the if statement can start with a short statement to execute before the condition.

和`for`语句一样，if语句在条件表达式之前可以执行一个简单的语句。

> Variables declared by the statement are only in scope until the end of the if.

在简短语句中声明的变量，只在存在于当前if的作用域中。

> (Try using v in the last return statement.)

试着在下面函数`pow`的最后一个return语句中使用变量v

```
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

## If and else

> Variables declared inside an if short statement are also available inside any of the else blocks.

在一个if语句的简短语句中声明的变量在其后面跟的任意个else代码块中都是可访问的。

> (Both calls to pow return their results before the call to fmt.Println in main begins.)

（在 main 的 fmt.Println 调用开始前，两次对 pow 的调用均已执行并返回其各自的结果。）

```
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```
