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