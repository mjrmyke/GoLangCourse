//Myke Walker

package main

import (
	"fmt"
)

func main() {
	BOOL := (true && false) || (false && true) || !(false && false)
	fmt.Println(BOOL)

	//Prints True
}
