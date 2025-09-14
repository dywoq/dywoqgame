import { Error } from "./error"

export declare namespace Console {
	export interface Stream {
		/**
		 * Writes the arguments to the stream.
		 * @param args Arguments to write
		 * @returns An Error object indicating whether the operation succeeded
		 */
		write(...args: any[]): Error

		/**
		 * Clears the stream buffer and returns the cleared buffer along with any error.
		 * @returns An object containing the buffer and an Error object
		 */
		flush(): {
			buffer: string[]
			error: Error
		}
	}

	export interface Streams {
		/**
		 * Standard output stream.
		 * Messages are stored in the game log buffer.
		 */
		stdout: Stream

		/**
		 * Error stream.
		 * Messages are stored in the error buffer.
		 */
		stderr: Stream

		/**
		 * Warning stream.
		 * Messages are stored in the warning buffer.
		 */
		stdwarn: Stream

		/**
		 * Debug stream.
		 * Messages are stored in the debug buffer.
		 */
		stddebug: Stream

		/**
		 * Fatal error stream.
		 * Messages are stored in the fatal error buffer and may stop the game.
		 */
		stdfatal: Stream
	}

	export interface StreamWriting {
		/**
		 * Writes arguments to stdout.
		 * @param args Arguments to write
		 * @returns An Error object
		 */
		print(...args: any[]): Error

		/**
		 * Writes arguments to stdout followed by a newline.
		 * @param args Arguments to write
		 * @returns An Error object
		 */
		println(...args: any[]): Error

		/**
		 * Writes arguments to the warning stream.
		 * @param args Arguments to write
		 * @returns An Error object
		 */
		warn(...args: any[]): Error

		/**
		 * Writes arguments to the error stream.
		 * @param args Arguments to write
		 * @returns An Error object
		 */
		error(...args: any[]): Error

		/**
		 * Writes arguments to the debug stream.
		 * @param args Arguments to write
		 * @returns An Error object
		 */
		debug(...args: any[]): Error

		/**
		 * Writes arguments to the fatal error stream.
		 * @param args Arguments to write
		 * @returns An Error object
		 */
		fatal(...args: any[]): Error
	}

	export const StreamWriting: StreamWriting
	export const Streams: Streams
}
