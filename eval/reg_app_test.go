package eval

import "testing"

func TestRegApp(t *testing.T) {
	tests := []struct {
		name    string
		regexps []RegExp
		input   string
		want    bool
	}{
		{
			"a,aの連接はaaにマッチする",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"aa",
			true,
		},
		{
			"a,bの連接はaaにマッチする",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("b"),
			},
			"ab",
			true,
		},
		{
			"a,a,aの連接はaaaにマッチする",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"aaa",
			true,
		},
		{
			"a,a,aの連接はaaにマッチしない",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"aa",
			false,
		},
		{
			"a,aの連接はaaaにマッチしない",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"aaa",
			false,
		},
		{
			"a,aの連接はabにマッチしない",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"ab",
			false,
		},
		{
			"a,aの連接はbaにマッチしない",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"ba",
			false,
		},
		{
			"a,aの連接はbbにマッチしない",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"bb",
			false,
		},
		{
			"a,aの連接は空文字にマッチしない",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"",
			false,
		},
		{
			"a,aの連接はbbにマッチしない",
			[]RegExp{
				MakeRegChar("a"),
				MakeRegChar("a"),
			},
			"bb",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex := MakeRegApp(tt.regexps)
			got := regex.Match(&tt.input)
			if got != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
