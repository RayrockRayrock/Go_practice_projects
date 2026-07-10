package main

import "fmt"



func main(){
	
	var operator string = ""
	var num1 float64 = 0.0
	var num2 float64 = 0.0
	fmt.Printf("Enter the first num: ")
	fmt.Scanln(&num1)
	fmt.Printf("Enter operator: ")
	fmt.Scanln(&operator)
	fmt.Printf("Enter the second num: ")
	fmt.Scanln(&num2)

	switch operator {

	case "+":
		fmt.Printf("Result is %f\n", num1 + num2)
	case "-":
		fmt.Printf("Result is %f\n", num1 - num2)
	case "*":
		fmt.Printf("Result is %f\n", num1 * num2)
	case "/":
		if num2 == 0{
			fmt.Printf("Can't divide by zero!\n")
			break;
		}
		fmt.Printf("Result is %f\n", num1 - num2)
		
	}
}
