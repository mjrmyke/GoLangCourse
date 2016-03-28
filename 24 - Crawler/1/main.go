package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.reddit.com/r/dota2")
	if err != nil {
		log.Println("http get error:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("readall error:", err)
	}

	fmt.Println(string(body))
}
