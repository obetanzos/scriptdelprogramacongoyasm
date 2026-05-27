package main

import "testing"

var resultado int64

func BenchmarkMultiplicarGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultado = MultiplicarGo(10, 5)
	}
}

func BenchmarkMultiplicarASM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resultado = MultiplicarASM(10, 5)
	}
}

func BenchmarkSumarVectorGo(b *testing.B) {
	vector := make([]int64, 100000)

	for i := range vector {
		vector[i] = int64(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resultado = SumarVectorGo(vector)
	}
}

func BenchmarkSumarVectorASM(b *testing.B) {
	vector := make([]int64, 100000)

	for i := range vector {
		vector[i] = int64(i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resultado = SumarVectorASM(vector)
	}
}
