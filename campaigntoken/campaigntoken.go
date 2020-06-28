package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"regexp"
	"strings"
)

func main() {
	novouuid := strings.Trim(strings.Replace(fmt.Sprint(uuid.New()), "-", "", -1), "[]")
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(novouuid, "")
	Substring := processedString[0:13]
	fmt.Println(AddCaracter(Substring) + "000")
}

func AddCaracter (uuidstringnum string) string {
	caracteres := []rune(uuidstringnum)
	dif := 13 - len(caracteres)
	if dif == 0 {
		return uuidstringnum
	}else {
		for i := len(caracteres); i < len(caracteres) + dif; i++ {
			uuidstringnum = uuidstringnum + "0"
		}
	}
	return uuidstringnum
}

//7630481714909002 token atual
//5622467846573000