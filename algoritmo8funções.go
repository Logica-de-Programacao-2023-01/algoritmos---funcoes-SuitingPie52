package main

import (
	"fmt"
)

func slicePares(x []int) ([]int, error) {

	pares := []int{}

	for _, ranX := range x {

		if ranX%2 == 0 {
			pares = append(pares, ranX)
		}
	}

	if len(pares) == 0 {

		return pares, fmt.Errorf("conjunto não tem par")

	}

	return pares, nil

}

func main() {

	conj := []int{2312, 32, 77, 10, 3456, 11, 1231, 3334}
	conjImpar := []int{3, 5, 7, 9, 11}

	fmt.Print("Os pares do seu conjunto são: ")
	fmt.Println(slicePares(conj))
	fmt.Print("Os pares do seu conjunto são: ")
	fmt.Println(slicePares(conjImpar))

}
