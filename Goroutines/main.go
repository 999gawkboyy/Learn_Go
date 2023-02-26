package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := []string{"qwe", "asd", "rty", "fgh", "feg", "dgd", "t3e", "ger", "qwg", "23r", "g4g"}
	for _, person := range people {
		go isSexy(person, c)
	}

	fmt.Println(len(people))

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is sexy"
}
