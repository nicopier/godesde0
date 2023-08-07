package arreglos_slice

import "fmt"

var tablaS []int = []int{2, 5, 4} // la diferencia entre slice y vector es que no se le asigna magnitud
var arreglo [10]int = [10]int{6, 98, 21, 644, 18, 36, 78, 9}

func MuestroSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]  // se puede hacer un slice de un vector
	porcion2 := arreglo[:5] // se puede hacer un slice de un vector
	porcion3 := arreglo[3:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20) // se crea un slice con 5 elementos y una capacidad de 20
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 1000000; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
