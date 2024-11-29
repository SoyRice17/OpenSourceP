package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name   string
	age    int
	grade  int
	score  int
	rating string
}

func randomName() string {
	char := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nameLength := rand.Intn(5) + 5
	name := ""
	for i := 0; i < nameLength; i++ {
		name += string(char[rand.Intn(len(char))])
	}
	return name
}

func randomScore() int {
	return rand.Intn(100)
}

func randomGrade() int {
	return rand.Intn(6)
}

func calculateScore(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else {
		return "F"
	}
}

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func createStudent() Student {
	name := randomName()
	age, _ := strconv.Atoi(getInput("Enter your age: "))
	grade := randomGrade()
	score := randomScore()
	rating := calculateScore(score)
	return Student{name, age, grade, score, rating}
}

func main() {
	students := []Student{}
	for {
		student := createStudent()
		students = append(students, student)
		if len(students) == 5 {
			break
		}
	}
	for _, student := range students {
		fmt.Println(student)
	}
}
