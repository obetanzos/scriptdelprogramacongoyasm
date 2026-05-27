# Integración de Go + Ensamblador x86-64 y Benchmarking de Rendimiento

## Descripción del Proyecto

Este proyecto tiene como objetivo demostrar la integración entre el lenguaje de programación Go y código en ensamblador x86-64 (AMD64), realizando una comparación de rendimiento entre implementaciones nativas en Go y funciones desarrolladas en lenguaje ensamblador.

La experimentación incluye:

- Implementación de operaciones matemáticas simples en Go y ASM.
- Integración de funciones ensamblador dentro de un programa Go.
- Manipulación de arreglos/vectores desde ensamblador.
- Uso de benchmarks para medir rendimiento.
- Análisis estadístico de resultados obtenidos.
- Comparativa entre optimizaciones modernas del compilador Go y código ASM manual.

---

# Objetivos

## Objetivo General

Analizar el rendimiento y funcionamiento de funciones desarrolladas en lenguaje ensamblador integradas en programas Go, comparando tiempos de ejecución frente a implementaciones nativas del lenguaje.

## Objetivos Específicos

- Implementar funciones matemáticas en Go puro.
- Implementar las mismas funciones en ensamblador x86-64.
- Integrar código ASM dentro de Go.
- Ejecutar pruebas de rendimiento utilizando benchmarks.
- Analizar estadísticamente los resultados.
- Evaluar ventajas y desventajas del uso de ensamblador.

---

# Tecnologías Utilizadas

| Tecnología | Descripción |
|---|---|
| Go | Lenguaje principal del proyecto |
| ASM x86-64 | Implementación de funciones de bajo nivel |
| Go Benchmarking | Herramienta integrada para pruebas de rendimiento |
| Windows 11 | Sistema operativo utilizado |
| Arquitectura AMD64 | Arquitectura objetivo |

---

# Estructura del Proyecto

```text
.
├── main.go
├── funciones_go.go
├── funciones_amd64.s
├── operaciones_test.go
├── go.mod
└── resultados_benchmark.txt
```

---

# Benchmarks

## Comando utilizado

```bash
go clean -testcache
go test -run=^TestDummy$ -bench=Benchmark -benchmem -count=30 > resultados_benchmark.txt
```

---

# Resultados de Benchmark

| Prueba | Versión | Corridas | Promedio ns/op | Desv. estándar | CV % |
|---|---|---:|---:|---:|---:|
| Multiplicación | Go puro | 29 | 0.3008 | 0.0700 | 23.27% |
| Multiplicación | Go + ASM | 30 | 2.1167 | 0.1583 | 7.48% |
| Suma vector | Go puro | 30 | 55172.57 | 9864.75 | 17.88% |
| Suma vector | Go + ASM | 30 | 44693.60 | 9857.84 | 22.06% |

---

# Conclusiones

- El uso de ensamblador no siempre garantiza mejor rendimiento.
- Los compiladores modernos pueden superar implementaciones ASM manuales en operaciones simples.
- El ensamblador sigue ofreciendo ventajas en procesamiento iterativo y acceso directo a memoria.
- La integración Go + ASM permite desarrollar aplicaciones híbridas de alto rendimiento.
- El benchmarking es fundamental para validar objetivamente mejoras de rendimiento.

---

# Licencia

Proyecto desarrollado con fines académicos y educativos.
