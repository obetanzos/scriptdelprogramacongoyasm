#include "textflag.h"

TEXT ·SumarASM(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    ADDQ b+8(FP), AX
    MOVQ AX, ret+16(FP)
    RET

TEXT ·MultiplicarASM(SB), NOSPLIT, $0-24
    MOVQ a+0(FP), AX
    IMULQ b+8(FP), AX
    MOVQ AX, ret+16(FP)
    RET

TEXT ·SumarVectorASM(SB), NOSPLIT, $0-32    
    MOVQ vector+0(FP), SI   // SI = Puntero al inicio del arreglo
    MOVQ vector+8(FP), CX   // CX = Longitud del vector (len)
    XORQ AX, AX             // AX = Acumulador de la suma (0)
    XORQ DX, DX             // DX = Índice de iteración (i = 0)

ciclo:
    CMPQ DX, CX
    JGE fin
    ADDQ (SI)(DX*8), AX
    INCQ DX
    JMP ciclo

fin:
    MOVQ AX, ret+24(FP)
    RET
    