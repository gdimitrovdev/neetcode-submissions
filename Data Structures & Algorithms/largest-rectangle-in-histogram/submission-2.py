class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        n = len(heights)

        stack = []
        left = []
        for i in range(n):
            while stack and heights[stack[-1]] >= heights[i]:
                stack.pop()
            if stack:
                left.append(stack[-1])
            else:
                left.append(-1)
            stack.append(i)

        stack = []
        right = [n] * n
        for i in range(n - 1, -1, -1):
            while stack and heights[stack[-1]] >= heights[i]:
                stack.pop()
            if stack:
                right[i] = stack[-1]
            stack.append(i)

        ans = 0

        for i in range(n):
            ans = max(ans, heights[i] * (right[i] - left[i] - 1))

        return ans
        