package main

import "fmt"

//10. Write a program that swaps the values of two variables without using a 3rd variable.

func swap(a, b int) (int, int) {

	a = a + b
	b = a - b
	a = a - b

	return a, b

}

func main() {

	var x int
	var y int
	fmt.Println("first number:")
	fmt.Scan(&x)
	fmt.Println("second number:")
	fmt.Scan(&y)

	x, y = swap(x, y)
	fmt.Println("after swap:")
	fmt.Println("a=", x)
	fmt.Println("b=", y)
}
