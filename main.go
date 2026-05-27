package main

import "fmt"

// Declaración de funciones implementadas en ensamblador
func SumarASM(a int64, b int64) int64
func MultiplicarASM(a int64, b int64) int64
func SumarVectorASM(vector []int64) int64

func main() {
	var a int64 = 10
	var b int64 = 5

	vector := []int64{1, 2, 3, 4, 5}

	fmt.Println("=====================================")
	fmt.Println("Integración Go + Ensamblador")
	fmt.Println("Arquitectura x86-64")
	fmt.Println("=====================================")

	// Comparativa de Sumas y Multiplicaciones
	fmt.Printf("Suma Go:  %d | Suma ASM:  %d\n", SumarGo(a, b), SumarASM(a, b))
	fmt.Printf("Mult Go:  %d | Mult ASM:  %d\n", MultiplicarGo(a, b), MultiplicarASM(a, b))

	// Comparativa de Vectores
	sumaGo := SumarVectorGo(vector)

	// 2. SOLUCIÓN COMPLETA: Ahora pasas el objeto vector directamente, limpio y sin operadores ampersand `&`
	sumaASM := SumarVectorASM(vector)

	fmt.Printf("Vec Go:   %d | Vec ASM:   %d\n", sumaGo, sumaASM)

	fmt.Println("=====================================")
	fmt.Println("Programa finalizado correctamente")
	fmt.Println("=====================================")
}
