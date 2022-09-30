package circulo

import (
	"fmt"
	"math"
)

type Circulo struct {
	Radio int
}

func (c *Circulo) Area() string {
	area := math.Pi * math.Pow(float64(c.Radio), 2)
	fmt.Println("Valor del Radio Proporcionado: ", c.Radio)
	return fmt.Sprintf("\n El Area del Circulo es: %.2f \n", area)
}
func (c *Circulo) Perimetro() string {
	perimetro := 2 * math.Pi * float64(c.Radio)
	fmt.Println("Valor del Radio Proporcionado: ", c.Radio)
	return fmt.Sprintf("\n El Perimetro del Circulo es: %.2f \n", perimetro)
}
