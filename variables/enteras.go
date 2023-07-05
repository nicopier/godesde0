package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int //siempre que inicializmaos una variable si lo le pasamos valor va a tomar el menor valor que permite la variable
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun= ", intcomun)
	fmt.Println("intcomun= ", intde32)
	fmt.Println("intcomun= ", intde64)
}
