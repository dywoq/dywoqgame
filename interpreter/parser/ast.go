package parser

import "github.com/dywoq/dywoqgame/interpreter/token"

type ModuleStatementType string

type Node interface {
	Node()
}

type BinaryExpression struct {
	Operator byte   `json:"operator"`
	Operands []Node `json:"operands"`
}

type ModuleStatement struct {
	Type       ModuleStatementType `json:"type"`
	Identifier string              `json:"identifier"`
}

type InstructionStatement struct {
	Arguments  []InstructionArgumentStatement `json:"arguments"`
	Identifier string                         `json:"identifier"`
}

type InstructionArgumentStatement struct {
	Type  token.Kind `json:"type"`
	Value string     `json:"value"`
}

type FunctionArgumentStatement struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type DeclarationFunctionValueStatement struct {
	Args []FunctionArgumentStatement `json:"args"`
	Body []Node                      `json:"body"`
}

type DeclarationVariableValueStatement struct {
	Value string `json:"value"`
}

type DeclarationStatement struct {
	Identifier string `json:"identifier"`
	Value      Node   `json:"value"`
	Exported   bool   `json:"exported"`
	Declared   bool   `json:"declared"`
}

type File struct {
	Statements []Node `json:"statements"`
}

const (
	ModuleStatementDeclaration ModuleStatementType = "declaration"
	ModuleStatementImporting   ModuleStatementType = "importing"
)

func (BinaryExpression) Node()                  {}
func (ModuleStatement) Node()                   {}
func (InstructionStatement) Node()              {}
func (InstructionArgumentStatement) Node()      {}
func (FunctionArgumentStatement) Node()         {}
func (DeclarationFunctionValueStatement) Node() {}
func (DeclarationVariableValueStatement) Node() {}
func (DeclarationStatement) Node()              {}
func (File) Node()                              {}
