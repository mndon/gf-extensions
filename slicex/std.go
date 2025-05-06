package slicex

import (
	"cmp"
	"slices"
)

// Equal 判断两个切片 s1 和 s2 是否相等。
// 示例：
// s1 := []int{1, 2, 3}
// s2 := []int{1, 2, 3}
// fmt.Println(fwslices.Equal(s1, s2)) // true
func Equal[S ~[]E, E comparable](s1, s2 S) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc 使用自定义的相等函数 eq 判断两个切片 s1 和 s2 是否相等。
// 示例：
// s1 := []string{"a", "b", "c"}
// s2 := []string{"A", "B", "C"}
// eq := func(a, b string) bool { return strings.ToLower(a) == strings.ToLower(b) }
// fmt.Println(fwslices.EqualFunc(s1, s2, eq)) // true
func EqualFunc[S ~[]E, E any](s1, s2 S, eq func(E, E) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Compare 比较两个切片 s1 和 s2 的元素。
// 返回 0 表示相等，-1 表示 s1 < s2，1 表示 s1 > s2。
// 示例：
// s1 := []int{1, 2, 3}
// s2 := []int{1, 2, 4}
// result := fwslices.Compare(s1, s2)
// fmt.Println(result) // -1
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int {
	return slices.Compare(s1, s2)
}

// CompareFunc 比较两个切片 s1 和 s2，返回一个整数：
// 如果 s1 < s2，返回负数；
// 如果 s1 == s2，返回 0；
// 如果 s1 > s2，返回正数。
// 比较是通过函数 cmp 来执行的，返回的值应符合 Go 的 `sort.Interface` 比较规则。
// 示例：
// fmt.Println(CompareFunc([]int{1, 2, 3}, []int{1, 2, 4}, func(a, b int) int { return a - b })) // 输出: -1
// fmt.Println(CompareFunc([]string{"apple", "banana"}, []string{"apple", "banana"}, func(a, b string) int { return strings.Compare(a, b) })) // 输出: 0
func CompareFunc[S ~[]E, E any](s1, s2 S, cmp func(E, E) int) int {
	return slices.CompareFunc(s1, s2, cmp)
}

// Index 返回切片 s 中第一个等于 v 的元素的索引，找不到则返回 -1。
// 示例：
// s := []int{1, 2, 3, 2}
// fmt.Println(fwslices.Index(s, 2)) // 1
func Index[S ~[]E, E comparable](s S, v E) int {
	return slices.Index(s, v)
}

// IndexFunc 返回切片 s 中第一个满足 f(x) 为 true 的元素索引。
// 如果没有找到匹配元素，则返回 -1。
// 示例：
// numbers := []int{10, 20, 30, 40}
// idx := fwslices.IndexFunc(numbers, func(n int) bool { return n > 25 })
// fmt.Println(idx) // 输出 2
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

// Contains 判断切片 s 是否包含元素 v。
// 示例：
// s := []int{1, 2, 3}
// fmt.Println(fwslices.Contains(s, 2)) // true
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return slices.Contains(s, v)
}

// ContainsFunc 判断切片 s 是否包含至少一个满足函数 f 的元素。
// 如果存在满足条件的元素，则返回 true；否则返回 false。
// 示例：
//
// fmt.Println(ContainsFunc([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })) // 输出: true
// fmt.Println(ContainsFunc([]string{"apple", "banana", "cherry"}, func(s string) bool { return len(s) > 6 })) // 输出: false
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Insert 在索引 i 处插入值 v 到切片 s 中。
// 示例：
// s := []int{1, 2, 4, 5}
// s = fwslices.Insert(s, 2, 3)
// fmt.Println(s) // [1 2 3 4 5]
func Insert[S ~[]E, E any](s S, i int, v ...E) S {
	return slices.Insert(s, i, v...)
}

// Delete 删除切片 s 中从索引 i 到 j（不包括 j）的元素。
// 示例：
// s := []int{1, 2, 3, 4, 5}
// s = fwslices.Delete(s, 1, 3)
// fmt.Println(s) // [1 4 5]
func Delete[S ~[]E, E any](s S, i, j int) S {
	return slices.Delete(s, i, j)
}

// DeleteFunc 删除切片 s 中满足函数 del 的元素。
// 示例：
// s := []int{1, 2, 3, 4, 5}
// del := func(v int) bool { return v%2 == 0 }
// s = fwslices.DeleteFunc(s, del)
// fmt.Println(s) // [1 3 5]
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S {
	return slices.DeleteFunc(s, del)
}

// Replace 替换切片 s 中索引 i 到 j（不包括 j）的元素为 v。
// 示例：
// s := []int{1, 2, 3, 4, 5}
// s = fwslices.Replace(s, 1, 4, 8, 9)
// fmt.Println(s) // [1 8 9 5]
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return slices.Replace(s, i, j, v...)
}

// Clone 返回切片 s 的副本。
// 示例：
// s := []int{1, 2, 3}
// clone := fwslices.Clone(s)
// fmt.Println(clone) // [1 2 3]
func Clone[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}

// Compact 将连续的重复元素替换为单个副本。
// 示例：
// s := []int{1, 1, 2, 2, 2, 3}
// s = fwslices.Compact(s)
// fmt.Println(s) // [1 2 3]
func Compact[S ~[]E, E comparable](s S) S {
	return slices.Compact(s)
}

// CompactFunc 移除切片 s 中相邻的重复元素，
// 相邻元素是否相等由函数 eq 决定，并返回处理后的切片。
// 示例：
// fmt.Println(CompactFunc([]int{1, 1, 2, 2, 2, 3, 4, 4}, func(a, b int) bool { return a == b })) // 输出: [1 2 3 4]
// fmt.Println(CompactFunc([]string{"a", "a", "b", "b", "c"}, func(a, b string) bool { return a == b })) // 输出: ["a" "b" "c"]
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return slices.CompactFunc(s, eq)
}

// Grow 扩展切片 s 的容量至指定的 n。
// 如果 n 小于或等于当前切片的容量，则不做任何修改。
// 示例：
// s := []int{1, 2, 3}
// Grow(&s, 10)
// fmt.Println(s) // 输出: [1 2 3 0 0 0 0 0 0 0]
func Grow[S ~[]E, E any](s S, n int) S {
	return slices.Grow(s, n)
}

// Reverse 反转切片 s。
// 示例：
// s := []int{1, 2, 3}
// fwslices.Reverse(s)
// fmt.Println(s) // [3 2 1]
func Reverse[S ~[]E, E any](s S) {
	slices.Reverse(s)
}

// Concat 将多个切片连接成一个新的切片。
// 示例：
// s1 := []int{1, 2}
// s2 := []int{3, 4}
// s3 := []int{5, 6}
// result := Concat(s1, s2, s3)
// fmt.Println(result) // 输出: [1 2 3 4 5 6]
func Concat[S ~[]E, E any](slicesToConcat ...S) S {
	return slices.Concat(slicesToConcat...)
}

// Sort 对切片 s 进行升序排序。
// 示例：
// s := []int{3, 1, 2}
// Sort(s)
// fmt.Println(s) // 输出: [1 2 3]
func Sort[S ~[]E, E cmp.Ordered](x S) {
	slices.Sort(x)
}

// SortFunc 使用自定义的比较函数对切片 s 进行排序。
// 示例：
// s := []int{3, 1, 2}
// SortFunc(s, func(a, b int) int { return b - a })
// fmt.Println(s) // 输出: [3 2 1]
func SortFunc[S ~[]E, E any](s S, cmp func(E, E) int) {
	slices.SortFunc(s, cmp)
}

// SortStableFunc 使用自定义的比较函数对切片 s 进行稳定排序。
// 稳定排序保持相等元素的原始相对顺序。
// 示例：
// s := []int{3, 1, 2}
// SortStableFunc(s, func(a, b int) int { return b - a })
// fmt.Println(s) // 输出: [3 2 1]
func SortStableFunc[S ~[]E, E any](s S, cmp func(E, E) int) {
	slices.SortStableFunc(s, cmp)
}

// IsSorted 检查切片 s 是否已经排序。
// 如果切片已经排序返回 true，否则返回 false。
// 示例：
// s := []int{1, 2, 3}
// fmt.Println(IsSorted(s)) // 输出: true
func IsSorted[S ~[]E, E cmp.Ordered](x S) bool {
	return slices.IsSorted(x)
}

// IsSortedFunc 检查切片 s 是否已经按照指定的比较函数排序。
// 如果已经排序返回 true，否则返回 false。
// 示例：
// s := []int{3, 2, 1}
// fmt.Println(IsSortedFunc(s, func(a, b int) int { return a - b })) // 输出: false
func IsSortedFunc[S ~[]E, E any](s S, cmp func(E, E) int) bool {
	return slices.IsSortedFunc(s, cmp)
}

// Min 返回切片 s 中的最小值。
// 示例：
// s := []int{3, 1, 2}
// fmt.Println(Min(s)) // 输出: 1
func Min[S ~[]E, E cmp.Ordered](x S) E {
	return slices.Min(x)
}

// MinFunc 返回切片 s 中根据自定义比较函数计算的最小值。
// 示例：
// s := []int{3, 1, 2}
// fmt.Println(MinFunc(s, func(a, b int) int { return a - b })) // 输出: 1
func MinFunc[S ~[]E, E any](s S, cmp func(E, E) int) E {
	return slices.MinFunc(s, cmp)
}

// Max 返回切片 s 中的最大值。
// 示例：
// s := []int{3, 1, 2}
// fmt.Println(Max(s)) // 输出: 3
func Max[S ~[]E, E cmp.Ordered](x S) E {
	return slices.Max(x)
}

// MaxFunc 返回切片 s 中根据自定义比较函数计算的最大值。
// 示例：
// s := []int{3, 1, 2}
// fmt.Println(MaxFunc(s, func(a, b int) int { return a - b })) // 输出: 3
func MaxFunc[S ~[]E, E any](s S, cmp func(E, E) int) E {
	return slices.MaxFunc(s, cmp)
}

// BinarySearch 在切片 s 中执行二分查找，返回元素值所在的索引，
// 如果元素不存在则返回负值。
// 示例：
// s := []int{1, 2, 3, 4, 5}
// fmt.Println(BinarySearch(s, 3)) // 输出: 2
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) {
	return slices.BinarySearch(x, target)
}

// BinarySearchFunc 在切片 s 中执行二分查找，查找的依据是自定义比较函数 cmp。
// 返回值是元素值的索引，若元素不存在则返回负值。
// 示例：
// s := []int{1, 2, 3, 4, 5}
// fmt.Println(BinarySearchFunc(s, 3, func(a, b int) int { return a - b })) // 输出: 2
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	return slices.BinarySearchFunc(x, target, cmp)
}
