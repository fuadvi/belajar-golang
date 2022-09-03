package main

import "fmt"

func Tambah(nilai1 int, nilai2 int) {

	fmt.Println(nilai1 + nilai2)

}

func main() {
	for i := 0; i < 10; i++ {
		Tambah(i, i+1)
	}
}
