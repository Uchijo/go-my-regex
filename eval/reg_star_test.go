package eval

import "testing"

func TestRegStar(t *testing.T) {
	tests := []struct {
		name    string
		subject RegExp
		input   string
		want    bool
	}{
		{
			"aのクリーネ閉包はaにマッチする",
			MakeRegChar("a"),
			"a",
			true,
		},
		{
			"aのクリーネ閉包は空文字にマッチする",
			MakeRegChar("a"),
			"",
			true,
		},
		{
			"aのクリーネ閉包はaaaaにマッチする",
			MakeRegChar("a"),
			"aaaa",
			true,
		},
		{
			"aのクリーネ閉包はaaaaaaaaにマッチする",
			MakeRegChar("a"),
			"aaaaaaaa",
			true,
		},
		{
			"aのクリーネ閉包はaaaaaaabにマッチしない",
			MakeRegChar("a"),
			"aaaaaaab",
			false,
		},
		{
			"aのクリーネ閉包はbaaaaaaaにマッチしない",
			MakeRegChar("a"),
			"baaaaaaa",
			false,
		},
		{
			"aのクリーネ閉包はaaabaaaaにマッチしない",
			MakeRegChar("a"),
			"aaabaaaa",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex := MakeRegStar(tt.subject)
			got := regex.Match(&tt.input)
			if got != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
