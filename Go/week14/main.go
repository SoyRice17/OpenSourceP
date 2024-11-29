package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calcCost(visitors []visitor) int {
	totalCost := 0
	for _, v := range visitors {
		totalCost += v.cost
	}
	return totalCost
}

func main() {
	var numVisitors int
	fmt.Print("Enter the number of visitors: ")
	fmt.Scanln(&numVisitors)

	var vs []visitor
	vs = make([]visitor, numVisitors)
	//visitors := make([]visitor, numVisitors)
	for i := 0; i < numVisitors; i++ {
		var age int
		fmt.Printf("Enter the age for visitor %d: ", i+1)
		fmt.Scanln(&age)
		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default:
			vs[i] = visitor{age: age, cost: 7000}
		}
	}
	fmt.Printf("Total cost: %d\n", calcCost(vs))
}
