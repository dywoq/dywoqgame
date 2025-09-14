export declare namespace Console {
	export interface Stream {
		/**
		 * write writes the arguments to the stream/
		 * @param args Arguments to write
		 */
		write(...args: any[]): void
		/**
		 * flush clears the stream buffer and 
		 * returns the cleared stream buffer.
		 */
		flush(): string[]
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
		print(...args: any[]): void
		/**
		 * println writes the arguments with newline to stdout.
		 * @param args Arguments to write
		 */
		println(...args: any[]): void
		/**
		 * warn writes the arguments to stdwarn.
		 * @param args Arguments to write
		 */
		warn(...args: any[]): void
		/**
		 * error writes the arguments to stderr.
		 * @param args Arguments to write
		 */
		error(...args: any[]): void
		/**
		 * debug writes the arguments to stddebug.
		 * @param args Arguments to write
		 */
		debug(...args: any[]): void
		/**
		 * fatal writes the arguments to stdfatal.
		 * @param args Arguments to write
		 */
		fatal(...args: any[]): void
	}

	export const StreamWriting: StreamWriting;
	export const Streams: Streams;
}
