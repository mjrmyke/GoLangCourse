//Myke Walker

package main

import (
	"fmt"
	"math"
)

func main() {

	var a float64
	var b float64
	var z float64

	fmt.Printf("%s", "Please enter a number")
	fmt.Scanln(&a)

	fmt.Printf("%s", "Please enter a larger number")
	fmt.Scanln(&b)

	fmt.Printf("%s\n", "When the second number is divided by the first number, the remainder is: ")

	z = math.Remainder(a, b)
	fmt.Printf("%f\n", z)

}
