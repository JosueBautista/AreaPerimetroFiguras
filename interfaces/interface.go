package interfaces

import (
	"fmt"
	)

type figura interface {
	Area() string
	Perimetro() string
}

func Calcular(figura Figura) {
	fmt.Println(figura.Area())
	fmt.Println(figura.Perimetro())
}
