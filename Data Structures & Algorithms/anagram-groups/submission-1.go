func groupAnagrams(strs []string) [][]string {
    sortedWords := map[string][]string{}

    for _, s := range strs {
        runes := []rune(s)
        sort.Slice(runes, func(i, j int) bool {
            return runes[i] < runes[j]
        })
        ss := string(runes)

        if _, ok := sortedWords[ss]; !ok {
            sortedWords[ss] = make([]string, 0)
        }

        sortedWords[ss] = append(sortedWords[ss], s)
    }

    ans := make([][]string, 0)

    for _, v := range sortedWords {
        ans = append(ans, v)
    }

    return ans
}
