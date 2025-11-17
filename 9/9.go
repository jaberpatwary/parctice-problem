/*
11. Write a program that reads the units in meter and converts it to feet. (1 inch = 2.54
cm).
*/
package main

import "fmt"

func meterTofeet(meter float32) float32 {
	feet := meter * 3.28084

	return feet

}

func main() {

	var meter float32

	fmt.Printf("Enter the meter:")
	fmt.Scan(&meter)

	feet := meterTofeet(meter)

	fmt.Println("feet:", feet)

}
