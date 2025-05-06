package slicex

// Filter 遍历切片 s，并返回所有满足 fn 函数条件的元素组成的新切片。
// 示例：
// s := []int{1, 2, 3, 4, 5}
// result := fwslices.Filter(s, func(item int) bool { return item%2 == 0 })
// fmt.Println(result) // [2, 4]
func Filter[T any](s []T, fn func(item T) bool) []T {
	n := make([]T, 0, len(s))
	for i := range s {
		if fn(s[i]) {
			n = append(n, s[i])
		}
	}
	return n
}
