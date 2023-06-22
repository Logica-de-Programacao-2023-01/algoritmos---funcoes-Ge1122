package main

import (
	"errors"
	"fmt"
	"sort"
)
func main() {
	lista := []int{2,5,1,0,21,50,41,122,13}
	ordem, err := crescente(lista)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ordem)
	}
}

func crescente(lista []int) ([]int, error) {
	if len(lista) == 0 {
		return nil, errors.New("A lista está vazia")
	}
	sort.Ints(lista)
	return lista, nil
}

