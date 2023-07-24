package main

import "fmt"

func main() {
	fmt.Println("Bem-vindo(a) ao conversor")
	fmt.Println("Selecione uma das opções:")
	fmt.Println("1. De metros para centímetros")
	fmt.Println("2. De centímetros para metros")

	var selecao int
	fmt.Scanln(&selecao)

	switch selecao {
	case 1:
		fmt.Println("Quantos metros?")
		var metros float32
		fmt.Scanln(&metros)
		centimetros := metros * 100
		fmt.Printf("%.2f metros são %.2f centímetros", metros, centimetros)
	case 2:
		fmt.Println("Quantos centímetros?")
		var centimetros float32
		fmt.Scanln(&centimetros)
		metros := centimetros / 100
		fmt.Printf("%.2f centímetros são %.2f metros", centimetros, metros)
	default:
		fmt.Println("Selecione uma opção correta")
	}

}
