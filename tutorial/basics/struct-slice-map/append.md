# Appending to a slice

> It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

向切片中追加新元素很常见，所以Go提供了内建的`append`函数。 内建包的[文档](https://golang.org/pkg/builtin/#append)中这样描述`append`

```
func append(s []T, vs ...T) []T
```

> The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

`append`的第一个参数`s`是一个元素类型为T的切片，剩下的参数是要往切片中追加的`T`类型的值。

> The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

`append`函数的返回值是一个切片，该切片包含了源切片中的所有元素以及`append`参数中提供的新值。

> If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

如果切片`s`的底层数组太小，不足以放下要追加进来的值的话，Go将分配一个更大的数组。返回的切片将指向新分配的数组。

(To learn more about slices, read the [Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) article.)

### Code Example

```
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Output
// len=0 cap=0 []
// len=1 cap=2 [0]
// len=2 cap=2 [0 1]
// len=5 cap=8 [0 1 2 3 4]
```

[SourceURL](https://tour.golang.org/moretypes/15)
