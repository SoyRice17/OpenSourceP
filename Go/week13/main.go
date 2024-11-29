package main

import (
	"fmt"
	"os"
	"reflect"
)

func text(strs ...string) {
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	slice := os.Args[:1]
	fmt.Println(slice[0])
	fmt.Printf("%T\n", slice[0])
	slice = append(slice, "i", "n", "h", "a")
	fmt.Println(slice, len(slice))
	text("123", "456", "789")
}
