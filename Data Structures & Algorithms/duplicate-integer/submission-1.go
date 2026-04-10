func hasDuplicate(nums []int) bool {
    elem := make(map[int]bool)

    for _, num := range nums {
        if elem[num] {
            return true
        }

        elem[num] = true
    }

    return false
}
