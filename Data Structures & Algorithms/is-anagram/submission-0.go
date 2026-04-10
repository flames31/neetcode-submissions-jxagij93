func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapS := make(map[rune]int)
    mapT := make(map[rune]int)
    
    for _, c := range s {
        mapS[c]++
    }

    for _, c := range t {
        mapT[c]++
    }

    for k, v := range mapS {
        if mapT[k] != v {
            return false
        }
    }

    return true
}
