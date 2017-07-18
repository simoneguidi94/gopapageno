package generator

const (
	_YI = iota
	_EQ = iota
	_TA = iota
	_NO = iota
)

func precToString(prec uint16) string {
	switch prec {
	case _YI:
		return "_YI"
	case _EQ:
		return "_EQ"
	case _TA:
		return "_TA"
	case _NO:
		return "_NO"
	}
	return "UNKNOWN_PREC"
}
