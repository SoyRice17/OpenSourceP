package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[1] = time.Unix(1457894000, 0)
	// fmt.Println(dates[1])

	dates := [3]time.Time{
		time.Unix(1257894000, 0),
		time.Unix(1457894000, 0),
		time.Unix(1657894000, 0),
	}
	// fmt.Println(dates[0], dates[1], dates[2]) // 배열 원소들
	// fmt.Println(dates)                       // 배열 자체
	// fmt.Printf("%#v\n", dates) // 배열 리터럴 (리터럴이란? 소스코드에 직접 표현되는 값)

	//for i := 0; i < len(dates); i++ { // 안전함ss
	for i, v := range dates { // 파이썬 enumerate와 유사함
		fmt.Println(i, v)
	}
}
