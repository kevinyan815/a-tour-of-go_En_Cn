# Nil interface values
# nil 接口值
> A nil interface value holds neither value nor concrete type.

一个nil接口值既不保存值也不保存具体类型

> Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

在nil接口上调用方法会产生一个运行时错误，因为在接口内部元组中未包含具体类型来指示调用哪个具体方法。

```
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

上面执行会产生运行时错误

```
(<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xe1d84]

goroutine 1 [running]:
main.main()
	/tmp/sandbox038911927/prog.go:12 +0x84
```

[SourceURL](https://tour.golang.org/methods/13)
