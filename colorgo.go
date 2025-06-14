package colorgo

import (
	"fmt"
)

// pequeño codigo para poner colores en consola para el lenguaje de GO

const ESCAPE = "\033"

var colores = map[string]string{
	// colores basicos
	"ROJO":     "0;31",
	"VERDE":    "0;32",
	"AMARILLO": "0;33",
	"AZUL":     "0;34",
	"CELESTE":  "0;36",
	"BLANCO":   "0;37",
	"NEGRO":    "0;30",
	"VIOLETA":  "0;35",
	"RESET":    "0",

	// subrayados
	"SUB_ROJO":     "4;31",
	"SUB_VERDE":    "4;32",
	"SUB_AMARILLO": "4;33",
	"SUB_AZUL":     "4;34",
	"SUB_CELESTE":  "4;36",
	"SUB_BLANCO":   "4;37",
	"SUB_NEGRO":    "4;30",
	"SUB_VIOLETA":  "4;35",

	// fondo
	"F_ROJO":     "41",
	"F_VERDE":    "42",
	"F_AMARILLO": "43",
	"F_AZUL":     "44",
	"F_CELESTE":  "46",
	"F_BLANCO":   "47",
	"F_NEGRO":    "40",
	"F_VIOLETA":  "45",
}

// formateo de colores: ROJO, VERDE, AMARILLO, AZUL, VIOLETA, RESET, BLANCO ...
func Formateo(color string) string {

	return fmt.Sprintf("%s[%sm", ESCAPE, colores[color])
}
