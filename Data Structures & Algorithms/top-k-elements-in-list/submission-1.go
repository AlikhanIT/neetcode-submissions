func topKFrequent(nums []int, k int) []int {
    // 1. Считаем частоты: map[число]количество
    counts := make(map[int]int)
    for _, n := range nums {
        counts[n]++
    }

    // 2. Создаем слайс только из уникальных чисел
    uniqueNums := make([]int, 0, len(counts))
    for n := range counts {
        uniqueNums = append(uniqueNums, n)
    }

    // 3. Сортируем только уникальные числа по их частоте в мапе
    sort.Slice(uniqueNums, func(i, j int) bool {
        return counts[uniqueNums[i]] > counts[uniqueNums[j]]
    })

    // 4. Возвращаем первые K элементов
    return uniqueNums[:k]
}