package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

func main() {

	time := time.Now()                      //calls for current time in string
	minutes := time.Minute()                //call time.minute function for int of current minute
	fmt.Printf("%s\n", "Current time is: ") //Print string
	fmt.Printf("%s\n", time)                //Print Time - type strring

	templ, error := template.ParseFiles("tpl.gohtml") // Parse template file
	if error != nil {
		log.Fatalln(error)
	}

	if minutes%2 == 0 {
		fmt.Printf("%s\n", "The current minute is even")
		fmt.Printf("%d\n", minutes)

		error = templ.Execute(os.Stdout, nil) //Executes template if minute is even.
		if error != nil {
			log.Fatalln(error)
		}

	} else {
		fmt.Printf("%s\n", "The current minute is odd")
		fmt.Printf("%d\n", minutes)
	}

}
