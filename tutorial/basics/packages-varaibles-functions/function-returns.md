# Function Returns

## Multiple results

> A function can return any number of results.

在Go中一个函数可以返回任意多个值。

> The swap function returns two strings.

下面的`swap`函数返回了两个字符串值。

```
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

## Named return values

> Go's return values may be named. If so, they are treated as variables defined at the top of the function.

Go的返回值可以被命名，命名返回值会被当做定义在函数顶部的变量来对待。

> These names should be used to document the meaning of the return values.

命名返回值的名称应当是反映其意义的名称。

> A return statement without arguments returns the named return values. This is known as a "naked" return.

没有实参的return语句将返回命名返回值。This is known as a "naked" return（大家自己体会:)）。

> Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。

```
package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```
