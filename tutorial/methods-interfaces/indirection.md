# Methods and pointer indirection

> Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

经过比较前两节的程序，你可能注意到了，拥有指针参数的函数必须接收一个指针。

```
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

> while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

但是用指针作为接收者的方法在被调用时可以使用一个值或者一个指针来作为接收者。

```
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

> For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

对于表达式`v.Scale(5)`即使`v`是一个值不是指针，带指针接收者的方法也能被直接调用。也就是说，由于`Scale`方法有一个指针接收者，为方便起见，Go 会将语句`v.Scale(5)`解释为`(&v).Scale(5)`。

### Code Example

```
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

[SourceURL](https://tour.golang.org/methods/6)

