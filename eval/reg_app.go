package eval

type RegApp struct {
	RegExp
	elems []RegExp
}

func (ra RegApp) Match(s *string) bool {
	rs := *s
	for _, elem := range ra.elems {
		for i := 0; i <= len(rs); i++ {
			ss := rs[:i]

			// この正規表現についてのループを終わらせ、入力文字を切り詰める
			if elem.Match(&ss) {
				rs = rs[i:]
				break
			}

			// 連接の中で一つでもマッチしない物があるとダメ
			if len(rs) == i {
				return false
			}
		}
	}
	return rs == ""
}

func MakeRegApp(elems []RegExp) RegApp {
	return RegApp{elems: elems}
}
