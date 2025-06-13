package colorgo

import "fmt"

// pequeño codigo para poner colores en consola para el lenguaje de GO

// colores basicos
const ESCAPE = "\033"
const ROJO = "0;31"
const VERDE = "0;32"
const AMARILLO = "0;33"
const AZUL = "0;34"
const CELESTE = "0;36"
const BLANCO = "0;37"
const NEGRO = "0;30"
const VIOLETA = "0;35"
const RESET = "0"

// subrayados
const SUB_ROJO = "4;31"
const SUB_VERDE = "4;32"
const SUB_AMARILLO = "4;33"
const SUB_AZUL = "4;34"
const SUB_CELESTE = "4;36"
const SUB_BLANCO = "4;37"
const SUB_NEGRO = "4;30"
const SUB_VIOLETA = "4;35"

// fondo
const F_ROJO = "41"
const F_VERDE = "42"
const F_AMARILLO = "43"
const F_AZUL = "44"
const F_CELESTE = "46"
const F_BLANCO = "47"
const F_NEGRO = "40"
const F_VIOLETA = "45"

// formateo de colores: ROJO, VERDE, AMARILLO, AZUL, VIOLETA, RESET, BLANCO ...
func Formateo(color string) string {

	return fmt.Sprintf("%s[%sm", ESCAPE, color)
}
