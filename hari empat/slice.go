package main

import "fmt"

func main() {
	var mount = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustust",
		"oktober",
		"september",
		"november",
		"desember",
	}

	slice1 := mount[2:4]

	// fmt.Println(cap(slice1))
	// fmt.Println(len(slice1))

	slice12 := append(slice1, "fuad")
	slice1[0] = "bukan maret"
	fmt.Println(slice12)
	fmt.Println(slice1)
	fmt.Println(mount)

	// membuat slice di awal
	newSlice := make([]string, 3, 5)

	newSlice[0] = "teuku"
	newSlice[1] = "fuad"
	newSlice[2] = "maulana"

	fmt.Println(newSlice)

}
