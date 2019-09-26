
# Mutating Maps

> Insert or update an element in map m:

写入或更新一个元素到`map` `m`中
```
m[key] = elem
```
> Retrieve an element:

读取一个元素
```
 elem = m[key]
```
> Delete an element:

删除一个元素
```
delete(m, key)
```
> Test that a key is present with a two-value assignment:

用双值赋值测试一个键是否存在
```
elem, ok = m[key]
```

> If key is in m, ok is true. If not, ok is false.

> If key is not in the map, then elem is the zero value for the map's element type.

如果键存在于`m`中，`ok`的值是`true`否则为`false`。

如果键不存在，则取出的值是map元素类型对应的零值。

> Note: If elem or ok have not yet been declared you could use a short declaration form:

注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：

```
elem, ok := m[key]
```

### Code Example

```
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```

[SourceURL](https://tour.golang.org/moretypes/22)
