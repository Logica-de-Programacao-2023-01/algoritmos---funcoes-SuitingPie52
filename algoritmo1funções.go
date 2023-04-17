package main

import "fmt"

func aritmetica(x []int) (float64, error) {
	soma := 0
	media := 0.0

	if len(x) == 0 {
		return 0, fmt.Errorf("cannot divide 0 by 0")
	}

	for _, ranX := range x {

		soma += ranX

	}
	media = float64(soma) / float64(len(x))

	return media, nil
}

func main() {
	conj := []int{3, 8, 9, 20, 7}
	conjVazio := []int{}

	fmt.Print("Sua média é: ")
	fmt.Println(aritmetica(conj))
	fmt.Print("Sua média é: ")
	fmt.Println(aritmetica(conjVazio))
}
