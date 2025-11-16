package main

/*'Write a program that computes the result using the formula
𝑟𝑒𝑠𝑢𝑙𝑡 = 𝑎 ∗ (𝑏 − 𝑐) + 1.37'*/

import "fmt"

func main() {

	var a float64
	var b float64
	var c float64
	fmt.Println("Enter a:")
	fmt.Scanf("%f", &a)

	fmt.Println("Enter b:")
	fmt.Scanf("%f", &b)

	fmt.Println("Enter c:")
	fmt.Scanf("%f", &c)

	result := a*(b-c) + 1.37
	fmt.Println("result:", result)

}
