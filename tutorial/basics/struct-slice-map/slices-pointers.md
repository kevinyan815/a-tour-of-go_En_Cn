# Slices are like references to arrays

> A slice does not store any data, it just describes a section of an underlying array.

一个切片不存储任何数据，它只是描述了底层数组的一段。

> Changing the elements of a slice modifies the corresponding elements of its underlying array.

改变一个切片的元素将会更改它的底层数组中对应的元素。

Other slices that share the same underlying array will see those changes.

其他共享了相同底层数组的切片也将会看到这些改变。

### Code Example

```
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

[SourceURL](https://tour.golang.org/moretypes/8)
