package main

import (
	"fmt"
	"strconv"
)
func main() {
	//p1
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))

	//p2
	urls := []string{
		"www.google.com",
		"www.facebook.com",
	}

	//p3
	for i := 2; i < 8; i++ {
		urls = append(urls, "www.web" + strconv.Itoa(i) + ".com")
	}
	
	//p4
	for _, url := range urls {
		fmt.Println(url)
	}

	//p5
	fmt.Println(sum(2, 3))
	fmt.Println(sum(2, 3, 4))
	/////////////////////////
	fmt.Println(urls[4])
}

// p1
func fibonacci(num int) int {
	if num <= 1 {
		return 0
	} else if num == 2 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

//p5
func sum (nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
