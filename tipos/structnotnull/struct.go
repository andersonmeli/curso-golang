package main

import "fmt"

type Config struct {
	host string
	port float64
}

func main() {
	//value
	var c1 Config
	c2 := Config{}
	c3 := *new(Config)

	//reference
	c4 := &Config{}
	c5 := new(Config)

	fmt.Println(&c1 == nil)
	fmt.Println(&c2 == nil)
	fmt.Println(&c3 == nil)
	fmt.Println(c4 == nil)
	fmt.Println(c5 == nil)

	fmt.Println(c1, c2, c3, c4, c5)
}