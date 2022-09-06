package main

import (
	"errors"
	"fmt"
)

func bagi(nilai int, pembagi int) (int, error) {
	if nilai == 0 {
		return 0, errors.New("pembagi tidak bisa di bagi dengan 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := bagi(0, 10)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("errors", err.Error())
	}
}
