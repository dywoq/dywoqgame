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

export declare let console: ConsoleStreamsPeek & ConsoleStreamsWriting & ConsoleStreams;

//#endregion

//#region type conversions / casting
export declare interface ConversionString {
	toNumber(str: string): { result: number, error: Error }
	toBoolean(str: string): { result: boolean, error: Error }
	toUnicode(str: string): { result: string[], error: Error }
}
export declare interface ConversionNumber {
	toString(num: number): { result: string, error: Error }
	toBoolean(num: number): { result: boolean, error: Error }
}
export declare interface ConversionBoolean {
	toString(b: boolean): { result: string, error: Error }
	toNumber(b: boolean): { result: number, error: Error }
}
export declare interface ConversionArray {
	toString<T>(a: T[]): { result: string, error: Error }
}
export declare interface Conversion {
	string: ConversionString
	number: ConversionNumber
	boolean: ConversionBoolean
	array: ConversionArray
}
export declare let ConversionString: ConversionString;
export declare let ConversionNumber: ConversionNumber;
export declare let ConversionBoolean: ConversionNumber;
export declare let ConversionArray: ConversionArray;
export declare let conversion: Conversion;
//#endregion

//#region reflect
export type ReflectKind =
	"number" |
	"string" |
	"boolean" |
	"array" |
	"object" |
	"function" |
	"undefined" |
	"null"

export declare interface ReflectObjectKind {
	/** Returns the kind of object. */
	kind(): ReflectKind
}

export declare interface ReflectObjectProperties {
	/** Returns the properties of object. */
	properties(): string[]
	/** Reports whether the object has property. */
	hasProperty(property: string): boolean
	/** Returns the property if exists. */
	get(property: string): any
	/** Sets the property if exists. */
	set(key: string, value: any): void
}

export declare interface ReflectObjectCall {
	/** Calls the function/method of the object dynamically with the given arguments. */
	call(method: string, ...args: any[]): any
}

export declare interface ReflectObject extends ReflectObjectKind, ReflectObjectProperties, ReflectObjectCall { }

/** Creates and returns new reflect object with the given value. */
export declare function createReflectObject(value: any): ReflectObject

export declare interface ReflectChecks {
	/** Reports whether the value is number. */
	isNumber(value: any): boolean
	/** Reports whether the value is string. */
	isString(value: any): boolean
	/** Reports whether the value is boolean. */
	isBoolean(value: any): boolean
	/** Reports whether the value is array. */
	isArray(value: any): boolean
	/** Reports whether the value is object. */
	isObject(value: any): boolean
	/** Reports whether the value is function. */
	isFunction(value: any): boolean
	/** Reports whether the value is undefined. */
	isUndefined(value: any): boolean
	/** Reports whether the value is null. */
	isNull(value: any): boolean
	/** Reports whether the value is primitive. */
	isPrimitive(value: any): boolean;
}
export declare let ReflectChecks: ReflectChecks;
//#endregion

//#region iterating
export declare interface IteratorBase<T> {
	/** Returns any error encountered during iteration. */
	error(): Error
	/** Returns the current position. */
	position(): number
	/** Returns the current element. */
	value(): T
	/** Advances the iterator to the next element and returns true if there is a next element. */
	next(): boolean
	/** Resets the iterator to its initial state. */
	reset(): void
	/** Returns the current length of the slice. */
	length(): number
	/** Returns the current array. */
	data(): T[]
}
export declare interface IteratorForward<T> extends IteratorBase<T> { }
export declare interface IteratorReverse<T> extends IteratorBase<T> { }

export declare let IteratorForward: { new <T>(data: T[]): IteratorForward<T> }
export declare let IteratorReverse: { new <T>(data: T[]): IteratorReverse<T> }

export declare interface IteratorMethod<T> {
	iterator: {
		forward(): IteratorForward<T>
		reverse(): IteratorReverse<T>
	}
}

//#endregion 
