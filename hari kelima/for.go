package main

import "fmt"

func main() {
	number := 12345
	result := 0

	for number != 0 {
		remainder := number % 10
		result = result*10 + remainder
		number /= 10
	}

	fmt.Println(result)
}
