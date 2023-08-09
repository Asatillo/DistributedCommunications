package main

import (
	"fmt"
)

func main(){
	r1 := request{2, 3}
	result := product(r1)
	fmt.Println(result)
}

func product(r1 request) int{
	return r1.first*r1.second
}
type request struct{
	first int
	second int
}