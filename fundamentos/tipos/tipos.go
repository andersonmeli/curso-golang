package main

import (
	"fmt"
	"math"
	"reflect"
)

func main(){
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal. só positivos...uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println( "O byte é", reflect.TypeOf(b))

	//com sinal
	i1 := math.MaxInt64
	fmt.Println("O Valor máximo de int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))

	//números reais float32, float64
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal de 49.99 é", reflect.TypeOf(49.99))//float64

	//boollean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//string
	s1 := "Olá meu nome é anderson"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string s1 é", len(s1))

	//string com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	anderson`
	fmt.Println("O tamanho da string s2 é", len(s2))


	//char????
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char) // não existe o tipo char em Go
}
