//https://projecteuler.net/problem=9
//Myke Walker

package main

import (
	"fmt"
)

func foo() int {
	max := 1000
	for a := 1; a < 1001; a++ {
		for b := a + 1; b < 1001; b++ {
			c := max - a - b
			if a*a+b+b == c*c {
				return (a * b * c) //returns 250800
			}

		}
	}
	return 0
}

func main() {
	fmt.Println(foo())
}
