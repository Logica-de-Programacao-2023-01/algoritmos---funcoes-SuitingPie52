ALGORITMO INCOMPLETO
package main

import (
	"fmt"
)

func sliceCrescente(x []int) ([]int, error) {

	if len(x) == 0 {

		return x, fmt.Errorf("slice vazio")

	}

	idxMenor := 0
	y := []int{}
	menor := x[0]
	i := 0

	for len(x) >= 2 {

		for i = 0; len(x) > i && len(x) != 0; i++ {

			if x[i] < menor {

				menor = x[i]
				idxMenor = i

			}

		}

		y = append(y, menor)

		if idxMenor == len(x)-1 {
			x = append(x[:idxMenor])
		} else if idxMenor == 0 {
			x = append(x[idxMenor+1:])
		} else {
			x = append(x[:idxMenor], x[idxMenor+1:]...)
		}

		i = 0
		menor = x[0]

	}

	return y, nil

}

func main() {

	conj := []int{34, 3, 434, 23, 3, 22, 2, 0}
	conjVazio := []int{}

	fmt.Print("Seus números em ordem crescente são: ")
	fmt.Println(sliceCrescente(conj))
	fmt.Print("Seus números em ordem crescente são: ")
	fmt.Println(sliceCrescente(conjVazio))

}
