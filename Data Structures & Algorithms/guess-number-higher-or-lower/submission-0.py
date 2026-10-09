# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num: int) -> int:

class Solution:
    def guessNumber(self, n: int) -> int:
        l = 1
        r = 2 ** 31 - 1

        while True:
            attempt = l + (r - l) // 2
            ans = guess(attempt)

            if ans == -1:
                r = attempt - 1
            elif ans == 1:
                l = attempt + 1
            else:
                return attempt

        return l
        