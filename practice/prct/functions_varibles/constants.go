package main

import "fmt"

// Constants cannot be declared using the := syntax.
const Pi = 3.14

func main(){
	const World = "Some chinese!"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}