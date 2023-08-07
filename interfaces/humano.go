package interfaces

type Humano interface {
	EstarVivo() bool
	Respirar()
	Pensar()
	Comer()
	Sexo() string
}
