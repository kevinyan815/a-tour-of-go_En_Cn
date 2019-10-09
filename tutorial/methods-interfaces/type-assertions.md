# Type assertions

> A type assertion provides access to an interface value's underlying concrete value.

类型断言提供对一个接口值底层具体值的访问

```
t := i.(T)
```

> This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

这个语句断言接口值`i`保存了一个具体类型`T`的值，并把底层类型为T的值赋给变量`t`

> If i does not hold a T, the statement will trigger a panic.

如果`i`保存的不是一个`T`类型的值，这个语句将会引发`panic`。

> To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

为了测试接口值是否保存了指定的类型值，类型断言语句可以返回两个值：底层类型值和一个报告类型断言是否成功的布尔值。
```
t, ok := i.(T)
```

> If i holds a T, then t will be the underlying value and ok will be true.

如果`i`保存了一个`T`类型的值，`t`的值会是底层值，`ok`的值是`true`。

> If not, ok will be false and t will be the zero value of type T, and no panic occurs.

如果不是，`ok`的值会是`false`，`t`的值将会是类型`T`的零值。而且不会有`panic`发生。

> Note the similarity between this syntax and that of reading from a map.

请注意这种语法和读取一个映射时的相同之处。

### Code Example

```
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

[SourceURL](https://tour.golang.org/methods/15)
