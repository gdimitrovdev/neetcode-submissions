class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        def solve(max_cap):
            curr_weight = 0
            needed_days = 1
            i = 0

            while i < len(weights):
                if curr_weight + weights[i] > max_cap:
                    needed_days += 1
                    curr_weight = weights[i]
                    i += 1
                    continue

                curr_weight += weights[i]
                i += 1

            return needed_days

        l, r = max(weights), sum(weights)
        ans = r

        while l < r:
            m = l + (r - l) // 2
            
            if solve(m) <= days:
                ans = m
                r = m - 1
            else:
                l = m + 1

        if solve(l) <= days:
            ans = l

        return ans
        