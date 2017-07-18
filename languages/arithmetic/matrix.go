package arithmetic

const _YI = _YIELDS_PREC
const _EQ = _EQ_PREC
const _TA = _TAKES_PREC
const _NO = _NO_PREC

/*
The unpacked precedence matrix
*/
var _PREC_MATRIX [][]uint16 = [][]uint16{
	[]uint16{_YI, _YI, _YI, _EQ, _YI, _TA},
	[]uint16{_NO, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_YI, _YI, _TA, _TA, _YI, _TA},
	[]uint16{_NO, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_YI, _YI, _TA, _TA, _TA, _TA},
	[]uint16{_YI, _YI, _YI, _YI, _YI, _EQ},
}
/*
The packed precedence matrix
*/
var _PREC_MATRIX_BITPACKED []uint64 = []uint64{
	765799921477154880, 64, 
}

/*
getPrecedence returns the precedence between two tokens using the bitpacked precedence matrix.
*/
func getPrecedence(token1 uint16, token2 uint16) uint16 {
	tv1 := tokenValue(token1)
	tv2 := tokenValue(token2)

	flatElemPos := tv1*_NUM_TERMINALS + tv2
	elem := _PREC_MATRIX_BITPACKED[flatElemPos/32]
	pos := uint((flatElemPos % 32) * 2)

	return uint16((elem >> pos) & 0x3)
}