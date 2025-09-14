export declare namespace Reflect {
	export enum ObjectKind {
		Number,
		Array,
		String,
		Boolean,
		Undefined,
		Function,
		Class
	}

	export interface Object {
		/**
		 * kind returns the kind/type of the object.
		 */
		kind(): ObjectKind

		/**
		 * name returns the name of the object (variable, function, or class name).
		 */
		name(): string

		/**
		 * value returns the value of the object if it is a primitive (number, string, boolean).
		 */
		value(): any

		/**
		 * length returns the length if the object is an array.
		 */
		length(): number

		/**
		 * properties returns an array of property names if the object is an object/class.
		 */
		properties(): string[]

		/**
		 * call calls the object if it is a function.
		 * @param args Arguments to pass to the function
		 */
		call(...args: any[]): any

		/**
		 * getProperty retrieves the value of a property from the object.
		 * @param prop Property name
		 */
		getProperty(prop: string): any

		/**
		 * hasProperty checks if a property exists on the object.
		 * @param prop Property name
		 */
		hasProperty(prop: string): boolean

		/**
		 * keys returns an array of the object's own enumerable property names.
		 */
		keys(): string[]

		/**
		 * values returns an array of the object's own property values.
		 */
		values(): any[]

		/**
		 * clone creates a shallow copy of the object.
		 */
		clone(): Object

		/**
		 * isArray returns true if the object is an array.
		 */
		isArray(): boolean

		/**
		 * isClass returns true if the object is a class.
		 */
		isClass(): boolean

		/**
		 * isFunction returns true if the object is a function.
		 */
		isFunction(): boolean

		/**
		 * isPrimitive returns true if the object is a primitive value.
		 */
		isPrimitive(): boolean

		/**
		 * invokeMethod calls a named method on the object.
		 * @param method Method name
		 * @param args Arguments to pass
		 */
		invokeMethod(method: string, ...args: any[]): any

		/**
		 * hasMethod checks if the object has a specific method.
		 * @param method Method name
		 */
		hasMethod(method: string): boolean

		/**
		 * map applies a map function to each property of the object.
		 * @param callback Function called with (value, key)
		 */
		map(callback: (value: any, key: string) => any): Object

		/**
		 * filter filters the object's properties using a callback.
		 * @param callback Function called with (value, key), returns true to keep property
		 */
		filter(callback: (value: any, key: string) => boolean): Object
	}

	/**
	 * createObject wraps a value into a Reflect Object.
	 * If no value is provided, creates an undefined Reflect Object.
	 * @param val The value to wrap
	 */
	export function createObject(val?: any): Object

	/**
	 * deepEqual performs a deep comparison between two values to determine if they are equivalent.
	 * Supports primitives, arrays, objects, and nested structures.
	 * Functions and classes are compared by reference.
	 * @param x First value
	 * @param y Second value
	 * @returns True if the values are deeply equal, false otherwise
	 */
	export function deepEqual(x: any, y: any): boolean
}
