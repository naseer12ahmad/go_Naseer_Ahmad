package main

import (
	"fmt"
)

func main() {

	// num1 := 85

	// num2 := 100

	var num1 int
	// var num2 int
	// var num3 = 70
	// var num4 = 84
	// var num5 = 55
	// var num6 = 69
	// var num7 = 40
	// var num8 = 50
	// var num9 = 0
	// var num10 = 39

	fmt.Println("Inset number: ")
	fmt.Scan(&num1)

	if num1 > 85 && num1 < 100 {
		fmt.Println("A")
	} else if num1 > 70 && num1 < 84 {
		fmt.Println("B")
	} else if num1 > 55 && num1 < 69 {
		fmt.Println("C")
	} else if num1 > 40 && num1 < 50 {
		fmt.Println("D")
	} else if num1 > 0 && num1 < 39 {
		fmt.Println("E")
	} else {
		fmt.Println("Invalid number")
	}

}
