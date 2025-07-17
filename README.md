# colorgo - Colores en Consola para Go 🎨

`colorgo` es un pequeño módulo para Go que te permite añadir fácilmente colores y efectos de texto a tus aplicaciones de consola. Con una interfaz simple y una variedad de opciones, puedes mejorar la legibilidad y estética de tus outputs.

## Instalación

```bash
go get github.com/Urban20/ColorGo
```

## Uso Básico

Importa el módulo y usa la función `Formateo`:

```go
package main

import (
	"fmt"
	"github.com/Urban20/ColorGo"
)

func main() {
	rojo := colorgo.Formateo("ROJO")
	reset := colorgo.Formateo("RESET")
	
	fmt.Printf("%s¡Error!%s Algo salió mal\n", rojo, reset)
}
```

## Guía de Colores y Efectos

### Colores Básicos
```go
fmt.Println(
	colorgo.Formateo("ROJO") + "Rojo" + colorgo.Formateo("RESET"),
	colorgo.Formateo("VERDE") + "Verde" + colorgo.Formateo("RESET"),
	colorgo.Formateo("AMARILLO") + "Amarillo" + colorgo.Formateo("RESET"),
)
```

| Nombre    | Muestra       |
|-----------|---------------|
| ROJO     | <span style="color:red">Texto rojo</span> |
| VERDE    | <span style="color:green">Texto verde</span> |
| AMARILLO | <span style="color:gold">Texto amarillo</span> |
| AZUL     | <span style="color:blue">Texto azul</span> |
| CELESTE  | <span style="color:deepskyblue">Texto celeste</span> |
| VIOLETA  | <span style="color:purple">Texto violeta</span> |
| BLANCO   | Texto blanco |
| NEGRO    | Texto negro |
| RESET    | Restaura colores |

### Texto Subrayado
```go
fmt.Println(colorgo.Formateo("SUB_AZUL") + "Texto subrayado" + colorgo.Formateo("RESET"))
```

| Nombre         | Efecto               |
|----------------|----------------------|
| SUB_ROJO      | Texto rojo subrayado |
| SUB_VERDE     | Texto verde subrayado|
| SUB_AMARILLO  | Texto amarillo subrayado|
| SUB_AZUL      | Texto azul subrayado |
| SUB_CELESTE   | Texto celeste subrayado|
| SUB_VIOLETA   | Texto violeta subrayado|

### Fondos Coloreados
```go
fmt.Println(colorgo.Formateo("F_AMARILLO") + "Fondo amarillo" + colorgo.Formateo("RESET"))
```

| Nombre      | Efecto          |
|-------------|-----------------|
| F_ROJO      | Fondo rojo      |
| F_VERDE     | Fondo verde     |
| F_AMARILLO  | Fondo amarillo  |
| F_AZUL      | Fondo azul      |
| F_CELESTE   | Fondo celeste   |
| F_VIOLETA   | Fondo violeta   |
| F_BLANCO    | Fondo blanco    |
| F_NEGRO     | Fondo negro     |

## Ejemplo Completo

```go
package main

import (
	"fmt"
	"github.com/Urban20/ColorGo"
)

func main() {
	// Mensajes de estado
	exito := colorgo.Formateo("VERDE") + "✓ Éxito" + colorgo.Formateo("RESET")
	alerta := colorgo.Formateo("AMARILLO") + "⚠ Alerta" + colorgo.Formateo("RESET")
	error := colorgo.Formateo("ROJO") + "✗ Error" + colorgo.Formateo("RESET")
	
	// Texto con formato combinado
	destacado := colorgo.Formateo("F_AZUL") + colorgo.Formateo("SUB_BLANCO") + "¡Destacado!" 
	
	fmt.Printf("%s: Operación completada\n", exito)
	fmt.Printf("%s: Campos incompletos\n", alerta)
	fmt.Printf("%s: Archivo no encontrado\n\n", error)
	fmt.Println(destacado + colorgo.Formateo("RESET"))
}
```

## Notas Importantes

1. **Siempre usa RESET** al final de los textos coloreados
2. **Combina efectos** concatenando múltiples códigos:
   ```go
   // Texto rojo subrayado con fondo amarillo
   combo := colorgo.Formateo("SUB_ROJO") + colorgo.Formateo("F_AMARILLO")
   ```
3. **Compatibilidad**: Funciona en terminales que soportan códigos ANSI (Linux, macOS, Windows 10+)
4. Para **Windows antiguo**, necesitarás habilitar el soporte ANSI:
   ```go
   // Añade esto al inicio de tu aplicación para Windows
   import "golang.org/x/sys/windows"
   
   func init() {
       stdout := windows.Handle(os.Stdout.Fd())
       var originalMode uint32
       windows.GetConsoleMode(stdout, &originalMode)
       windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
   }
   ```
