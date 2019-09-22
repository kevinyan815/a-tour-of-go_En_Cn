# Slice defaults

> When slicing, you may omit the high or low bounds to use their defaults instead.

在进行切片时，你可以利用它的默认行为来忽略上下界。

> The default is zero for the low bound and the length of the slice for the high bound.

默认行为是0和切片长度作为切片的上下界。

> For the array

对下面的数组`a`来说

```
var a [10]int
```

> these slice expressions are equivalent:

这些切片表达式是相等的

```
a[0:10]
a[:10]
a[0:]
a[:]
```

### Code Example

```
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

[SourceURL](https://tour.golang.org/moretypes/10)
