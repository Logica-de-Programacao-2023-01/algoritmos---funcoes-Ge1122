package main

import 	"fmt"


func main() {
	slice := []int{3, 10, 4, 36, 15}
	fmt.Print(segundo(slice))
}

func segundo(slice []int) int {
	var maior int
	var segundomaior int
	for i := 0; i < len(slice); i++ {
		if slice[i] > maior {
			segundomaior = maior
			maior = slice[i]
		}
	}
	return segundomaior
}
