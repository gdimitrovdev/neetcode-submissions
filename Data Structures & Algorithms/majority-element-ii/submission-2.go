func majorityElement(nums []int) []int {
    if len(nums) == 1 {
        return []int{nums[0]}
    }

    var num1, num2 int = -1, -1
    var count1, count2 int = 0, 0

    for i := 0; i < len(nums); i++ {
        if nums[i] == num1 {
            count1++
        } else if nums[i] == num2 {
            count2++
        } else if count1 == 0 {
            count1 = 1
            num1 = nums[i]
        } else if count2 == 0 {
            count2 = 1
            num2 = nums[i]
        } else {
            count1--
            count2--
        }
    }

    count1 = 0
    count2 = 0

    for _, num := range nums {
        if num == num1 {
            count1++
        }
        if num == num2 {
            count2++
        }
    }

    var ans []int;
    if count1 > len(nums) / 3 {
        ans = append(ans, num1)
    }
    if count2 > len(nums) / 3 && num1 != num2 {
        ans = append(ans, num2)
    }

    return ans
}
