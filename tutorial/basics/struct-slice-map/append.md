# Appending to a slice

> It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

向切片中追加新元素很常见，所以Go提供了内建的`append`函数。 内建包的[文档](https://golang.org/pkg/builtin/#append)中这样描述`append`

```
func append(s []T, vs ...T) []T
```

> The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

(To learn more about slices, read the Slices: usage and internals article.)
