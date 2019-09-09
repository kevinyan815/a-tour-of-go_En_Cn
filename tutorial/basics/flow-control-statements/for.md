
# For

> Go has only one looping construct, the for loop.

Go只有一种循环结构, for循环。

> The basic for loop has three components separated by semicolons:

for循环有三个部分组成，每一部分由分号分隔。

>the init statement: executed before the first iteration

初始化语句：在第一次迭代之前执行。

>the condition expression: evaluated before every iteration

条件表达式： 在每次迭代前计算求值。

>the post statement: executed at the end of every iteration

后置语句：在每次迭代后执行。

> The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

初始化语句通常来说会是一个短标签变量赋值声明，在这里声明的变量只在当前for语句中可见。

> The loop will stop iterating once the boolean condition evaluates to false.

当条件计算求值为false时，循环将会停止。

> Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

不像C、Java或者JavaScript这些其他的编程语言，for语句的三部分不用括号包括起来，而且`{}`无论什么情况下都是必须的。

```
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

[Source URL](https://tour.golang.org/flowcontrol/1)


> The init and post statements are optional.

初始化语句和后置语句是可选的，比如下面的代码

```
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```


