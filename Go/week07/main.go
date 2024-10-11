package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// = var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Print("이름을 입력하세요: ")
	name, err := reader.ReadString('\n')
	fmt.Println(name)
	fmt.Println(err)
}
