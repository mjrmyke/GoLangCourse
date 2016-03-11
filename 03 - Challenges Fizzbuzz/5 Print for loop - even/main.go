//Myke Walker

package main

import (
	"fmt"
)

func main() {
	var count = 100
	for i := 1; i < count; i++ {
		if i%2 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
