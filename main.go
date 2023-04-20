package main

import (
	"calculater/calc"
	"fmt"
)

func main() {

	var choice string

	data := calc.Input{
		Num1: 5,
		Num2: 6,
	}

	fmt.Println("Enter first num")
	fmt.Scan(data.Num1)

	fmt.Println("Enter your operater")
	fmt.Scan(&choice)

	fmt.Println("Enter second num")
	fmt.Scan(data.Num2)

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
	fmt.Println(data.Addition())
}
