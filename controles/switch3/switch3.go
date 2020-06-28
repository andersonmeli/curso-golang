package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcão"
	default:
		return "não sei!"
	}
}

func main(){
	fmt.Println(tipo(3.2))
	fmt.Println(tipo(30))
	fmt.Println(tipo("teste"))
	fmt.Println(tipo(func(){}))
	fmt.Println(tipo(time.Now()))
}