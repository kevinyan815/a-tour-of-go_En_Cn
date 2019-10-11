# Errors

> Go programs express error state with error values.

Go程序通过`error`值表示错误的状态。

> The error type is a built-in interface similar to fmt.Stringer:

`error`类型是一个类似`fmt.Stringer`的内建接口类型。
```
type error interface {
    Error() string
}
```

> (As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

（与 fmt.Stringer 类似，fmt 包在打印值时也会查找error接口。）

> Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

函数经常返回一个`error`值，调用函数的代码也应当通过测试`error`值是否是`nil`来进行错误处理。

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

> A nil error denotes success; a non-nil error denotes failure.

error 为 nil 时表示成功；非 nil 的 error 表示失败。

### Code Example

```
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

[SourceURL](https://tour.golang.org/methods/19)
