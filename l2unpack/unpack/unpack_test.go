package unpack

import "testing"

func TestUnpack(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for _, c := range cases {
		if got := Unpack(c.input); got != c.want {
			t.Errorf("Unpack(%q) == %q, want %q", c.input, got, c.want)
		}
	}
}
