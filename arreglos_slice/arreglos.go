package arreglos_slice

import (
	"fmt"
)

// var tabla [10]int //el numero que va dentro de [] es para deir cuantos elementos va a tener.
var tabla [10]int = [10]int{10, 0, 59}

var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[4] = 54

	tabla2 := [10]string{"pablo", "juan", "pepito"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		if tabla[i] != 0 {
			fmt.Println(tabla[i])
		}
	}
	matriz[7][24] = 15
	fmt.Println(matriz)
}
