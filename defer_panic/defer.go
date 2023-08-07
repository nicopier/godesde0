package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	fmt.Println("este es un primer mensaje")
	defer fmt.Println("Este es el mensaje Final") //defer ejecuta lo que se le pase al final del codigo siempre.

	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que generó Panic: %v", reco) // Fatalf aborta el sistema usando os.Exit para no recursionar tanto
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}
