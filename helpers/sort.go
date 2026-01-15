package helpers
import "slices"

var array = [][]int{
    {1, 2},
    {3, 4},
    {5, 6},
    {7, 8},
    {9, 10},
}

// По возрастаню по первому элементу
sort.Slice(array, func(i, j int) bool {
    return array[i][0] < array[j][0]
})

// По убыванию по первому элементу
sort.Slice(array, func(i, j int) bool {
    return array[i][0] > array[j][0]
})

// По возрастанию по второму элементу
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][1] < intervals[j][1]  // Индекс 1
})

// По убыванию по второму элементу
sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][1] > intervals[j][1]  // Индекс 1
})