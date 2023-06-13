package main

import "fmt"

//Crie uma função que receba um slice de inteiros
//e retorne a média aritmética dos valores.

func main() {
	valores := []int{5, 5, 5, 5, 5}
	fmt.Println("A media é: ", medias(valores))
}
func medias(s []int) int {
	if len(s) == 0 {
		return 0
	}
	soma := 0
	for _, valor := range s {
		soma += valor
	}
	return int(soma) / int(len(s))
}
