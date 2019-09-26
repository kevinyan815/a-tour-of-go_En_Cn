# Map literals

> Map literals are like struct literals, but the keys are required.

`map`的字面量和结构体的字面量很像，但是`map`的键名是必须提供的。

```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

> If the top-level type is just a type name, you can omit it from the elements of the literal.

若顶级类型只是一个类型名，你可以在字面量的元素中省略它。向下面这样元素值中省略了结构体类型`Vertex`：

```
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

[SourceURL](https://tour.golang.org/moretypes/21)
