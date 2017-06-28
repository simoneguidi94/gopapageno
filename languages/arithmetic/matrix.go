package arithmetic

import (
	"math"
)

const _YI = _YIELDS_PREC
const _EQ = _EQ_PREC
const _TA = _TAKES_PREC
const _NO = _NO_PREC

/*
The unpacked precedence matrix
*/
var _ARITH_PREC_MATRIX [][]uint16 = [][]uint16{
	[]uint16{_TA, _YI, _YI, _TA, _YI, _TA},
	[]uint16{_TA, _TA, _YI, _TA, _YI, _TA},
	[]uint16{_YI, _YI, _YI, _EQ, _YI, _TA},
	[]uint16{_TA, _TA, _NO, _TA, _NO, _TA},
	[]uint16{_TA, _TA, _NO, _TA, _NO, _TA},
	[]uint16{_YI, _YI, _YI, _YI, _YI, _EQ},
}

/*
TODO: move this function in the generative part of papageno

bitPack packs the matrix into a slice of uint64 where a precedence value is represented by just 2 bits.
*/
func bitPack(matrix [][]uint16) []uint64 {
	newSize := int(math.Ceil(float64((len(matrix) * len(matrix))) / float64(32)))

	newMatrix := make([]uint64, newSize)

	setPrec := func(elem *uint64, pos uint, prec uint16) {
		bitMask := uint64(0x3 << pos)
		*elem = (*elem & ^bitMask) | (uint64(prec) << pos)
	}

	for i, _ := range matrix {
		for j, prec := range matrix[i] {
			flatElemPos := i*len(matrix) + j
			newElemPtr := &newMatrix[flatElemPos/32]
			newElemPos := uint((flatElemPos % 32) * 2)
			setPrec(newElemPtr, newElemPos, prec)
		}
	}

	return newMatrix
}

/*
The packed precedence matrix
*/
var _ARITH_PREC_MATRIX_BITPACKED []uint64 = []uint64{
	845194211396987010, 64,
}

/*
getPrecedence returns the precedence between two tokens using the bitpacked precedence matrix.
*/
func getPrecedence(token1 uint16, token2 uint16) uint16 {
	tv1 := tokenValue(token1)
	tv2 := tokenValue(token2)

	flatElemPos := tv1*_NUM_TERMINALS + tv2
	elem := _ARITH_PREC_MATRIX_BITPACKED[flatElemPos/32]
	pos := uint((flatElemPos % 32) * 2)

	return uint16((elem >> pos) & 0x3)
}
