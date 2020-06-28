package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struc to json
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 1899,
		Tags:  []string{"Promocao", "Eletronico"},
	}

	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	var p2 produto
	jsonString := `{"id":2,"nome":"Smartphone","preco":2500,"tags":["Promocao","Eletronico"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}