package regex

const _YI = _YIELDS_PREC
const _EQ = _EQ_PREC
const _TA = _TAKES_PREC
const _NO = _NO_PREC

/*
The unpacked precedence matrix
*/
var _PREC_MATRIX [][]uint16 = [][]uint16{
	[]uint16{_EQ, _YI, _YI, _YI, _YI, _YI, _YI, _YI, _YI, _YI, _YI, _YI, _YI},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _TA, _EQ, _TA, _YI, _NO, _EQ},
	[]uint16{_TA, _NO, _NO, _NO, _YI, _NO, _NO, _NO, _NO, _NO, _NO, _EQ, _NO},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _TA, _EQ, _TA, _YI, _NO, _EQ},
	[]uint16{_TA, _NO, _NO, _NO, _YI, _EQ, _NO, _NO, _NO, _NO, _NO, _TA, _NO},
	[]uint16{_TA, _NO, _NO, _NO, _EQ, _NO, _NO, _NO, _NO, _NO, _NO, _NO, _NO},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _YI, _NO, _EQ, _YI, _NO, _NO},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _TA, _NO, _TA, _YI, _NO, _NO},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _TA, _NO, _TA, _YI, _NO, _NO},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _TA, _EQ, _TA, _YI, _NO, _EQ},
	[]uint16{_TA, _NO, _EQ, _NO, _YI, _NO, _NO, _NO, _NO, _NO, _NO, _EQ, _NO},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _TA, _EQ, _TA, _YI, _NO, _EQ},
	[]uint16{_TA, _YI, _NO, _YI, _NO, _NO, _YI, _TA, _NO, _TA, _YI, _NO, _NO},
}
/*
The packed precedence matrix
*/
var _PREC_MATRIX_BITPACKED []uint64 = []uint64{
	14980984727966580737, 13831960122815723519, 14752790669143570427, 2755301147185181411, 3635077580152884711, 248719, 
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