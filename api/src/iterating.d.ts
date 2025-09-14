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

import { Error } from "./error"

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
