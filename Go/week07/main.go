package main

import (
	"fmt"
	"strings"
)

func main() {
	var foods string = "국?, 비빔?, 볶음?"
	var replacer *strings.Replacer = strings.NewReplacer("?", "밥")
	//= replacer := strings.NewReplacer("?", "밥")

	fmt.Println(foods)
	fmt.Println(replacer.Replace(foods))
}
