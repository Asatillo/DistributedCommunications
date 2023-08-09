package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var result1 = inc(1)
	fmt.Println(result1)

	var result2 = isEven(3)
	fmt.Println(result2)

	var result3 = devides(3, 9)
	fmt.Println(result3)

	var result4 = area(3, 4)
	fmt.Println(result4)

	var sum, err = stringAdd("4", "6")
	fmt.Println(sum, err)

	var result6 = isUpper("APPLE")
	fmt.Println(result6)
}

// p1
func inc(x int) int {
	return x + 1
}

// p2
func isEven(x int) bool {
	return x%2 == 0
}

// p3
func devides(first int, second int) bool {
	return second%first == 0
}

// p4
func area(a int, b int) int {
	return a * b
}

// p5
func stringAdd(x string, y string) (int, error) {
	xInt, xErr := strconv.Atoi(x)
	yInt, yErr := strconv.Atoi(y)
	if xErr != nil || yErr != nil {
		return -1, errors.New("Failed to convert")
	} else {
		return xInt + yInt, nil
	}
}

// p6
func isUpper(text string) bool {
	return strings.ToUpper(text) == text
}
