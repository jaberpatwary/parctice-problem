package main

import "fmt"

/*6. Write a program that displays the sum, difference, product, quotient (division), and
remainder of two integers provided by the user.*/

func main() {

	var number1 int
	var number2 int

	fmt.Println("number1")
	fmt.Scanf("%d", &number1)
	fmt.Println("number2")
	fmt.Scanf("%d", &number2)

	fmt.Println("sum:", number1+number2)
	fmt.Println("difference:", number1-number2)
	fmt.Println("product:", number1*number2)
	fmt.Println("quotient:", number1/number2)
	fmt.Println("reminder:", number1%number2)

}
