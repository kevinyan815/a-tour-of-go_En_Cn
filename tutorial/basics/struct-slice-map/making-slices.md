# Creating a slice with make

> Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

切片可以由内建函数`make`来创建; 这也是你创建变长数组的方式。

> The make function allocates a zeroed array and returns a slice that refers to that array:

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

```
a := make([]int, 5)  // len(a)=5
```

> To specify a capacity, pass a third argument to make:

传递第三个参数给`make`来指定切片的容量。

```
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

### Code Example

```
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}


// a len=5 cap=5 [0 0 0 0 0]
// b len=0 cap=5 []
// c len=2 cap=5 [0 0]
// d len=3 cap=3 [0 0 0]
```

[SourceURL](https://tour.golang.org/moretypes/13)
