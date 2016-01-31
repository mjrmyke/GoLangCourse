//Myke Walker

package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Printf("%s\n", "Please enter your name")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s %s\n", "Hello,", name) //Print string

}
