class Solution {
    public int findMin(int[] nums) {
        int l = 0, r = nums.length - 1;

        while (l < r) {
            int mid = l + (r - l) / 2;

            if (nums[l] > nums[mid]) {
                if (nums[mid - 1] > nums[mid]) {
                    return nums[mid];
                } else {
                    r = mid - 1;
                    continue;
                }
            }

            if (nums[mid] < nums[r]) {
                return nums[l];
            }

            if (nums[mid] > nums[r]) {
                l = mid + 1;
            }
        }

        return nums[l];
    }
}
