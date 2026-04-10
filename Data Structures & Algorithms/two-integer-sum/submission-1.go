func twoSum(nums []int, target int) []int {
    elem := make(map[int]int)

    for idx, num := range nums {
        if val, ok := elem[target - num]; ok {
            return []int{val, idx}
        }

        elem[num] = idx
    }

    return []int{0,0}
}
