package eval

type RegUnion struct {
	RegExp
	left  RegExp
	right RegExp
}

func (ru RegUnion) Match(s *string) bool {
	sl := *s
	sr := *s
	if result := ru.left.Match(&sl); result {
		*s = sl
		return true
	}
	if result := ru.right.Match(&sr); result {
		*s = sr
		return true
	}
	return false
}

func MakeRegUnion(left, right RegExp) RegUnion {
	return RegUnion{left: left, right: right}
}
