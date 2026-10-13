class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        l = 1
        r = max(piles)

        def round_up_div(number, div):
            if number // div != number / div:
                return number // div + 1
            return number // div

        res = r

        while l < r:
            m = l + (r - l) // 2

            ans = sum(round_up_div(pile, m) for pile in piles)

            if ans <= h:
                res = m
                r = m - 1
            else:
                l = m + 1

        ans = sum(round_up_div(pile, l) for pile in piles)
        if ans <= h:
            res = l

        return res
        