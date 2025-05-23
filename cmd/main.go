package main

import (
	"fmt"
	"learn-go-with-tests/arrays"
)

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	result := arrays.Sum(numbers)
	fmt.Printf("Sum of numbers: %d\n", result)
}
