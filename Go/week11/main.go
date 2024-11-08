// 두개의 수 사이의 소수 출력
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		for i := 3; i*i <= n; i += 2 {
			if n%i == 0 { // 나누어 떨어지는 수가 있으면
				return false // 소수가 아님
			}
		}
	}
	return true
}

func getInteger() (int, error) {
	r := bufio.NewReader(os.Stdin)
	a, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}

	a = strings.TrimSpace(a) // python strip
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	fmt.Print("첫 번째 정수 입력 : ")
	n1, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("두 번째 정수 입력 : ")
	n2, err := getInteger()
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) { // 나누어 떨어지는 수는 없어야함
			fmt.Printf("%d ", i)
		}
	}
}
