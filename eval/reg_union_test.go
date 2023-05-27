package eval

import "testing"

func TestRegUnion(t *testing.T) {
	tests := []struct {
		name  string
		left  RegExp
		right RegExp
		input string
		want  bool
	}{
		{
			"a,aの選択はaにマッチする",
			MakeRegChar("a"),
			MakeRegChar("a"),
			"a",
			true,
		},
		{
			"a,bの選択はbにマッチする",
			MakeRegChar("a"),
			MakeRegChar("b"),
			"b",
			true,
		},
		{
			"a,bの選択はcにマッチしない",
			MakeRegChar("a"),
			MakeRegChar("b"),
			"c",
			false,
		},
		{
			"a,bの選択は空文字にマッチしない",
			MakeRegChar("a"),
			MakeRegChar("b"),
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex := MakeRegUnion(tt.left, tt.right)
			got := regex.Match(&tt.input)
			if got != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
