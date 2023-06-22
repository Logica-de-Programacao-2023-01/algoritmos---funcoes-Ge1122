package main

import (
	"errors"
	"fmt"
	"strings"
)
func concatenacaostring(lista []string) (string, error) {
	if len(lista) == 0 {
		return "", errors.New("Slice está vazio")
	}
	return strings.Join(lista, ","), nil
}

func main() {
	lista := []string{"maçã", "banana", "laranja"}
	result, err := concatenacaostring(lista)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

