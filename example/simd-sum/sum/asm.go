//go:build ignore

package main

import (
	asm "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/operand"
)

//go:generate go run asm.go -out sum.s -stubs sum.go

func main() {
	asm.TEXT("Vectors", asm.NOSPLIT, "func(input []int64) int64")

	input := asm.GP64()
	inputLen := asm.GP64()
	asm.Load(asm.Param("input").Base(), input)
	asm.Load(asm.Param("input").Len(), inputLen)

	sum := asm.GP64()
	index := asm.GP64()
	asm.XORQ(sum, sum)
	asm.XORQ(index, index)

	asm.Comment("If we can't use two YMM vectors (inputLen < 8), fallback to a scalar sum.")
	asm.CMPQ(inputLen, operand.U8(8))
	asm.JL(operand.LabelRef("scalar_loop"))

	asm.Comment("Otherwise keep adding YMM vectors in the loop.")
	vecLeft := asm.YMM()
	vecRight := asm.YMM()
	asm.Comment("Y0 = input[0:4]")
	asm.VMOVDQU(operand.Mem{Base: input}, vecLeft)

	// We can efficiently calculate (in 1 cycle) the vector loop end
	// since 4 is a power of 2 like this: inputLen - (inputLen % 4).
	asm.Comment("loopEnd = inputLen - (inputLen % 4)")
	loopEnd := asm.GP64()
	asm.MOVQ(inputLen, loopEnd)
	asm.ANDQ(operand.I8(-4), loopEnd)

	/*
		i := 0
		Y0 := input[i:i+4]                   // [1, 2,  3,  4]
		for i += 4; i < loopEnd; i += 4 {
			Y1 := input[i:i+4]               // [5, 6,  7,  8]
			Y0 = Y0 + Y1                     // [6, 8, 10, 12]
		}
	*/
	asm.Label("vector_loop")
	{
		asm.ADDQ(operand.U32(4), index)
		asm.CMPQ(loopEnd, index)
		asm.JLE(operand.LabelRef("vector_loop_end")) // Exit the vector loop.

		asm.VMOVDQU(
			operand.Mem{
				Base:  input,
				Index: index,
				Scale: 8,
			},
			vecRight,
		)
		asm.VPADDQ(vecLeft, vecRight, vecLeft)

		asm.JMP(operand.LabelRef("vector_loop"))
	}
	asm.Label("vector_loop_end")

	asm.Comment("Horizontal reduction.")
	{
		vecRightLow := vecRight.AsX()
		asm.VEXTRACTI128(operand.U8(1), vecLeft, vecRightLow)

		vecLeftLow := vecLeft.AsX()
		asm.VPADDQ(vecLeftLow, vecRightLow, vecLeftLow)

		asm.VPSRLDQ(operand.U8(8), vecLeftLow, vecRightLow)
		asm.VPADDQ(vecLeftLow, vecRightLow, vecLeftLow)

		asm.VMOVQ(vecLeftLow, sum)

		asm.VZEROUPPER()
	}

	asm.Label("scalar_loop")
	asm.Comment("Summarize what we couldn't with SIMD.")
	asm.Comment("for i; i < inputLen; i++")
	{
		asm.CMPQ(inputLen, index)
		asm.JLE(operand.LabelRef("scalar_loop_end"))

		asm.ADDQ(
			operand.Mem{
				Base:  input,
				Index: index,
				Scale: 8,
			},
			sum,
		)
		asm.INCQ(index)
		asm.JMP(operand.LabelRef("scalar_loop"))
	}
	asm.Label("scalar_loop_end")

	asm.Store(sum, asm.ReturnIndex(0))
	asm.RET()

	asm.Generate()
}
