package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Println("Enter the value of num1")
	fmt.Scanln(&num1)

	fmt.Println("Enter the value of num2")
	fmt.Scanln(&num2)

	if num1 > num2 {
		fmt.Printf("The largest number is %d\n", num1)

	} else if num2 > num1 {
		fmt.Printf("The largest number is %d\n", num2)
	} else {
		fmt.Println("Both the numbers are equal")
	}
}
