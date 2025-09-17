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

	TypeSignedInteger8bit    Type = "type-signed-integer-8-bit"
	TypeSignedInteger16Bit   Type = "type-signed-integer-16-bit"
	TypeSignedInteger32bit   Type = "type-signed-integer-32-bit"
	TypeSignedInteger64Bit   Type = "type-signed-integer-64-bit"
	TypeUnsignedInteger8bit  Type = "type-unsigned-integer-8-bit"
	TypeUnsignedInteger16Bit Type = "type-unsigned-integer-16-bit"
	TypeUnsignedInteger32bit Type = "type-unsigned-integer-32-bit"
	TypeUnsignedInteger64Bit Type = "type-unsigned-integer-64-bit"
	TypeString               Type = "type-string"

	InstructionMov       Type = "instruction-mov"
	InstructionMake      Type = "instruction-make"
	InstructionConst     Type = "instruction-const"
	InstructionAdd       Type = "instruction-add"
	InstructionMinus     Type = "instruction-minus"
	InstructionDiv       Type = "instruction-div"
	InstructionMul       Type = "instruction-mul"
	InstructionRet       Type = "instruction-ret"
	InstructionTerminate Type = "instruction-terminate"
	InstructionStderr    Type = "instruction-stderr"
	InstructionStdout    Type = "instruction-stdout"
)
