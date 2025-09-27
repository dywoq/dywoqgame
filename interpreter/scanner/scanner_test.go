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
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			s := scanner.NewScanner()
			tokens, err := s.Scan(test.input)
			if err != nil {
				t.Fatal(err)
			}

			if test.kind != tokens[0].Kind {
				t.Errorf("got %s, want %s", tokens[0].Kind, test.kind)
			}
		})
	}
}
