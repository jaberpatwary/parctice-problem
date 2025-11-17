package main

import "fmt"

/*7.Write a program that calculates the number of seconds, minutes and hours in 1 year.*/

func main() {

	day := 365
	hours := day * 60
	minutes := hours * 60
	seconds := minutes * 60

	fmt.Println("hours:", hours)
	fmt.Println("minutes:", minutes)
	fmt.Println("second:", seconds)
}
