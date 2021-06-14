package main

import (
	"fmt"
	"math"
)

func dani(indice int) int {
	diferenca := int(math.Ceil(float64(indice) / 2))
	return indice + diferenca
}

func main(){
	fmt.Println(dani(11))
}

/* 0..  2..3..  5..6..  8..9..  11..12..  14..15
   0..  1..2..  3..4..  5..6..  7...8...  9...10*/