package main

import "fmt"

func main(){
	//Estrutura homogênica (mesmo tipo) e estática (tamanho fixo)
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 8.5, 9.4
	//notas[3] = 9.2 //dá erro de compilacao pois o array nao contem esse indice
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total/float64(len(notas))
	fmt.Printf("Média: %.2f\n", media)
}
