# Readers

> The io package specifies the io.Reader interface, which represents the read end of a stream of data.

`io`包中指定的`io.Reader`接口表示了一个流数据的读入端。

> The Go standard library contains many implementations of these interfaces, including files, network connections, compressors, ciphers, and others.

Go标准库中包含了很多这个接口的实现，包括网络连接、压缩、加密等。

> The io.Reader interface has a Read method:

`io.Reader`接口拥有一个`Read`方法。

```
func (T) Read(b []byte) (n int, err error)
```

> Read populates the given byte slice with data and returns the number of bytes populated and an error value. It returns an io.EOF error when the stream ends.

`Read`方法用数据填充给定的字节切片并返回填充的字节数和一个`error`值。读取到文件流的末尾时它会返回一个`io.EOF`错误。

> The example code creates a strings.Reader and consumes its output 8 bytes at a time.

示例代码创建了一个`strings.Reader`并且每次8个字节的节奏消费它的输出。

### Code Example

```
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// output
// n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
// b[:n] = "Hello, R"
// n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
// b[:n] = "eader!"
// n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
// b[:n] = ""
```

[SourceURL](https://tour.golang.org/methods/21)
