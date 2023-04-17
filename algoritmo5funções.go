package main

import "fmt"

func primeiroElementoIgual(conj []int, x int) int {

	idxIgual := -1

	for i := 0; i < len(conj); i++ {

		if conj[i] == x {

			idxIgual = i
			break

		}

	}

	return idxIgual

}

func main() {
	conj := []int{2412341413424, 32424, 32, 2, 3, 32}
	conjVazio := []int{0, 1, 2, 3}
	x := 32

	fmt.Println(primeiroElementoIgual(conj, x))
	fmt.Println(primeiroElementoIgual(conjVazio, x))
}
