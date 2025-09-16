package token

type Type string

const (
	Illegal Type = "illegal"
	Eof     Type = "eof"

	Identifier Type = "identifier"
	Integer    Type = "integer"
	String     Type = "string"

	Assign Type = "assign"

	ArithmeticPlus  Type = "arithmetic-plus"
	ArithmeticMinus Type = "arithmetic-minus"

	Bang Type = "logic-bang"

	LogicEqual    Type = "logic-equal"
	LogicNotEqual Type = "logic-not-equal"
	LogicLess     Type = "logic-less"
	LogicGreater  Type = "logic-greater"

	SeparatorComma     Type = "separator-comma"
	SeparatorSemicolon Type = "separator-semicolon"
	SeparatorLParen    Type = "separator-lparen"
	SeparatorRParen    Type = "separator-rparen"
	SeparatorLBrace    Type = "separator-lbrace"
	SeparatorRBrace    Type = "separator-rbrace"
)
