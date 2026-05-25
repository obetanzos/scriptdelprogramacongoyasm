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

TEXT ·SumarVectorASM(SB), NOSPLIT, $0-24
    MOVQ vector+0(FP), SI
    MOVQ n+8(FP), CX
    XORQ AX, AX
    XORQ DX, DX

ciclo:
    CMPQ DX, CX
    JGE fin
    ADDQ (SI)(DX*8), AX
    INCQ DX
    JMP ciclo

fin:
    MOVQ AX, ret+16(FP)
    RET
    
