package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nicopier/godesde0/ejercicios"
)

var fileName string = ("./files/txt/tabla.txt")

func GrabaTabla() {
	var texto string = ejercicios.EjercicioDos()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("error al crear el archivo" + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.EjercicioDos()
	if !Append(fileName, texto) { // el ! es para negar la funcion en vez de poner Append(fileName, texto) == false
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error al abrir el archivo" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto) // el _ va cuando la variable no se usa

	if err != nil {
		fmt.Println("error al escribir el archivo" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error al abrir el archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
