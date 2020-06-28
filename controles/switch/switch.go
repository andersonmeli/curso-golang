package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		//Em go o default do switch é sair do bloco assim que atende a condição, ao contrario de outras linguagens que usam o break para sair
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida!"
	}
}

func main(){
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(5.8))
}