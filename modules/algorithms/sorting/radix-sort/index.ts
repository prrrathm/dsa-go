// radixSort sorts an array using Radix Sort algorithm
function radixSort(arr: number[]): number[] {
	if (arr.length === 0) return arr; // nothing to sort

	// Step 1: Find the maximum number to determine number of digits
	const maxVal = Math.max(...arr);

	// Step 2: Apply counting sort for each digit (1s, 10s, 100s, ...)
	for (let exp = 1; Math.floor(maxVal / exp) > 0; exp *= 10) {
		countingSortByDigit(arr, exp);
	}

	return arr;
}

// countingSortByDigit sorts array based on digit at exponent place
function countingSortByDigit(arr: number[], exp: number): void {
	const n = arr.length;

	// Output array for sorted result
	const output = new Array<number>(n);

	// Count array for digits 0–9
	const count = new Array<number>(10).fill(0);

	// Step 1: Count occurrences of each digit
	for (const num of arr) {
		const digit = Math.floor(num / exp) % 10; // extract digit
		count[digit]++;
	}

	// Step 2: Convert count to cumulative positions
	for (let i = 1; i < 10; i++) {
		count[i] += count[i - 1];
	}

	// Step 3: Build output array (iterate backwards for stability)
	for (let i = n - 1; i >= 0; i--) {
		const digit = Math.floor(arr[i] / exp) % 10;
		output[count[digit] - 1] = arr[i];
		count[digit]--;
	}

	// Step 4: Copy sorted values back into original array
	for (let i = 0; i < n; i++) {
		arr[i] = output[i];
	}
}

// Example usage
const nums = [170, 45, 75, 90, 802, 24, 2, 66];

console.log(radixSort(nums));
// Output: [2, 24, 45, 66, 75, 90, 170, 802]
