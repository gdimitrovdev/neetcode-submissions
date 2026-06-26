func trap(height []int) int {
    ans := 0
    l, r := 0, len(height) - 1
    waterLevel := 0

    for l < r {
        if height[l] <= height[r] {
            if height[l] > waterLevel {
                ans += (r - l - 1) * (height[l] - waterLevel)
                ans -= waterLevel
                waterLevel = height[l]
            } else {
                ans -= height[l]
            }

            l++
        } else {
            if height[r] > waterLevel {
                ans += (r - l - 1) * (height[r] - waterLevel)
                ans -= waterLevel
                waterLevel = height[r]
            } else {
                ans -= height[r]
            }

            r--
        }
    }

    return ans
}
