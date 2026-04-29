def lengthOfLongestSubstring(s: str) -> int:
    seen = set()
    left = 0
    best = 0

    for right in range(len(s)):
        while s[right] in seen:
            print("seen", right)
            seen.remove(s[left])
            left += 1

        seen.add(s[right])
        best = max(best, right - left + 1)
        print("best", best)

    return best


print("Final Answer", lengthOfLongestSubstring("pwwkew"))
