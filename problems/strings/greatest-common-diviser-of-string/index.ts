function gcdOfStrings(str1: string, str2: string): string {
	// If the order of concatenation is different,
	// they do not share the same repeating pattern.
	if (str1 + str2 !== str2 + str1) {
		return "";
	}

	// Helper function to find gcd of two numbers.
	function gcd(a: number, b: number): number {
		while (b !== 0) {
			const temp = b;
			b = a % b;
			a = temp;
		}

		return a;
	}

	const length = gcd(str1.length, str2.length);

	return str1.slice(0, length);
}

console.log(gcdOfStrings("ABCABC", "ABC"));
console.log(gcdOfStrings("ABABAB", "ABAB"));
console.log(gcdOfStrings("LEET", "CODE"));
console.log(gcdOfStrings("AAAAAB", "AAA"));
