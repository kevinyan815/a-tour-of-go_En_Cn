# Exercise: Readers

> Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

实现一个`Reader`类型，它产生一个 ASCII 字符 'A' 的无限流。

```
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
```
