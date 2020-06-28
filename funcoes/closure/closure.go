package main

import "fmt"

func closure() func(){
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main(){
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
	//Conceito de closure. A função tem ideia do local que foi definida,
	//por isso o x é outro. Go sabe onde foi definido
}
