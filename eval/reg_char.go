package eval

type RegChar struct {
	RegExp
	char string
}

func (rc RegChar) Match(s *string) bool {
	if *s == rc.char {
		*s = ""
		return true
	}

	return false
}

func MakeRegChar(s string) RegChar {
	return RegChar{char: s}
}
