package main

import "fmt"

// var name = "Tim"
// var age int32 = 26


func main() {
	
	// name := "Tim"
	// age := 26
	// email := "tim@gmail.com"

	name, age, email := "Tim", 26, "tim@gmail.com"

	const someVar = true

	fmt.Println(name, age, email)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

}