package main

import "fmt"

//p1
type animal struct{
	specie string
	name string
}


func main(){
	//p1
	dog1 := animal{"dog", "one"}
	fmt.Println(dog1)
	fmt.Println(dog1.name)
	dog1.name = "Two"
	fmt.Println(dog1)
	
	//p2 
	dog1.SetName("Three")
	fmt.Println(dog1)

	//p3
	dog1.move()

	//p4 
	animalSlice := []animal{
		{"dog", "One"},
		{"dog", "Two"},
		{"cat", "Three"},
		{"cat", "Four"},
		{"bird", "Five"},
	}
	fmt.Println(animalSlice)
	//!!!!!!!!!!!!!!!!!!!
	fmt.Println(animalSlice[3])
	fmt.Println(animalSlice[2:4])

	//p5
	animalMap := make(map[string]animal)
	animalMap["A"] = animal{"dog", "One"}
	animalMap["B"] = animal{"dog", "Two"}
	animalMap["C"] = animal{"cat", "Three"}
	animalMap["D"] = animal{"bird", "Four"}

	fmt.Println(len(animalMap))
	animalMap["C"] = animal{"bird", "Three"}

	fmt.Println(len(animalMap))

	v1, ok1 := animalMap["B"]
	if (ok1){
		fmt.Println("B", v1)
		delete(animalMap, "B")
		fmt.Println(animalMap)
	}
}

//p2
func (a *animal) SetName(newName string) {
	a.name = newName;
} 

//p3
func (a animal) move() {
	fmt.Println(a.specie + " " + a.name + " move")
}