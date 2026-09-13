class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if len(t) > len(s): return ""

        chars = {}
        for c in t:
            if c not in chars:
                chars[c] = 0
            chars[c] += 1

        matches = 0
        needs_to_match = len(chars)
        
        shortest = float('inf')
        shortest_ans = ""

        seen = {c: 0 for c in chars}

        l, r = 0, 0
        while r < len(s):
            if s[r] in seen:
                seen[s[r]] += 1
                if seen[s[r]] == chars[s[r]]:
                    matches += 1

                while l <= r and (s[l] not in seen or seen[s[l]] > chars[s[l]]):
                    if s[l] in seen:
                        seen[s[l]] -= 1
                        if seen[s[l]] == chars[s[l]] - 1:
                            matches -= 1
                    l += 1

                if matches == needs_to_match:
                    while s[l] not in seen:
                        l += 1
                    if shortest > r - l + 1:
                        shortest = r - l + 1
                        shortest_ans = s[l:r+1]

            r += 1

        return shortest_ans
        