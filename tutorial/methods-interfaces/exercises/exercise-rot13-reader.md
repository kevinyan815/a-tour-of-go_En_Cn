# Exercise: rot13Reader

> A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

一个常见的模式是一个`io.Reader`包装另外一个`io.Reader`, 用来以某种方式修改流。

> For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

比如说，`gzip.NewReader`接收一个`io.Reader`(一个压缩数据流)返回一个也实现了`io.Reader`的`*gzip.Reader`(解压数据流)。

> Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

编写一个实现了`io.Reader`并从另一个`io.Reader`中读取数据的`rot13Reader`，通过应用`rot13`代换密码对数据流进行修改。

> The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。

ROT13是一种简易加密方式，通过用当前字母在字母表往后13位的字母替换当前字母的方式加密。

```
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

[SourceUrl](https://tour.golang.org/methods/23)
