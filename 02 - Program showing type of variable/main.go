package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	time := time.Now()                      //calls for current time in string
	minutes := time.Minute()                //call time.minute function for int of current minute
	fmt.Printf("%s\n", "Current time is: ") //Print string
	fmt.Printf("%s\n", time)                //Print Time - type strring

	fmt.Printf("%s\n", "Time variable is of type: ") //Print string
	fmt.Println(reflect.TypeOf(time))

	fmt.Printf("%s\n", "Minutes variable is of type: ") //Print string
	fmt.Println(reflect.TypeOf(minutes))

}
