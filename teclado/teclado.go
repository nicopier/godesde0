package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	scanner := bufio.NewScanner(os.Stdin) //lee lo que ingresa el usuario usando el os standar input
	fmt.Println("Ingrese numero 1:")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text()) //convierte el texto a entero ( nose le pone := y se se pone = ya que la primer variable ya fue creada...)
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error()) //aca se aborta la aplicacion y se muestra el error
		}
	}

	fmt.Println("Ingrese numero 2:")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text()) //convierte el texto a entero ( nose le pone := y se se pone = ya que la primer variable ya fue creada...)
		if err != nil {
			panic("el dato ingresado es incorrecto" + err.Error()) //aca se aborta la aplicacion y se muestra el error
		}
	}
	fmt.Println("Ingrese leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
