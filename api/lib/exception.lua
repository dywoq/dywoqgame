exception = {}

--- Throws an exception with the given message.
--- If an exception is already in the thrown state, this function throws
--- a new exception to indicate a nested error.
--- @param message any The error message to be set.
function exception.throw(message) end

--- Reports whether an exception has been thrown.
--- @return boolean
function exception.thrown() end

--- Gets the message of the thrown exception and clears the state.
--- Use this function to catch and handle an exception.
--- @return any The exception message, or nil if no exception was thrown.
function exception.get() end

--- Terminates the game immediately without cleaning up entities.
--- This is for fatal, unrecoverable errors.
function exception.terminate() end

--- Returns a table of all nested error messages.
--- @return table A table containing all messages in the exception chain.
function exception.get_chain_messages() end

--- Outputs the latest thrown exception message, 
--- only if the internal list of the exceptions is not empty. 
function exception.output() end
