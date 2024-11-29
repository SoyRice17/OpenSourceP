package main

import "fmt"

func main() {
	var student struct {
		id   int
		name string
		gpa  float32
	}
	student.id = 20241234
	student.name = "홍길동"
	student.gpa = 4.5
	fmt.Println(student.id)
	fmt.Println(student.name)
	fmt.Println(student.gpa)
}
