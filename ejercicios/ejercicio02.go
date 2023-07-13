package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EjercicioDos() string {
	var numero int
	var err error
	var texto string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese numero:")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ingrese un numero valido: ")
			} else {
				break
			}
		}
	}
	for multiplicador := 1; multiplicador <= 10; multiplicador++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, multiplicador, numero*multiplicador)
	}

	return texto
}
