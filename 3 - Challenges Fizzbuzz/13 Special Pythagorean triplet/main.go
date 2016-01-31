//https://projecteuler.net/problem=9
//Myke Walker
//Problem Description: Pythagorean triplet is a set of numbers where a^2 + b^2 = c^2
//										 There is only one pythagorean triplet in which a+b+c = 1000
//										 find the variables in which a+b+c = 1000, then multiple them together.
//Output: 250800

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
				return (a * b * c)
			}

		}
	}
	return 0
}

func main() {
	fmt.Println(foo())
}
