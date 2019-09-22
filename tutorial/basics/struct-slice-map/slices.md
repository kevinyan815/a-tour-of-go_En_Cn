# Slices
> An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

一个数组长度是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

> The type []T is a slice with elements of type T.

类型`[]T`是元素类型为T的切片。

> A slice is formed by specifying two indices, a low and high bound, separated by a colon:

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

```
a[low : high]
```

> This selects a half-open range which includes the first element, but excludes the last one.

它会选择一个半开放区间，区间包含第一个元素，但是不包含最后一个元素。

> The following expression creates a slice which includes elements 1 through 3 of a:

下面的表达式创建了一个切片包含数组a中的1-3下标对应的元素。

```
a[1:4]
```

### Code Example

```
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```

[SourceURL](https://tour.golang.org/moretypes/7)
