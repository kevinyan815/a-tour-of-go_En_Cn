# sync.Mutex

> We've seen how channels are great for communication among goroutines.

我们已经看到通道在作用于goroutines之间的通信时有多么的棒了。

> But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

但是如果我们不需要goroutine间的相互通信？如果我们只是想确保在同一时间只有一个goroutine能够访问一个变量。

> This concept is called mutual exclusion, and the conventional name for the data structure that provides it is mutex.

这个概念叫做互斥，对于提供互斥的数据结构，他们约定俗称的名字是`mutex`。

> Go's standard library provides mutual exclusion with sync.Mutex and its two methods:

    Lock
    Unlock
  
Go的标准库里通过`sync.Mutex`和它的两个方法`Lock` `Unlock`提供了mutual exclusion支持


> We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock as shown on the Inc method.

> We can also use defer to ensure the mutex will be unlocked as in the Value method.

我们可以定义一个在
