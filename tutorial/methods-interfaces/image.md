# Images

> Package image defines the Image interface:

image 包定义了 Image 接口：

```
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
> Note: the Rectangle return value of the Bounds method is actually an image.Rectangle, as the declaration is inside package image.

注意: Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。

> (See the documentation for all the details.)

（请参阅[文档](https://go-zh.org/pkg/image/#Rectangle）了解全部信息。）

> The color.Color and color.Model types are also interfaces, but we'll ignore that by using the predefined implementations color.RGBA and color.RGBAModel. These interfaces and types are specified by the image/color package

color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由 image/color 包定义。
