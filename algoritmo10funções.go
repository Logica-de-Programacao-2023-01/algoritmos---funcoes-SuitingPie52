package main

import (
	"fmt"
)

func sliceCrescente(x []int) ([]int, error) {

	if len(x) == 0 {

		return x, fmt.Errorf("slice vazio")

	}

	for i := 0; i < len(x); i++ {

		for j := 0; j < len(x); j++ {

			if x[j] > x[i] {

				x[j], x[i] = x[i], x[j]

			}
		}
	}

	return x, nil
}

func main() {

	conj := []int{34, 3, 434, 23, 3, 22, 2, 0}
	conjVazio := []int{}

	fmt.Print("Seus números em ordem crescente são: ")
	fmt.Println(sliceCrescente(conj))
	fmt.Print("Seus números em ordem crescente são: ")
	fmt.Println(sliceCrescente(conjVazio))

}
