package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	//var i int = 55

	var f float32 = 4.30 // f := 4.30
	i := 55

	fmt.Println(reflect.TypeOf(f))
	fmt.Printf("%s\n", strings.Title("hello world"))
	fmt.Println(math.Ceil(3.99))

	fmt.Printf("Value i: %d\n", i)
	fmt.Print("Value i: ", i, "\n")
	fmt.Println("Value i:", i)
}
