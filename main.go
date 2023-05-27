package main

import (
	"fmt"
)

func main() {

	fmt.Println(" =============== ")

	regex := RegApp{
		elems: []RegExp{
			RegUnion{
				left:  RegChar{char: "h"},
				right: RegChar{char: "o"},
			},
			RegUnion{
				left:  RegChar{char: "h"},
				right: RegChar{char: "o"},
			},
		},
	}
	input := "hhhg"
	fmt.Printf("regex.match(&input): %v\n", regex.match(&input))

	fmt.Println(" =============== ")

	regex2 := RegStar{
		subject: regex,
	}
	input = "hhoo"
	fmt.Printf("regex2.match(&input): %v\n", regex2.match(&input))
	input = "hoho"
	fmt.Printf("regex2.match(&input): %v\n", regex2.match(&input))
	input = "hoh"
	fmt.Printf("regex2.match(&input): %v\n", regex2.match(&input))
	input = "hohoh"
	fmt.Printf("regex2.match(&input): %v\n", regex2.match(&input))
}

type RegExp interface {
	match(s *string) bool
}

type RegChar struct {
	RegExp
	char string
}

func (rc RegChar) match(s *string) bool {
	if *s == rc.char {
		*s = ""
		return true
	}

	return false
}

type RegStar struct {
	RegExp
	subject RegExp
}

func (rs RegStar) match(s *string) bool {
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
			if rs.subject.match(&ss) {
				cs = cs[i:]
				break;
			}
		}
	}
}

type RegUnion struct {
	RegExp
	left  RegExp
	right RegExp
}

func (ru RegUnion) match(s *string) bool {
	sl := *s
	sr := *s
	if result := ru.left.match(&sl); result {
		*s = sl
		return true
	}
	if result := ru.right.match(&sr); result {
		*s = sr
		return true
	}
	return false
}

type RegApp struct {
	RegExp
	elems []RegExp
}

func (ra RegApp) match(s *string) bool {
	rs := *s
	for _, elem := range ra.elems {
		for i := 0; i <= len(rs); i++ {
			ss := rs[:i]

			// この正規表現についてのループを終わらせ、入力文字を切り詰める
			if elem.match(&ss) {
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
