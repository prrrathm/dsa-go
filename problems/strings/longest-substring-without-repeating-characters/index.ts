function lengthOfLongestSubstring(s: string): number {
	// Longest Substring without duplicate characters
	let left = 0;
	let best = 0;
	let seen = new Set();

	for (let right = 0; right < s.length; right++) {
		while (seen.has(s[right])) {
			seen.delete(s[left]);
			left++;
		}
		seen.add(s[right]);
		best = Math.max(best, right - left + 1);
	}

	return best;
}

console.log("Final", lengthOfLongestSubstring("abcabcbb"));
