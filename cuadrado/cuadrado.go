package cuadrado

import "fmt"

type Cuadrado struct {
	Ancho int
}

func (p *Cuadrado) Area() string {
	area := p.Ancho * p.Ancho
	fmt.Println("Valor proporcionado: ", p.Ancho)
	return fmt.Sprintf("El area del cuadrado es: %d", area)
}

func (p *Cuadrado) Perimetro() string {
	perimetro := 4 * p.Ancho
	fmt.Println("Valor proporcionado: ", p.Ancho)
	return fmt.Sprintf("\nEl perimetro del cuadrado es: %d \n", perimetro)
}
