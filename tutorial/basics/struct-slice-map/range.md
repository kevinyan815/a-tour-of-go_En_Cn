# Range

> The range form of the for loop iterates over a slice or map.

`range`形式的`for`循环用来迭代切片或映射。

> When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.

当迭代一个切片时，每次迭代将返回两个值。第一个是索引，第二个是索引对应位置的值的拷贝。

```
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```

# Range continued

> You can skip the index or value by assigning to _.

你可以通过赋值给`_`来跳过索引或者值。

```
for i, _ := range pow
for _, value := range pow
```

> If you only want the index, you can omit the second variable.

如果你只想要索引，你可以省略第二个迭代变量。

```
for i := range pow
```

### Code Example

```
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

[SourceURL](https://tour.golang.org/moretypes/17)
