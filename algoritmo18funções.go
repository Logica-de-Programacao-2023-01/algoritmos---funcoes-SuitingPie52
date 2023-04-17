ALGORITMO INCOMPLETO

package main

import (
	"fmt"
)

func aplicarFuncEmIntSlice(x []int, y func(int) (int, error)) (int, error) {

	funcRanX := 0
	somaFuncRanX := 0

	if len(x) == 0 {
		return 0, fmt.Errorf("seu slice está vazio")
	}

	for _, ranX := range x {
		funcRanX, _ = y(ranX)
		somaFuncRanX += funcRanX
	}

	return funcRanX, nil

}

func main() {
	x := []int{232, 23, 111, 1234}
	y := []int{}

	fmt.Printf("Sua função aplicada em %d é: ", x)
	fmt.Println(aplicarFuncEmIntSlice(x, somaAlgarismos))
	fmt.Printf("Sua função aplicada em %d é: ", y)
	fmt.Println(aplicarFuncEmIntSlice(y, somaAlgarismos))
}
