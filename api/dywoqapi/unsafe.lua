unsafe = {}

--- @class unsafe.numeric
unsafe.numeric = {}
--- Creates a new numeric value with the specified type.
--- For example, if type is "int32",
--- the game engine will create the variable with type "int32".
--- @param type string Supported types: "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64".
--- @param val integer The value that must be within the range of type.
--- @throw "out of range" if `val` is not within the range of `type`.
function unsafe.numeric.new(type, val) end

--- Assigns src to dst if their types are the same,
--- and none of them are out of range of their types.
--- @param dst table Destination.
--- @param src table Source.
--- @throw "type mismatch" if `src` and `dst` have different types.
function unsafe.numeric.assign_to(dst, src) end

--- Copies a value from src to dst if their types are the same,
--- and none of them are out of range of their types.
--- @param dst table Destination.
--- @param src table Source.
--- @throw "type mismatch" if `src` and `dst` have different types.
function unsafe.numeric.copy(dst, src) end

--- Adds numeric values. All passed arguments must be of the same type.
--- @param ... table Numeric values to add.
--- @return table A new numeric value representing the sum.
--- @throw "type mismatch" if arguments have different types.
--- 
--- @throw "overflow" if the result exceeds the type's range.
function unsafe.numeric.add(...) end

--- Subtracts numeric values. All passed arguments must be of the same type.
--- @param ... table Numeric values.
--- @return table A new numeric value representing the difference.
--- @throw "type mismatch" if arguments have different types.
--- 
--- @throw "underflow" if the result falls below the type's range.
function unsafe.numeric.minus(...) end

--- Divides numeric values. All passed arguments must be of the same type.
--- @param ... table Numeric values.
--- @return table A new numeric value representing the quotient.
--- @throw "type mismatch" if arguments have different types.
--- 
--- @throw "division by zero" if any divisor is zero.
function unsafe.numeric.divide(...) end

--- Multiplies numeric values. All passed arguments must be of the same type.
--- @param ... table Numeric values.
--- @return table A new numeric value representing the product.
--- @throw "type mismatch" if arguments have different types.
--- 
--- @throw "overflow" if the result exceeds the type's range.
function unsafe.numeric.multiply(...) end

