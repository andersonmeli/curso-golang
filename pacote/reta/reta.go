package main

import "math"

//Iniciando com letra maíuscula é PÚBLICO (visível dentro e fora do pacote)
//Iniciando com letra maíuscula é PRIVADO (visível apenas dentro do pacote)

//Exemplo
//Ponto é público
//ponto ou _Ponto é privado

type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64){
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}