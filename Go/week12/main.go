package main

import (
	"fmt"
	"reflect"
)

func main() {
	var gpa [5]float64 = [5]float64{2.3, 3.4, 4.5, 5.6, 6.7} // 배열
	gpa_slice := gpa[1:4]                                    // 슬라이스
	fmt.Println(gpa_slice, reflect.TypeOf(gpa_slice))
	fmt.Println(gpa, reflect.TypeOf(gpa))
}
