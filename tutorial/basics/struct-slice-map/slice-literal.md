# Slice literals

> A slice literal is like an array literal without the length.

切片的字面量就像是没有长度的数组字面量。

> This is an array literal:

这是一个数组字面量

```
[3]bool{true, true, false}
```

> And this creates the same array as above, then builds a slice that references it:

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

```
[]bool{true, true, false}
```

### Code Example

```
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

[SourceURL](https://tour.golang.org/moretypes/10)
