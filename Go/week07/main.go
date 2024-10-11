package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// = var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	fmt.Print("input your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}
}
