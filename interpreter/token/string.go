package token

import "encoding/json"

// String converts t into the Json string presentation.
// Returns <nil> instead of the string if t is nil.
func String(t *Token) string {
	if t == nil {
		return "<nil>"
	}
	json, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(json)
}
