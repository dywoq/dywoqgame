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