# Exercise: Errors

> Copy your Sqrt function from the earlier exercise and modify it to return an error value.

从以前的练习中拷贝`Sqrt`函数过来，修改它使它返回一个`error`值。

> Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

传递复数给Sqrt时，它应该返回一个非nil错误值，因为它不支持复数。

> Create a new type

创建一个新类型
```
type ErrNegativeSqrt float64
```
> and make it an error by giving it a method 

实现下面的方法让它变成一个error类型
```
func (e ErrNegativeSqrt) Error() string
```

> such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

这样，`ErrNegativeSqrt(-2).Error()`会返回 "cannot Sqrt negative number: -2"。

> Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

在`Error`方法中调用`fmt.Sprint(e)将会导致程序进入无限循环。你可以通过将`e`先装换类型`fmt.Sprint(float64(e))`来避免这种情况。为什么？

> Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.

修改Sqrt函数，使传给它一个负数时，返回一个`ErrNegativeSqrt`值。

```
package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

[SolutionURL](https://github.com/kevinyan815/a-tour-of-go_En_Cn/blob/master/tutorial/methods-interfaces/exercises/exercise-errors.go)

[SourceURL](https://tour.golang.org/methods/20)


###  为什么会出现无限循环


fmt.Sprint(e)将调用e.Error()将e转换为字符串。如果Error()方法调用fmt.Sprint(e)，则程序将递归直到内存溢出。可以通过将`e`转换成一个非错误类型(未实现Error接口）的值来避免这种情况。

实际上在Error方法中用error值直接调用`fmt`包中Print相关的函数都会导致无限循环。原因可以在fmt包的源码中找到。
```
		switch verb {
		case 'v', 's', 'x', 'X', 'q':
			// Is it an error or Stringer?
			// The duplication in the bodies is necessary:
			// setting wasString and handled, and deferring catchPanic,
			// must happen before calling the method.
			switch v := p.field.(type) {
			case error:
				wasString = false
				handled = true
				defer p.catchPanic(p.field, verb)
				p.printField(v.Error(), verb, plus, false, depth)
				return
```

fmt源码链接: https://github.com/golang/go/blob/2ed57a8cd86cec36b8370fb16d450e5a29a9375f/src/pkg/fmt/print.go#L639


