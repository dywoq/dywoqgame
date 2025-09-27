package scanner_test

import (
	"testing"

	"github.com/dywoq/dywoqgame/interpreter/scanner"
	"github.com/dywoq/dywoqgame/interpreter/token"
)

func TestScannerScan(t *testing.T) {
	tests := []struct {
		input string
		kind  token.Kind
	}{
		// tokenizing numbers
		{"23", token.KIND_INTEGER},
		{"23.12", token.KIND_FLOAT},

		// tokenizing strings
		{`"Hi!"`, token.KIND_STRING},

		// tokenizing keywords
		{"export", token.KIND_KEYWORD},
		{"module", token.KIND_KEYWORD},
		{"import", token.KIND_KEYWORD},
		{"nil", token.KIND_KEYWORD},
		{"declare", token.KIND_KEYWORD},

		// tokenizing separators
		{",", token.KIND_SEPARATOR},
		{"{", token.KIND_SEPARATOR},
		{"}", token.KIND_SEPARATOR},

		// tokenizing types
		{"str", token.KIND_TYPE},
		{"bool", token.KIND_TYPE},
		{"i8", token.KIND_TYPE},
		{"i16", token.KIND_TYPE},
		{"i32", token.KIND_TYPE},
		{"i64", token.KIND_TYPE},
		{"u8", token.KIND_TYPE},
		{"u16", token.KIND_TYPE},
		{"u32", token.KIND_TYPE},
		{"u64", token.KIND_TYPE},
		{"void", token.KIND_TYPE},

		// tokenizing identifiers
		{"Foo", token.KIND_IDENTIFIER},
		{"Foo_2_A", token.KIND_IDENTIFIER},
		{"Foo_SD_", token.KIND_IDENTIFIER},
		{"__PrivateFoo", token.KIND_IDENTIFIER},

		// tokenizing base instructions
		{"add", token.KIND_BASE_INSTRUCTION},
		{"sub", token.KIND_BASE_INSTRUCTION},
		{"mul", token.KIND_BASE_INSTRUCTION},
		{"div", token.KIND_BASE_INSTRUCTION},
		{"write", token.KIND_BASE_INSTRUCTION},
		{"store", token.KIND_BASE_INSTRUCTION},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			s := scanner.NewScanner()
			tokens, err := s.Scan(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if test.kind != tokens[0].Kind {
				t.Errorf("got %s (literal %s), want %s", tokens[0].Kind, tokens[0].Literal, test.kind)
			}
		})
	}
}
