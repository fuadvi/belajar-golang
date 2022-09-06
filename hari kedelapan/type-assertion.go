package main

import "fmt"

func random() interface{} {
	return 1
}

func main() {
	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("type tidak dikenali")

	}
}
