package main

import "fmt"

func someDez(x int) int {

	return x + 10

}

func aplicarFunc(conj []int, funcao func(int) int) ([]int, error) {

	i := 0

	for i = 0; i < len(conj); i++ {
		conj[i] = funcao(conj[i])
	}

	if len(conj) == 0 {
		return nil, fmt.Errorf("conjunto vazio")
	}

	return conj, nil
}

func main() {
	conj := []int{2312, 32, 77, 10}
	conjVazio := []int{}

	fmt.Print("Sua função aplicada no conjunto é: ")
	fmt.Println(aplicarFunc(conj, someDez))
	fmt.Print("Sua função aplicada no conjunto é: ")
	fmt.Println(aplicarFunc(conjVazio, someDez))

}
