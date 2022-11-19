package main

import (
	"fmt"
)

var num1 float32
var num2 float32
var operator string
var answer float32

func main() {
	fmt.Print(`enter first number: `)
	fmt.Scanln(&num1)
	fmt.Print(`enter second number: `)
	fmt.Scanln(&num2)
	fmt.Print(`enter the operator: + - * /   `)
	fmt.Scanln(&operator)
	fmt.Println(`Your want to calculate:`, num1, operator, num2)

	switch operator {
	case "+":
		fmt.Println("RESULT: ", num1+num2)
	case "-":
		fmt.Println("RESULT: ", num1-num2)
	case "/":
		if num2 == 0 {
			fmt.Println("It is bad idea to divide by zero ")
		}
		fmt.Println("RESULT: ", num1/num2)
	case "*":
		fmt.Println("RESULT: ", num1*num2)
	default:
		fmt.Printf("HOUSTON! WE HAVE A PROBLEM!!! (invalid operator:%s) Next time you'd better choose the operator from presented operators ;) BYE!", operator)
	}
}
