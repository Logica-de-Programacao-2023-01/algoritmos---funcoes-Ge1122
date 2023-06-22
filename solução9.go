package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	x := "Eu fui ao cinema."
	words, err := separarPalavras(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(words)
	}
}

func separarPalavras(x string) ([]string, error) {
	if len(x) == 0 {
		return nil, errors.New("A string está vazia")
	}
	words := strings.Split(x, " ")
	return words, nil
}
