package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(3) + 1 // 1~3 사이의 랜덤한 수 생성
	fmt.Printf("%d\n", target)
}
