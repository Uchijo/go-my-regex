package eval

import "testing"

func TestRegChar(t *testing.T) {
	tests := []struct {
		name  string
		char  string
		input string
		want  bool
	}{
		{
			name:  "aはaにマッチする",
			char:  "a",
			input: "a",
			want:  true,
		},
		{
			name:  "aはbにマッチしない",
			char:  "a",
			input: "b",
			want:  false,
		},
		{
			name:  "aは空文字にマッチしない",
			char:  "a",
			input: "",
			want:  false,
		},
		{
			name:  "空文字はaにマッチしない",
			char:  "",
			input: "a",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex := MakeRegChar(tt.char)
			got := regex.Match(&tt.input)
			if got != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, got)
			}
		})
	}
}
