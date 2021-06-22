package main

import (
	"fmt"
	"reflect"
)

type Creature struct {
	Species string
}

func main() {
	var creature Creature = Creature{Species: "shark"}

	fmt.Printf("1) %+v\n", creature)
	changeCreature(&creature)
	fmt.Printf("3) %+v\n", creature)


	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())
	fmt.Println(strSlice)
}

func changeCreature(creature *Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}