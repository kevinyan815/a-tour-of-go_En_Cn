## Pointers

> Go has pointers. A pointer holds the memory address of a value.

Go有指针，一个指针保存着一个值的内存地址。


> The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`。

```
var p *int
```

The `&` operator generates a pointer to its operand.

`&`操作符为它的操作数生成一个指针。

```
i := 42
p = &i
```

> The `*` operator denotes the pointer's underlying value.

`*`操作符取出指针指向地址的值。

```
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

> This is known as "dereferencing" or "indirecting".

这叫做“解引用”或者“重定向”

> Unlike C, Go has no pointer arithmetic.

与 C 不同，Go 没有指针运算。

### Code Example

```
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

[SourceURL](https://tour.golang.org/moretypes/1)
