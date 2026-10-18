class Solution {
    public int search(int[] nums, int target) {
        int l = 0, r = nums.length - 1;

        while (l < r) {
            int m = l + (r - l) / 2;

            if (nums[m] == target) {
                return m;
            } else if (nums[l] == target) {
                return l;
            } else if (nums[r] == target) {
                return r;
            }

            if (nums[l] < nums[m] && target > nums[l] && target < nums[m]) {
                r = m - 1;
                continue;
            }
            // m is still "before rotation" but target is not in nums[l] to nums[m] range
            if (nums[l] < nums[m]) {
                l = m + 1;
                continue;
            }

            // if we reach this then "rotation point" is between l and m
            if (target > nums[m] && target < nums[r]) {
                l = m + 1;
                continue;
            }

            r = m - 1;
        }

        if (nums[l] == target) {
            return l;
        }

        return -1;
    }
}
