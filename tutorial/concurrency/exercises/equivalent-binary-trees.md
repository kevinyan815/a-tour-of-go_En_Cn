# Exercise: Equivalent Binary Trees

> There can be many different binary trees with the same sequence of values stored at the leaves. For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13.

可以有许多二叉树在节点中存储着相同的值序列，比如说这里的两个二叉树的节点中就存储着序列1, 1, 2, 3, 5, 8, 13。
![binary trees](/tutorial/images/tree.png)

> A function to check whether two binary trees store the same sequence is quite complex in most languages. We'll use Go's concurrency and channels to write a simple solution.

一个检查两个二叉树是否是相等二叉树的函数在大多数编程语言中实现起来非常复杂。我们将使用Go的并行和通道实现一个简单的解决方案。

> This example uses the tree package, which defines the type:

这个练习将会使用到`golang.org/x/tour/tree`包

```
type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}
```


> 1. Implement the Walk function.

1. 实现函数`Walk`


> 2. Test the Walk function.

    The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

    Create a new channel ch and kick off the walker:

    go Walk(tree.New(1), ch)
    Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
    
2. 测试编写的`Walk`函数
   `tree.New(k)`会构造一个结构随机但总是排序好的二叉树，数据节点携带的值为k, 2k, 3k, ......, 10k.
   创建一个新的通道ch并开始调用函数`Walk`:
   
   go Walk(tree.New(1), ch)
   然后从通道读取并打印10个值。他们应该是1, 2, 3 ......, 10.

>3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

3. 实现`Same`函数, `Same`会使用`Walk`函数来判断t1和t2是否存储了相同的值。

> 4. Test the Same function.

    Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

4. 测试`Same`函数

    Same(tree.New(1), tree.New(1)) 应该返回true, Same(tree.New(1), tree.New(2)) 应该返回false.
    

> The documentation for Tree can be found [here](https://godoc.org/golang.org/x/tour/tree#Tree).

Tree包的文档可以在[这里找到](https://godoc.org/golang.org/x/tour/tree#Tree)


### 题解

附加一个解题答案[equivalent-binary-trees.go](https://github.com/kevinyan815/a-tour-of-go_En_Cn/blob/master/tutorial/concurrency/exercises/equivalent-binary-trees.go)，使用了深度优先搜索来中序遍历二叉树（使用深度优先搜索则一定会用到递归）将遍历到的每个节点放到对应的通道中，然后在`Same`函数中比较从两个通道中读取到的值是否都一样。


    
