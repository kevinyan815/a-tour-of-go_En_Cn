# Slice length and capacity

> A slice has both a length and a capacity.

一个切片拥有长度和容量两个属性

> The length of a slice is the number of elements it contains.

切片的长度是切片包含的元素的个数。

> The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

> The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

一个切片的长度和容量可以通过表达式`len(s)`和`cap(s)`获得

> You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

你可以通过再切片给它提供足够的容量来改变一个切片的长度。试着更改下面示例中的切片操作让它扩展到超过其容量的长度，看看会发生什么。

### Code Example

```
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

[SourceURL](https://tour.golang.org/moretypes/11)
