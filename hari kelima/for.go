package main

import "fmt"

func main() {
	number := 123
	result := 0
	// 321

	for number != 0 {
		mod := number % 10
		result = result*10 + mod
		number /= 10
	}

	fmt.Println(result)
}
