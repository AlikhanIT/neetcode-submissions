func twoSum(nums []int, target int) []int {
    A := make([][2]int, len(nums))
    for i, num := range nums {
        A[i] = [2]int{num, i}
    }

    sort.Slice(A, func(i, j int) bool {
        return A[i][0] < A[j][0]
    })

    i, j := 0, len(nums)-1
    for i < j {
        cur := A[i][0] + A[j][0]
        if cur == target {
            if A[i][1] < A[j][1] {
                return []int{A[i][1], A[j][1]}
            } else {
                return []int{A[j][1], A[i][1]}
            }
        } else if cur < target {
            i++
        } else {
            j--
        }
    }
    return []int{}
}