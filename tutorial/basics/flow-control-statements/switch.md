# Switch
> A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

Switch语句是编写一连串`if - else`语句的简便写法。它会运行与条件表达式具有相同的值的第一个case模块。

> Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

Go的switch与C, C++, Java, Javascript和PHP中的switch很像，不同的是Go只会运行选中的case块，而不是随后的所有case块。实际上，其他语言中则case块结尾处需要编写的`break`语句在Go中是自动提供的。另外一个重要的不同是，Go的switch中case不必是常量，case的取值也不必是整数。

## Supplement
