# Interfaces are implemented implicitly

> A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

一个类型通过实现接口所有的方法来实现该接口。即无专门的显式声明也没有"implements"关键字。

> Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

隐式接口从接口实现中解耦了接口定义，这样接口的实现可以出现在任何包中，无需提前准备。

### Code Example

```
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

```

[SourceURL](https://tour.golang.org/methods/10)
