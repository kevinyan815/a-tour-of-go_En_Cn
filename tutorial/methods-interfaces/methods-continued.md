# Methods continued

> You can declare a method on non-struct types, too.

你也可以在一个非结构体类型上定义方法

> In this example we see a numeric type MyFloat with an Abs method.

在下面这个例子中我们可以看到一个数字类型`MyFloat`和它的方法`Abs`。

> You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

你只能为包中定义了类型的接收者声明方法。不能
