package main

import "fmt"

func recursiveLoop(number int) int {
	total := 1

	for i := number; i > 0; i-- {
		total *= i
	}

	return total
}

func recursiveFunc(number int) int {
	if number == 1 {
		return 1
	}

	return number * recursiveFunc(number-1)
}

func main() {
	fmt.Println(recursiveLoop(5))
	fmt.Println(recursiveFunc(5))

}
