package main

import "fmt"


func main() {
	
	for {
		var option int
		fmt.Println("\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("\n1. Start")
		fmt.Println("2. Exit")
		
		fmt.Print("\nSelect option: ")
		fmt.Scanf("%d", &option)

		
		

		if option == 1 {
			var num1, num2 int
			var operand string

			fmt.Print("\nEnter first number: ")
			fmt.Scanf("%d", &num1)

			fmt.Print("Enter an operand(+,-,*,/): ")
			fmt.Scanf("%s", &operand)

			fmt.Print("Enter second number: ")
			fmt.Scanf("%d", &num2)

			if operand == "+" {
				result := num1 + num2
				fmt.Printf("\n%d + %d = %d\n", num1, num2, result)
			} else if operand == "-" {
				result := num1 - num2
				fmt.Printf("\n%d - %d = %d\n", num1, num2, result)

			} else if operand == "*" {
				result := num1 * num2
				fmt.Printf("\n%d * %d = %d\n", num1, num2, result)
			} else if operand == "/" {
				if num2 != 0 {
					result := num1 / num2
					fmt.Printf("\n%d / %d = %d\n", num1, num2, result)
				} else {
					fmt.Println("\nZero division not allowed")
				}
			} else {
				fmt.Printf("\n\nUnsupported operand --> %s\nEnter one of this operands (+,-,*,/)\n", operand)
			}
		} else {
			break
		}
			
	}
}