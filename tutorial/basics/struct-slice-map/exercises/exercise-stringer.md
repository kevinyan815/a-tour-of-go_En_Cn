# Exercise: Stringers

> Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

用`IPAddr`类型实现`fmt.Stringer`，打印点号分隔的IP地址。

> For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".

比如说，`IPAddr{1, 2, 3, 4}`应该打印为`"1.2.3.4"`。

```
package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

[SourceURL](https://tour.golang.org/methods/18)

[Solution](https://github.com/kevinyan815/a-tour-of-go_En_Cn/blob/master/tutorial/basics/struct-slice-map/exercises/exercise-stringer.go)
