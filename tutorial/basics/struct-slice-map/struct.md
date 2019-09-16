## Structs

> A `struct` is a collection of fields.

结构体是字段的集合。

```
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

## Struct Fields

> Struct fields are accessed using a dot.

结构体的字段通过点号`.`来访问。

```
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

##  Pointers to structs

> Struct fields can be accessed through a struct pointer.

结构体的字段可以通过结构体的指针来访问。

> To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.

要使用结构体的指针`p`访问结构体的字段`X`我们可以这样写`(*p).X`。不过这么写太啰嗦了，Go允许我们直接写`p.X`访问字段X。不需要显示地解引用。
