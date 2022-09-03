package main

import "fmt"

func goodBy(name string) string {
	return "good bye " + name
}

func main() {
	result := goodBy

	fmt.Println(result("fuad"))
}
