package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 13
	f := 12.9

	fmt.Printf("Value i: %d, Value f: %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f) invalid operation: i * f (mismatched types int and float64)
	//fmt.Printf("%d * %f = %f\n", i, f, i*int(f))
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)
	fmt.Println(reflect.TypeOf(i))
}
