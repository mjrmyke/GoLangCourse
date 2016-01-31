//Myke Walker

package main

import (
	"fmt"
)

func main() {
	k := 0
	for i := 1; i < 1000; i++ {

		if i%3 == 0 {
			k = k + i

		} else if i%5 == 0 {
			k = k + i

		}
	}
	fmt.Println(k)
}
