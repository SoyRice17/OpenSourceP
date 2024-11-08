package main

import (
	"fmt"
	"log"
	"week11/greeting"
	"week11/keyboard"
)

func main() {
	//greeting.Hi("Inha")
	//greeting.Hello("Havard")
	greeting.AllGreetings("Inha")
	n, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
