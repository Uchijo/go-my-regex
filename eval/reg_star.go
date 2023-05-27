package eval

type RegStar struct {
	RegExp
	subject RegExp
}

func (rs RegStar) Match(s *string) bool {
	cs := *s
	for {
		slen := len(cs)
		// 空文字はマッチ（0回繰り返しに相当）
		if len(cs) == 0 {
			*s = ""
			return true
		}

		// 頭から全体を見て、毎度後ろを削っていく
		// マッチしたらcsを更新して最初から
		for i := slen; i >= 0; i-- {
			// 対象となる部分文字が空になるまでマッチしない場合は失敗
			if i == 0 {
				return false
			}
			ss := cs[:i]
			if rs.subject.Match(&ss) {
				cs = cs[i:]
				break
			}
		}
	}
}

func MakeRegStar(r RegExp) RegStar {
	return RegStar{subject: r}
}
