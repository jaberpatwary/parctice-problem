package main

import "fmt"

func main() {

	var name string
	var age int
	var class float32
	var isStudent bool

	fmt.Println("Name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Age:")
	fmt.Scanf("%d", &age)
	fmt.Println("Class")
	fmt.Scanf("%f", &class)

	fmt.Println("Student")
	fmt.Scanf("%t", &isStudent)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(class)
	fmt.Println(isStudent)
}
