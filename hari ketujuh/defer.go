package main

import "fmt"

func logger() {
	fmt.Println("selesain di jalankan")
}

func runAplication(num int) {
	fmt.Println("run aplication")

	result := num / 10

	fmt.Println("result ", result)
}

func main() {
	defer logger()

	runAplication(10)

}
