
# Zero values

> Variables declared without an explicit initial value are given their zero value.

变量声明时没有指定初始值那么将被会被初始化为类型对应的零值。

> The zero value is:

类型的零值有：

```
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
```

## Supplement

(补充各种类型的零值）
 
- 0 for all integer types,
- 0.0 for floating point numbers,
- false for booleans,
- "" for strings,
- nil for interfaces, slices, channels, maps, pointers and functions.

> The elements of an array or struct will have its fields zeroed if no value is specified. 

结构体数组中的元素如果没有被初始化，那么每个结构体中字段都将被赋予类型对应的零值

```
type T struct {
    n int
    f float64
    next *T
}
fmt.Println([2]T{}) // [{0 0 <nil>} {0 0 <nil>}]
```
