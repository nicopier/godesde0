package mapas

import "fmt"

func Mapas() {
	paises := make(map[string]string, 5)

	paises["Argentina"] = "Buenos Aires"
	paises["Brasil"] = "Brasilia"
	//fmt.Println(paises)
	//fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":       39,
		"Real Madrid":     38,
		"Chivas":          37,
		"Rosario Central": 30,
	}
	fmt.Println(campeonato)
	/*
		for equipo, puntaje := range campeonato {
			fmt.Printf("Equipo: %s tiene %d puntos\n", equipo, pu3ntaje)
		}
	*/
	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
}
