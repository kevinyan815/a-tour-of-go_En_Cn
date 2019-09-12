# Switch
> A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Switch语句是编写一连串`if - else`语句的简便写法。它会运行与条件表达式具有相同的值的第一个case模块。

> Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

Go的switch与C, C++, Java, Javascript和PHP中的switch很像，不同的是Go只会运行选中的case块，而不是随后的所有case块。实际上，其他语言中则case块结尾处需要编写的`break`语句在Go中是自动提供的。另外一个重要的不同是，Go的switch中case不必是常量，case的取值也不必是整数。


### Code Example
```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

```

[SourceURL](https://tour.golang.org/flowcontrol/9)

## Supplement

相比较 C 和 Java 等其它语言而言，Go 语言中的 switch 结构使用上更加灵活。它接受任意形式的表达式：
```
switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
}
```
变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。前花括号 { 必须和 switch 关键字在同一行。

