package main

import "fmt"

func Tambah(nilai1 int, nilai2 int) int {
	return nilai1 + nilai2
}

func main() {
	tambah := Tambah(1, 2)

	fmt.Println(tambah)
}
