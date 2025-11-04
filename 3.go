package main

/*5. Write a program that takes 3 inputs from the user: 1 integer, 1 floating point
number and 1 character. Then display what the user entered using cout*/

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
