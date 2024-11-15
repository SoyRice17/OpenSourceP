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
	fmt.Println(dates[0], dates[1], dates[2])
}
