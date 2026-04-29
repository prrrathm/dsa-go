// Given an encoded string, return its decoded string.

// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

// You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

// The test cases are generated so that the length of the output will never exceed 105.

// Example 1:

// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"
// Example 2:

// Input: s = "3[a2[c]]"
// Output: "accaccacc"
// Example 3:

// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"

function decodeString(s: string): string {
	// Stack to store repeat counts (k values)
	const countStack: number[] = [];

	// Stack to store previous strings before encountering '['
	const stringStack: string[] = [];

	// Current string being built
	let currentString = "";

	// Current number being formed (could be multiple digits like "12")
	let currentNumber = 0;

	for (const char of s) {
		// Case 1: If character is a digit, build the number
		if (char >= "0" && char <= "9") {
			currentNumber = currentNumber * 10 + Number(char);
		}

		// Case 2: Opening bracket → push current state to stacks
		else if (char === "[") {
			countStack.push(currentNumber); // Save repeat count
			stringStack.push(currentString); // Save current string

			// Reset for new substring
			currentNumber = 0;
			currentString = "";
		}

		// Case 3: Closing bracket → decode the current substring
		else if (char === "]") {
			const repeatCount = countStack.pop()!; // Get last repeat count
			const previousString = stringStack.pop()!; // Get last saved string

			// Repeat current substring and append to previous string
			currentString = previousString + currentString.repeat(repeatCount);
		}

		// Case 4: Regular character → append to current string
		else {
			currentString += char;
		}
	}

	return currentString;
}

function decodeStringMe(s: string): string {
	const countStack: number[] = [];
	const stringStack: string[] = [];

	let currentCount = 0;
	let currentString = "";
	console.log("encoded string =>", s);
	for (const i of s) {
		if (i >= "0" && i <= "9") {
			currentCount = currentCount * 10 + Number(i);
			console.log("\ni=n :", " curentCount=", currentCount);
		} else if (i === "[") {
			countStack.push(currentCount);
			stringStack.push(currentString);
			console.log(
				"i=[ :",
				" curentCount=",
				currentCount,
				" curentString=",
				currentString,
			);

			currentCount = 0;
			currentString = "";
		} else if (i === "]") {
			const repeatCount = countStack.pop()!;
			const prevString = stringStack.pop()!;
			console.log(
				"i=] :",
				" repeatCount=",
				repeatCount,
				" prevString=",
				prevString,
			);

			currentString = prevString + currentString.repeat(repeatCount);
			console.log("\nsub string: ", currentString);
		} else {
			currentString += i;
			console.log("i=s :", " curentString=", currentString);
		}
	}
	return currentString;
}

console.log("\ndecoded string: ", decodeStringMe("3[a]2[bc]")); // "aaabcbc"
// console.log(decodeStringMe("3[a2[c]]")); // "accaccacc"
// console.log(decodeStringMe("2[abc]3[cd]ef")); // "abcabccdcdcdef"
