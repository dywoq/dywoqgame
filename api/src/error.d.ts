export declare namespace Error {
	export interface Error {
		/**
		 * message returns the error message.
		 * If empty, error is not active.
		 */
		message(): string

		/**
		 * reset clears the error.
		 */
		reset(): void
	}
	/**
	 * error creates a new Error with the given message.
	 * @param message The given message.
	 */
	export function create(message: string): Error
	/**
	 * none creates a new Error with the empty message.
	 */
	export function none(): Error
}
