package main

import (
	"errors"
	"fmt"
)

func funcao(slice []int, f func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice está vazia")
	}
	result := make([]int, len(slice))
	for i, num := range slice {
		result[i] = f(num)
	}
	return result, nil
}

func double(x int) int {
	return x * 2
}

func main() {
	slice := []int{10, 20, 30, 40, 50}
	result, err := funcao(slice, double)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
