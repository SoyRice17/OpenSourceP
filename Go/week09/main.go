package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	goalNumber := rand.Intn(6) + 1 // 1~3 사이의 랜덤한 수 생성
	fmt.Printf("%d\n", goalNumber)

	for i := 0; i < 3; i++ {
		fmt.Printf("%d번의 기회가 있습니다. / 숫자 입력 : ", 3-i)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guessedNumber, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guessedNumber)

		if guessedNumber == goalNumber {
			fmt.Println("정답입니다.")
		} else if guessedNumber > goalNumber {
			fmt.Println("입력하신 수는 정답보다 큰 수 입니다.")
		} else {
			fmt.Println("입력하신 수는 정답보다 작은 수 입니다.")
		}
	}
}
