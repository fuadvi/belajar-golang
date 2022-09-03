package main

import "fmt"

func reverst(number int) {

	resul := 0

	for number != 0 {
		mod := number % 10
		resul = number*10 + mod
		number /= 10
	}

	fmt.Print(resul)
}

func main() {
	reverst(123)
}
