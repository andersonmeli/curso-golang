package main

import "fmt"

func main(){
	//var aprovados map[int]string
	//Maps devem ser inicializados
	aprovados := make(map[int]string)
	aprovados[12345] = "Maria"
	aprovados[67890] = "Pedro"
	aprovados[21312] = "Ana"

	for cpf, nome := range aprovados{
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345])
	delete(aprovados, 21312)
	fmt.Println(aprovados[21312])
}
