package main

// SumarGo realiza una suma básica de dos enteros de 64 bits de forma nativa.
func SumarGo(a int64, b int64) int64 {
	return a + b
}

// MultiplicarGo realiza una multiplicación de dos enteros de 64 bits de forma nativa.
func MultiplicarGo(a int64, b int64) int64 {
	return a * b
}

// SumarVectorGo recorre un slice de enteros y acumula su suma de forma nativa.
func SumarVectorGo(vector []int64) int64 {
	var suma int64 = 0
	for i := 0; i < len(vector); i++ {
		suma += vector[i]
	}
	return suma
}
