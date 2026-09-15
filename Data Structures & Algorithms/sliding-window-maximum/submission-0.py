from heapq import heapify_max, heappop_max, heappush_max

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        if k == 1: return nums

        heap = nums[:k]
        heapify_max(heap)
        curr_window = {}
        for i in range(k):
            if nums[i] not in curr_window:
                curr_window[nums[i]] = 0
            curr_window[nums[i]] += 1

        ans = []

        for i in range(len(nums) - k + 1):
            num = heap[0]
            while len(heap) > 0 and num not in curr_window:
                heappop_max(heap)
                num = heap[0]
            if len(heap) != 0: ans.append(num)

            curr_window[nums[i]] -= 1
            if curr_window[nums[i]] == 0:
                del curr_window[nums[i]]
        
            if i + k != len(nums):
                if nums[i + k] not in curr_window:
                    curr_window[nums[i + k]] = 0
                curr_window[nums[i + k]] += 1
                heappush_max(heap, nums[i + k])

        return ans
        