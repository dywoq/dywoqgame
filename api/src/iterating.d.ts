import { Error } from "./error"

export declare namespace Iterating {
	export interface Base<T> {
		/**
		 * position returns the current position of iterator.
		 */
		position(): number

		/**
		 * reset resets the iterator, setting the position by zero,
		 * and resetting the possible encountered error.
		 */
		reset(): void

		/**
		 * value returns the current element of the Forward iterator.
		 */
		value(): number

		/**
		 * next advances the iterator to the next element and returns true if there are more
		 * elements to iterate over.
		 * It increments the current position and check if it still within the bounds of the data slice.
		 */
		next(): boolean

		/**
		 * length returns the length of the current array.
		 */
		length(): number

		/**
		 * error returns the possible encountered error.
		 */
		error(): Error.Error
	}

	export interface Forward<T> extends Base<T> { }
	/**
	 * createForward creates new forward iterator.
	 */
	export function createForward<T>(s: T[]): Forward<T>

	export interface Reverse<T> extends Base<T> { }
	/**
	 * createForward creates new forward iterator.
	 */
	export function createReverse<T>(s: T[]): Reverse<T>
}
