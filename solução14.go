package main

import (
	"errors"
	"fmt"
)

func main() {
	lista1 := []int{1, 2, 3, 4, 5}
	lista2 := []int{6,7,8,9,10,11}
	i, err := newSlice(lista1, lista2)
  
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}

func newSlice(lista1 []int, lista2 []int) ([]int, error) {
	if len(lista1) == 0 && len(lista2) == 0 {
		return nil, errors.New("Uma das slices está vazia")
	}
	numMap := make(map[int]bool)
	for _, numero := range lista1 {
		numMap[numero] = true
	}
	var resultado []int
	for _, numero := range lista2 {
		if numMap[numero] {
			resultado = append(resultado, numero)
		}
	}
	return resultado, nil
}
