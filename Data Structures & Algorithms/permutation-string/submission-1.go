func checkInclusion(s1 string, s2 string) bool {
    l1, l2 := len(s1), len(s2)

    if l1 > l2 {
        return false
    }

    toFind := make(map[byte]int)

    for i := 0; i < l1; i++ {
        toFind[s1[i]]++
    }

    for i := 0; i < l2; i++ {
        if i < l1 {
            toFind[s2[i]]--
        } else {
            isPerm := true
            for _, v := range toFind {
                if v != 0 {
                    isPerm = false
                    break
                }
            }

            if isPerm {
                return true
            } else {
                toFind[s2[i]]--
                toFind[s2[i-l1]]++
            }
        }
    }

    isPerm := true
    for _, v := range toFind {
        if v != 0 {
            isPerm = false
        }
    }

    return isPerm
}
