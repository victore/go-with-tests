package arrays

import "fmt"

func Sum(numbers [5]int) int {
	letters := [7]string{"A", "B", "C", "D", "E", "F", "G"}

	for _, letter := range letters {
		fmt.Println("letter: ", letter)
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	fmt.Printf("Sum of numbers: %d\n", result)
}
