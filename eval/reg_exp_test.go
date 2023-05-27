// 色々組み合わせるテスト
package eval

import (
	"testing"
)

func TestRegExp1(t *testing.T) {
	regex := MakeRegApp(
				[]RegExp{
					MakeRegChar("a"),
					MakeRegStar(MakeRegUnion(MakeRegChar("a"), MakeRegChar("b"))),
				},
			)
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			"a(a|b)*がaにマッチする",
			"a",
			true,
		},
		{
			"a(a|b)*がaaaaaaaにマッチする",
			"aaaaaaa",
			true,
		},
		{
			"a(a|b)*がabbbbbbbにマッチする",
			"abbbbbbb",
			true,
		},
		{
			"a(a|b)*がabababababaにマッチする",
			"abababababa",
			true,
		},
		{
			"a(a|b)*が空文字にマッチしない",
			"",
			false,
		},
		{
			"a(a|b)*がbにマッチしない",
			"b",
			false,
		},
		{
			"a(a|b)*がbabaにマッチしない",
			"baba",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := regex.Match(&tt.input)
			if got != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
