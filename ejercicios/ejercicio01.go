package ejercicios

import (
	"fmt"
	"strconv"
)

func EjercicioCeroUno(entrada string) (int, string) {
	valor_convertido, err := strconv.Atoi(entrada)
	if err != nil {
		return 0, "hubo un error " + err.Error()
	}

	if valor_convertido > 100 {
		fmt.Println("El numero es mayor a 100")
	} else {
		fmt.Println("El numero es menor a 100")
	}
	return valor_convertido, entrada
}
