package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile" // go get github.com/headfirstgo/datafile
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if !matched { // 중복된 이름이 없으면
			names = append(names, line) // 이름 배열에 추가
			counts = append(counts, 1)  // 카운트 배열에 1 추가
		}
	}

	for i, name := range names {
		fmt.Printf("Votes for %s: %d\n", name, counts[i])
	}

	//lines, err := datafile.GetStrings("votes.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//counts := make(map[string]int)
	//for _, line := range lines {
	//	counts[line]++
	//}
	//for name, count := range counts {
	//	fmt.Printf("Votes for %s: %d\n", name, count)
	//}
}
