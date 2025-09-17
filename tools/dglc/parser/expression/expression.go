package expression

// Expression represents expression from AST tree.
type Expression struct {
	Root  any
	Left  *Expression
	Right *Expression
}

// New creates a new pointer to expression.
func New(root any, left *Expression, right *Expression) *Expression {
	e := &Expression{root, left, right}
	return e
}
