package token_test

import (
	"testing"

	"github.com/dywoq/dywoqgame/interpreter/token"
)

func TestIsIdentifier(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		// valid ones
		{"pi", true},
		{"pi2_As", true},
		{"ghy.sewe", true}, // why is it valid: . is the only way to get the exported symbols from the modules, example: A.Foo

		// wrong ones
		{"2pi", false},
		{"#pi", false},
		{"/pi", false},
		{"pi)", false},
		{"pi(", false},

		{"export", false},
		{"import", false},
		{"module", false},
		{"nil", false},
		{"declare", false},

		{"}", false},
		{"{", false},
		{",", false},

		{"str", false},
		{"bool", false},
		{"i8", false},
		{"i16", false},
		{"i32", false},
		{"i64", false},
		{"u8", false},
		{"u16", false},
		{"u32", false},
		{"u64", false},
		{"void", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := token.IsIdentifier(test.input)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
