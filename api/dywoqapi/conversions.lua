conversions = {}

--- @class conversions.string
conversions.string = {}
--- Converts `val` string to boolean.
--- If `val` is not `true`, `false`, `1`, `0`, `TRUE` or `FALSE`, it returns `nil`.
--- @param val string The string to convert.
function conversions.string.to_boolean(val) end

--- Converts `val` string to integer,
--- returns `nil` if the operation wasn't successful.
--- @param val string The string to convert.
function conversions.string.to_integer(val) end

--- @class conversions.integer
conversions.integer = {}
--- Converts `val` integer to the string,
--- returns `nil` if the operation wasn't successful.
--- @param val integer The integer to convert.
function conversions.integer.to_string(val) end

--- Converts `val` integer to the boolean,
--- returns `nil` if `val` is not 1 or 0.
--- @param val integer The integer to convert.
function conversions.integer.to_boolean(val) end

--- @class conversions.boolean
conversions.boolean = {}
--- Converts `val` boolean to the string,
--- returns `nil` if the operation wasn't successful.
--- @param val boolean The boolean to convert.
function conversions.boolean.to_string(val) end

--- Converts `val` boolean to the integer,
--- returns `nil` if the operation wasn't successful.
--- @param val boolean The boolean to convert.
function conversions.boolean.to_integer(val) end
