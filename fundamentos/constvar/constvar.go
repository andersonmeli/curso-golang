package main

import (
	"fmt"
	m "math" //tambem é possível usar labels parar os imports
)

func main(){
	const PI float64 = 3.1415
	var raio = 3.2 //Tipo (float64) inferido pelo compilador

	//Forma reduzida de criar uma variável
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area) //Na concatenação o proprío Go adiciona o espaço

	//Outra forma de declaraçao de const e var são com blocos
	const(
		a = 1
		b = 2
	)

	var(
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false //também é possível múltiplas atribuiçoes numa única linha
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
