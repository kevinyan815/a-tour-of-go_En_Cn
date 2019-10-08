# Interface values with nil underlying values

# 底层值为 nil 的接口值

> If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

> In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)

在有些语言中这会触发一个空指针异常，但是在Go中编写可以优雅处理被`nil`接收者调用的方法情况很常见（就像例子中的方法`M`一样）。

> Note that an interface value that holds a nil concrete value is itself non-nil.

注意：一个保存了`nil`具体指的接口值，其本身不是`nil`。

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

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// output:
// (<nil>, *main.T)
// <nil>
// (&{hello}, *main.T)
// hello
```

[SourceURL](https://tour.golang.org/methods/12)
