package main

import (
	"errors"
	"fmt"
)
func main() {
	x:= 63
	y, err := sum(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(y)
	}
}

func sum(x int) (int, error) {
	if x <=-1 {
		return 0, errors.New("Número negativo")
	}
	var sum = 0
	for x != 0 {
		digito := x % 10
		sum += digito
		x /= 10
	}
	return sum, nil
}
