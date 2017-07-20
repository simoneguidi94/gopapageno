package xml

const _YI = _YIELDS_PREC
const _EQ = _EQ_PREC
const _TA = _TAKES_PREC
const _NO = _NO_PREC

/*
The unpacked precedence matrix
*/
var _PREC_MATRIX [][]uint16 = [][]uint16{
	[]uint16{_EQ, _YI, _YI, _YI, _YI, _YI, _YI, _YI, _YI},
	[]uint16{_TA, _TA, _TA, _TA, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_TA, _TA, _TA, _TA, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_TA, _TA, _TA, _TA, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_TA, _TA, _TA, _TA, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_TA, _YI, _EQ, _NO, _YI, _YI, _YI, _YI, _YI},
	[]uint16{_TA, _TA, _TA, _TA, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_TA, _TA, _TA, _TA, _NO, _TA, _TA, _TA, _TA},
	[]uint16{_TA, _YI, _EQ, _EQ, _YI, _YI, _YI, _YI, _YI},
}
/*
The packed precedence matrix
*/
var _PREC_MATRIX_BITPACKED []uint64 = []uint64{
	16909532993153400833, 12302321268114041514, 5417706, 
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