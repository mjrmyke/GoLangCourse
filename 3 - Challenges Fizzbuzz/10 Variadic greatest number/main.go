//Myke Walker

package main

import (
	"fmt"
)

func largest(x ...int) int {
	var bignum int
	for _, i := range x {
		if i > bignum {
			bignum = i
		}
	}
	return bignum
}

func main() {
	numbers := largest(2, 4, 6, 8, 10, 100, 200, 500, 201, 101, 9, 7, 5, 3, 1)
	fmt.Println(numbers)

}
