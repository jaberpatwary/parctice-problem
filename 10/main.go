/*
12. Write a program that displays a temperature in Celsius and in Kelvin after reading
the temperature in Fahrenheit from the user. Use the conversion formulae 5(𝐹 −
32) = 9𝐶 and 𝐶 = 𝐾 − 273.15.
*/
package main

import "fmt"

func convertTemp(F float32) (float32, float32) {
	C := (5.0 / 9.0) * (F - 32)
	K := C + 273.15

	return C, K

}

func main() {

	var F float32
	fmt.Printf("enter temperature:")
	fmt.Scan(&F)
	C, K := convertTemp(F)

	fmt.Println("Celsius", C)
	fmt.Println("Kelvin", K)

}
