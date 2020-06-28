package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("linha")

	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.1415
	//fmt.Println("O valor de X é " + x) Isso n é possíve por incompatibilidade de tipos
	xs := fmt.Sprint(x)
	fmt.Println("O valor de X é " + xs)
	fmt.Println("O valor de X é", x)

	fmt.Printf("O valor de X é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Opa!"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
