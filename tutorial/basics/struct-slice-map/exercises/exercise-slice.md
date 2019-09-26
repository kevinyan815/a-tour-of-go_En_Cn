# Exercise: Slices

> Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

实现`Pic`函数。它返回一个长度为`dy`的切片，每个切片元素是一个长度为`dx`类型为8位无符号整数的切片。当你运行程序时，程序将整数解释为灰阶值然后展示出图片。

> The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

图像的选择取决于您。有趣的函数包括（x + y）/ 2，x * y和x ^ y。

> (You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(你需要用一个循环在切片[][]unit8中分配每个[]unit8切片)

> (Use uint8(intValue) to convert between types.)

(使用unit8(intValue)来转换类型)

```
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}
```
