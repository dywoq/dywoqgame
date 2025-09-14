// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//#region "error handling"
export declare interface Error {
	/** Returns the error message. Empty string means there's no error. */
	message(): string
	/** Resets the error, setting the message to an empty string. */
	reset(): void
}
/** Creates and returns new Error instance, with message.  */
export function createError(message?: string): Error
//#endregion

//#region "console"
export type ConsoleStreamType = "log" | "error" | "fatal" | "warning" | "debug"

export declare interface ConsoleStreamWriter {
	/** 
	 * Writes the data to the stream.
	 * May return a error.
	*/
	write(...data: any[]): Error
}

export declare interface ConsoleStreamReader {
	/** 
	 * Returns the data of stream.
	 * May return a error.
	 */
	read(): { got: any[], error: Error }
}

export declare interface ConsoleStreamFlusher {
	/** 
	 * Flushes and returns the data of stream.
	 * May return a error.
	 */
	flush(): { flushed: any[], error: Error }
}

export declare interface ConsoleStreamClear {
	/**
	 * Clears the data of stream without returning it.
	 * May return a error.
	 * @param stream The chosen stream.
	 */
	clear(stream: ConsoleStreamType): Error
}

export declare interface ConsoleStream extends ConsoleStreamFlusher, ConsoleStreamReader, ConsoleStreamWriter, ConsoleStreamClear { }

export declare interface ConsoleStreams {
	/** Stream of logs. */
	log: ConsoleStream
	/** Stream of errors. */
	error: ConsoleStream
	/** Stream of fatal errors. */
	fatal: ConsoleStream
	/** Stream of warnings. */
	warning: ConsoleStream
	/** Stream of debugging messages. */
	debug: ConsoleStream
}

export declare interface ConsoleStreamsPeek {
	/** 
	 * peek returns the latest message in the given count.
	 * May return a error.
	 */
	peek(stream: ConsoleStreamType, count?: number): { got: any[], error: Error }
}

export declare interface ConsoleStreamsWriting {
	/**
	 * Writes the arguments to the chosen stream.
	 * May return a error.
	 * @param stream The chosen stream.
	 * @param args Arguments to write.
	 */
	write(stream: ConsoleStreamType, ...args?: any[]): Error
	/**
	 * Writes the arguments to the chosen stream with new line.
	 * May return a error.
	 * @param stream The chosen stream.
	 * @param args Arguments to write.
	 */
	writeln(stream: ConsoleStreamType, ...args?: any[]): Error
}

export declare let ConsoleStreams: ConsoleStreams
export declare let ConsoleStreamsWriting: ConsoleStreamsWriting
export declare let ConsoleStreamsPeek: ConsoleStreamsPeek;

//#endregion