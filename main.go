package main

import (
	"fmt"

	"github.com/nicopier/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(26031)
	fmt.Println(estado)
	fmt.Println(texto)
}
