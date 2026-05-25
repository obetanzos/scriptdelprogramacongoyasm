package main

import "fmt"

// Declaración de funciones implementadas en ensamblador.
// En Go no se usa extern como en C.
// Estas funciones deben implementarse en un archivo .s.
func SumarASM(a int64, b int64) int64
func MultiplicarASM(a int64, b int64) int64
func SumarVectorASM(vector *int64, n int64) int64

func main() {
	var a int64 = 10
	var b int64 = 5

	vector := []int64{1, 2, 3, 4, 5}

	fmt.Println("=====================================")
	fmt.Println("Integración Go + Ensamblador")
	fmt.Println("Arquitectura x86-64")
	fmt.Println("=====================================")

	fmt.Println("Suma ASM:", SumarASM(a, b))
	fmt.Println("Multiplicación ASM:", MultiplicarASM(a, b))
	fmt.Println("Suma vector ASM:", SumarVectorASM(&vector[0], int64(len(vector))))

	fmt.Println("=====================================")
	fmt.Println("Programa finalizado correctamente")
	fmt.Println("=====================================")
}
