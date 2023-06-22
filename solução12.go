package main

import (
	"errors"
	"fmt"
	"strings"
)


func main() {
	lista := []string{"cavalo", "mundo", "orelha", "brasilia", "DF"}
	x, err := maiuscula(lista)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}

func maiuscula(lista []string) (string, error) {
	if len(lista) == 0 {
		return "", errors.New("Slice está vazia")
	}
	var result []string
	for _, y := range lista {
		if len(y) > 0 && strings.ToUpper(string(y[0])) == string(y[0]) {
			result = append(result, y)
		}
	}
	return strings.Join(result, " "), nil
}
