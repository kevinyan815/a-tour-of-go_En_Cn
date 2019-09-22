# Slices of slices

> Slices can contain any type, including other slices.

切片的值可以是任意类型，包括其他切片。

### Code Example

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// Output
//X _ X
//O _ X
//_ _ O
```

[SourceURL](https://tour.golang.org/moretypes/14)
