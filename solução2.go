package main

func QuantasVogais(x string) int {
	Frequencia_vogais := make(map[rune]int)
	quantidade_vogais := 0
	for _, letra := range x {
		Frequencia_vogais[letra] += 1
	}
	for letra, count := range Frequencia_vogais {
		if letra == 'a' || letra == 'e' || letra == 'i' || letra == 'o' || letra == 'u' || letra == 'A' || letra == 'E' || letra == 'I' || letra == 'O' || letra == 'U' {
			quantidade_vogais += count
		}
	}
	return quantidade_vogais
}
