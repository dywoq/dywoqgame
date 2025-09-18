package expression

import (
	"encoding/json"
	"fmt"
)

// Expression represents expression from AST tree.
type Expression struct {
	Root  any         `json:"root"`
	Left  *Expression `json:"left,omitempty"`
	Right *Expression `json:"right,omitempty"`
}

// New creates a new pointer to expression.
func New(root any, left *Expression, right *Expression) *Expression {
	e := &Expression{root, left, right}
	return e
}

func (expr *Expression) String() string {
	prettyExpr := Expression{
		Root: expr.Root,
	}
	if expr.Left != nil {
		prettyExpr.Left = expr.Left
	}
	if expr.Right != nil {
		prettyExpr.Right = expr.Right
	}
	data, err := json.MarshalIndent(prettyExpr, "", "  ")
	if err != nil {
		return fmt.Sprintf("error formatting JSON: %v", err)
	}
	return string(data)
}
