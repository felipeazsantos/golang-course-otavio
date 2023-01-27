package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade int `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	// converter map ou estruct para json
	cachorroEmJSON, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJSON)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)

	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

	// converter json para struct ou map
	//json.Unmarshal()
}
