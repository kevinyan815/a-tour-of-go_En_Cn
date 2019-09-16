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

## Struct Literals

> A struct literal denotes a newly allocated struct value by listing the values of its fields.

结构体直接量语法通过直接列出字段的值来新分配一个结构体。

> You can list just a subset of fields by using the `Name:` syntax. (And the order of named fields is irrelevant.)

使用`Name: `语法可以仅列出部分字段。（字段名的顺序无关。）

> The special prefix `&` returns a pointer to the struct value.

特殊的前缀`&`返回一个指向结构体的指针。

```
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```
