package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicaçao entre as goroutines
// Channel é um tipo

func doisTresQuatroVezes(base int, c chan int){
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	// enquanto o canal c nao for lido ele nao segue

	time.Sleep(time.Second)
	c <- 3 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 4 * base // enviando dados para o canal
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a,b := <-c, <-c //Recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}
