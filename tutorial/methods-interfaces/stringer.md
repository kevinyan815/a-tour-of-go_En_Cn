# Stringers

> One of the most ubiquitous interfaces is Stringer defined by the fmt package.

在Go中`fmt`包中定义的`Stringer`接口是最普遍的接口之一。

```
type Stringer interface {
    String() string
}
```

> A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

`Stringer`是一个可以用字符串描述自己的类型。`fmt`包（还有其他包）会查找此接口来打印值。

### Code Example

```
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

// output:
// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```

[SourceURL](https://tour.golang.org/methods/17)
