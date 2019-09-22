# Nil slices

> The zero value of a slice is nil.

切片的零值是`nil`。

> A nil slice has a length and capacity of 0 and has no underlying array.

一个`nil`切片的长度和容量是0，并且没有指向任何底层数组。

### Code Example

```
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

[SourceURL](https://tour.golang.org/moretypes/12)
