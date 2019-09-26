# Exercise: Maps

> Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

实现`WordCount`函数。它预期返回一个记录了字符串`s`中每个单词出现次数的`map`。

You might find strings.Fields helpful.

### Code Example

```
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	wordCounter := make(map[string]int)
	for _, word := range words {
		wordCounter[word] += 1
	}
	return wordCounter
}

func main() {
	wc.Test(WordCount)
}
```

[SourceURL](https://tour.golang.org/moretypes/23)
