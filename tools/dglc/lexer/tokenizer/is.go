package tokenizer

import "github.com/dywoq/dywoqgame/tools/dglc/lexer/token"

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func isKeyword(literal string) token.Type {
	switch literal {
	case "int8":
		return token.TypeSignedInteger8bit
	case "int16":
		return token.TypeSignedInteger16Bit
	case "int32":
		return token.TypeSignedInteger32bit
	case "int64":
		return token.TypeSignedInteger64Bit
	case "uint8":
		return token.TypeUnsignedInteger8bit
	case "uint16":
		return token.TypeUnsignedInteger16Bit
	case "uint32":
		return token.TypeUnsignedInteger32bit
	case "uint64":
		return token.TypeUnsignedInteger64Bit
	case "string":
		return token.TypeString
	default:
		return token.Illegal
	}
}
