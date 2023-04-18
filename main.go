package main

import (
	"fmt"
)

type Input struct {
	num1 int
	num2 int
}

func (data Input) Addition() int {
	return data.num1 + data.num2
}

func (data Input) Substraction() int {
	return data.num1 - data.num2

}

func (data Input) Multiplication() int {
	return data.num1 * data.num2

}

func (data Input) Division() int {
	return data.num1 / data.num2

}

func main() {

	var choice string

	data := Input{

		num1: 5,
		num2: 6,
	}

	fmt.Println("Enter first num")
	fmt.Scan(data.num1)

	fmt.Println("Enter your operater")
	fmt.Scan(&choice)

	fmt.Println("Enter second num")
	fmt.Scan(data.num2)

	switch choice {
	case "+":
		fmt.Printf("Addition: %d \n", data.Addition())
	case "-":
		fmt.Printf("Difference: %d \n", data.Substraction())
	case "*":
		fmt.Printf("Product: %d \n", data.Multiplication())
	case "/":
		fmt.Printf("Division: %d \n", data.Division())
	default:
		fmt.Println("Wrong option try with +, -, / and *")
	}

}
