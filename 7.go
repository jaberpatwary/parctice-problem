package main

import "fmt"

/*9. Write a program that swaps the values of two variables.*/

func swap(a, b int) (int, int) {

	return b, a
}

func main() {
	var x, y int

	fmt.Println("first number:")
	fmt.Scan(&x)
	fmt.Println("second number:")
	fmt.Scan(&y)

	x, y = swap(x, y)

	fmt.Println("after the swap:")
	fmt.Println("x=", x)
	fmt.Println("y=", y)

}
