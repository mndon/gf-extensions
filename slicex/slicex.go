package slicex

type ordered interface {
	int | int32 | int16 | int8 | int64 | uint | uint32 | uint16 | uint8 | uint64 | float32 | float64 | string
}

type digit interface {
	int | int32 | int16 | int8 | int64 | uint | uint32 | uint16 | uint8 | uint64 | float32 | float64
}

// GroupBy
// @Description: 数组分组
// @param s 数组
// @param field 分组字段
// @return map[T2][]T
func GroupBy[T any, T2 comparable](s []T, field func(t T) T2) map[T2][]T {
	m := make(map[T2][]T)
	for i, _ := range s {
		m[field(s[i])] = append(m[field(s[i])], s[i])
	}
	return m
}

// Filter
// @Description: 列表过滤
// @param s 数组
// @param predicate 过滤函数
// @return []T
func Filter[T any](s []T, test func(item T) bool) []T {
	n := make([]T, 0, len(s))
	for i := range s {
		if test(s[i]) {
			n = append(n, s[i])
		}
	}
	return n
}

// Map
// @Description: 列表批处理
// @param s 数组
// @param mapper 处理函数
// @return []T2
func Map[T any, K any](s []T, m func(item T) K) []K {
	k := make([]K, len(s))
	for i := range s {
		k[i] = m(s[i])
	}
	return k
}

// Contains
// @Description: 元素是否在列表内
// @param arr
// @param item
// @return bool
func Contains[T comparable](s []T, item T) bool {
	for i := range s {
		if s[i] == item {
			return true
		}
	}
	return false
}

// Distinct
// @Description: 数组去重
// @param arr
// @param item
// @return bool
func Distinct[T comparable](s []T) []T {
	uniqueMap := make(map[T]struct{})
	var result []T

	for _, value := range s {
		// 如果值不存在于 map 中，将其添加到结果切片和 map 中
		if _, ok := uniqueMap[value]; !ok {
			uniqueMap[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// Concat
// @Description: 数组合并
// @param s
// @return []T
func Concat[T any](s ...[]T) []T {
	totalLen := 0
	for i := range s {
		totalLen += len(s[i])
	}

	output := make([]T, 0, totalLen)
	for i := range s {
		output = append(output, s[i]...)
	}

	return output
}

// Max
// @Description: 最大值
// @param s
// @return T
func Max[T digit](s []T) T {
	if len(s) == 0 {
		var t T
		return t
	}
	max := s[0]
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// Min
// @Description: 最小值
// @param s
// @return T
func Min[T digit](s []T) T {
	if len(s) == 0 {
		var t T
		return t
	}
	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

// Every
// @Description:  均满足
// @param s
// @param test
// @return bool
func Every[T any](s []T, test func(item T) bool) bool {
	for i := range s {
		if !test(s[i]) {
			return false
		}
	}
	return true
}

// Union
// @Description: 合并输入数据，并去重
// @param s
// @return []T
func Union[T comparable](s ...[]T) []T {
	if len(s) == 0 {
		return []T{}
	}
	hash := make(map[T]struct{})
	for i := range s {
		for j := range s[i] {
			hash[s[i][j]] = struct{}{}
		}
	}
	output := make([]T, len(hash))
	i := 0
	for k := range hash {
		output[i] = k
		i++
	}
	return output
}

// IndexOf
// @Description: 获取元素在数据的索引，不存在返回-1
// @param s
// @param item
// @return int
func IndexOf[T comparable](s []T, item T) int {
	for i := range s {
		if s[i] == item {
			return i
		}
	}
	return -1
}

// IndexOfFunc
// @Description: 获取元素在数据的索引，不存在返回-1
// @param s
// @param test
// @return int
func IndexOfFunc[T any](s []T, test func(item T) bool) int {
	for i := range s {
		if test(s[i]) {
			return i
		}
	}
	return -1
}

// Clone
// @Description: 深拷贝数组
// @param s
// @return []T
func Clone[T any](s []T) []T {
	c := make([]T, len(s))
	for i := range s {
		c[i] = s[i]
	}
	return c
}

func SortFunc[T any](s []T, less func(a T, b T) bool) []T {
	c := Clone(s)
	quickSortFunc(c, 0, len(c)-1, less)
	return c
}

// Sort creates a new slice that is sorted in ascending order. The
// given slice is not changed.
func Sort[T ordered](s []T) []T {
	c := Clone(s)
	quickSort(c, 0, len(s)-1)
	return c
}
