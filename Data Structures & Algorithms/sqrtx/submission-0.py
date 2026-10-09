class Solution:
    def mySqrt(self, x: int) -> int:
        l, r = 0, x

        while True:
            m = l + (r - l) // 2

            if m * m > x:
                r = m - 1
            elif m * m == x:
                return m
            elif (m + 1) * (m + 1) > x:
                return m
            else:
                l = m + 1

        return l
        