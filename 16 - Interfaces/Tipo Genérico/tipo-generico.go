package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1 : "String",
		float64(100): true,
		"String": "String",
		true: float64(12),
	}

	fmt.Println(mapa)
}
