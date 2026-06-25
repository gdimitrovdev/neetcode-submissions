func maxArea(heights []int) int {
    l, r := 0, len(heights) - 1
    maxArea := 0

    for l <= r {
        if min(heights[l], heights[r]) * (r - l) > maxArea {
            maxArea = min(heights[l], heights[r]) * (r - l)
        }

        if heights[l] < heights[r] {
            l++
        } else {
            r--
        }
    }

    return maxArea
}
