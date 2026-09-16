class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        close = {
            '(': ')',
            '[': ']',
            '{': '}'
        }

        for c in s:
            if c in "([{":
                stack.append(close[c])
            elif len(stack) == 0:
                return False
            elif stack[-1] != c:
                return False
            else:
                stack.pop(-1)

        return len(stack) == 0
        