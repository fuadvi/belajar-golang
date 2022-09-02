package main

import "fmt"

func main() {

	var bulan [3]string
	bulan[0] = "teuku"
	bulan[1] = "fuad"
	bulan[2] = "maulana"

	fmt.Println(bulan[0])
	fmt.Println(bulan[1])
	fmt.Println(bulan[2])

	var age = [3]int{
		20,
		19,
		18,
	}

	fmt.Println(age[0])
	fmt.Println(age[1])
	fmt.Println(age[2])
}
