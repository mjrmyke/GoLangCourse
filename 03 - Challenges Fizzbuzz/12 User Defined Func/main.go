//Myke Walker

package main

import (
	"fmt"
)

func foo(x ...int) int {
	var numba int
	for _, i := range x {
		numba += i
	}
	return numba
}

func main() {
	fmt.Println(foo(1, 2))
	fmt.Println(foo(1, 2, 3))
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))
	fmt.Println(foo())
}
