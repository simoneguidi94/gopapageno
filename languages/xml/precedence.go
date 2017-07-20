package xml

const (
	_YIELDS_PREC = iota
	_EQ_PREC     = iota
	_TAKES_PREC  = iota
	_NO_PREC     = iota
)

func precToString(prec uint16) string {
	switch prec {
	case _YIELDS_PREC:
		return "YIELDS_PREC"
	case _EQ_PREC:
		return "EQ_PREC"
	case _TAKES_PREC:
		return "TAKES_PREC"
	case _NO_PREC:
		return "NO_PREC"
	}
	return "UNKNOWN_PREC"
}
