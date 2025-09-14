import { Error } from "./error"

export declare namespace Console {
	export interface Stream {
		/**
		 * write writes the arguments to the stream/
		 * @param args Arguments to write
		 */
		write(...args: any[]): Error
		/**
		 * flush clears the stream buffer and 
		 * returns the cleared stream buffer.
		 */
		flush(): {
			buffer: string[]
			error: Error
		}
	}

	export interface Streams {
		/**
		 * Standard output stream.
		 * Messages go to the game log buffer.
		 */
		stdout: Stream
		/**
		 * Error stream. 
		 * Messages go to error buffer.
		 */
		stderr: Stream
		/**
		 * Warn stream.
		 */
		stdwarn: Stream
		/**
		 * Debug stream.
		 */
		stddebug: Stream
		/**
		 * Fatal errors streams.
		 */
		stdfatal: Stream
	}

	export interface StreamWriting {
		/**
		 * print writes the arguments to stdout.
		 * @param args Arguments to write
		*/
		print(...args: any[]): Error
		/**
		 * println writes the arguments with newline to stdout.
		 * @param args Arguments to write
		 */
		println(...args: any[]): Error
		/**
		 * warn writes the arguments to stdwarn.
		 * @param args Arguments to write
		 */
		warn(...args: any[]): Error
		/**
		 * error writes the arguments to stderr.
		 * @param args Arguments to write
		 */
		error(...args: any[]): Error
		/**
		 * debug writes the arguments to stddebug.
		 * @param args Arguments to write
		 */
		debug(...args: any[]): Error
		/**
		 * fatal writes the arguments to stdfatal.
		 * @param args Arguments to write
		 */
		fatal(...args: any[]): Error
	}

	export const StreamWriting: StreamWriting;
	export const Streams: Streams;
}
