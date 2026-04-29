function kidsWithCandies(candies: number[], extraCandies: number): boolean[] {
	function max(arr: number[]) {
		var max = 0;
		for (let i = 0; i < arr.length; i++) {
			if (arr[i] > max) {
				max = arr[i];
			}
		}
		return max;
	}
	const maxVal = max(candies);
	console.log("max val", maxVal);
	return candies.map((x) => x + extraCandies >= maxVal);
}

console.log(kidsWithCandies([2, 3, 5, 1, 3], 3));
console.log(kidsWithCandies([4, 2, 1, 1, 2], 1));
