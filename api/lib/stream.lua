stream = {}

--- @class stream.error
stream.error = {}
--- Writes a message to the error stream.
--- @param message string The message to write.
function stream.error.write(message) end

--- Flushes the error stream, ensuring all messages are written.
function stream.error.flush() end

--- @class stream.warn
stream.warn = {}
--- Writes a message to the warning stream.
--- @param message string The message to write.
function stream.warn.write(message) end

--- Flushes the warning stream.
function stream.warn.flush() end

--- @class stream.fatal
stream.fatal = {}
--- Writes a message to the fatal stream and terminates the application immediately without cleaning.
--- @param message string The message to write.
function stream.fatal.write(message) end

--- Flushes the fatal stream.
function stream.fatal.flush() end

--- @class stream.out
stream.out = {}
--- Writes a message to the standard output stream.
--- @param message string The message to write.
function stream.out.write(message) end

--- Flushes the standard output stream.
function stream.out.flush() end

--- @class stream.debug
stream.debug = {}
--- Writes a message to the debug stream. This stream is typically only active when debug mode is on.
--- @param message string The message to write.
function stream.debug.write(message) end

--- Flushes the debug stream.
function stream.debug.flush() end
